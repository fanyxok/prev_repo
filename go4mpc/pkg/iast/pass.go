package iast

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"math"
	"os"
	"path/filepath"
	"s3l/mpcfgo/config"
	"s3l/mpcfgo/pkg/calib"
	"s3l/mpcfgo/pkg/functional"
	"s3l/mpcfgo/pkg/lz"
	"strconv"
	"strings"

	cg "github.com/dave/jennifer/jen"
	"github.com/lanl/clp"
	"golang.org/x/tools/go/ast/astutil"
)

/*
*

	Constant propagation of Const

*
*/

func ConstPropagate(fset *token.FileSet, pkg *ast.Package) {
	log.Printf("[Pass]: Constant Propagete Pkg [%v]\n", pkg.Name)
	astutil.Apply(pkg,
		// replace Const variable to BasicLit
		func(c *astutil.Cursor) bool {
			switch n := c.Node().(type) {
			case *ast.Ident:
				if n.Obj == nil {
					break
				}
				if n.Obj.Pos() == n.Pos() {
					break
				}
				if n.Obj.Kind == ast.Con {
					c.Replace(NewExpr(fset, n.Obj.Decl.(*ast.ValueSpec).Values[0]))
				}
			}
			return true
		},
		// Eval BasicLit
		func(c *astutil.Cursor) bool {
			switch n := c.Node().(type) {
			case *ast.BinaryExpr:
				x, okx := Unparen(n.X).(*ast.BasicLit)
				y, oky := Unparen(n.Y).(*ast.BasicLit)
				if okx && oky {
					c.Replace(Eval(n.Op, x, y))
				}
			}
			return true
		},
	)
}

// eval int-based type only
func Eval(tok token.Token, x *ast.BasicLit, y *ast.BasicLit) *ast.BasicLit {
	if x.Kind != token.INT || y.Kind != token.INT {
		log.Panicf("Eval typed %v, %v\n", x.Kind, y.Kind)
	}
	xv, errx := strconv.ParseInt(x.Value, 10, 64)
	yv, erry := strconv.ParseInt(y.Value, 10, 64)
	if errx != nil || erry != nil {
		log.Panicf("Eval %v, %v\n", x, y)
	}

	c := int64(0)
	switch tok {
	case token.ADD:
		c = xv + yv
	case token.SUB:
		c = xv - yv
	case token.MUL:
		c = xv * yv
	case token.QUO:
		c = xv / yv
	case token.REM:
		c = xv % yv
	case token.AND:
		c = xv & yv
	case token.OR:
		c = xv | yv
	case token.XOR:
		c = xv & yv
	case token.SHL:
		c = xv << yv
	case token.SHR:
		c = xv >> yv
	default:
		log.Panicf("Not Impl %#v", tok)
	}
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: strconv.FormatInt(c, 10),
	}
}

/*
*

	Constant propagation of Var

*
*/
func VarPropagate(fset *token.FileSet, pkg *ast.Package) {
	log.Printf("[Pass]: Var Propagete Pkg [%v]\n", pkg.Name)
	decls := []*ast.DeclStmt{}
	// find all basiclit define
	ast.Inspect(pkg,
		func(n ast.Node) bool {
			if n == nil {
				return true
			}
			if d, ok := n.(*ast.DeclStmt); ok {
				vspec := d.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				if _, ok := vspec.Values[0].(*ast.BasicLit); ok {
					decls = append(decls, d)
					return false
				}
			}
			return true
		},
	)
	// check if re-assigned
	for _, decl := range decls {
		vs := decl.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
		ident := vs.Names[0]
		resigned := false
		ast.Inspect(pkg, func(n ast.Node) bool {
			if n == nil {
				return true
			}
			switch x := n.(type) {
			case *ast.AssignStmt:
				if id, ok := x.Lhs[0].(*ast.Ident); ok {
					if id.Name == ident.Name && id.Obj == ident.Obj {
						resigned = true
						return false
					}
				}
			}
			return true
		})
		// replace use ident to lit
		if !resigned {
			astutil.Apply(pkg,
				func(c *astutil.Cursor) bool {
					return true
				},
				func(c *astutil.Cursor) bool {
					if d, ok := c.Node().(*ast.Ident); ok {
						if d.Name == ident.Name && d.NamePos != ident.NamePos {
							c.Replace(vs.Values[0])
							return true
						}
					}
					return true
				},
			)
		}
	}
	astutil.Apply(pkg,
		// replace Const variable to BasicLit
		func(c *astutil.Cursor) bool {
			return true
		},
		// Eval BasicLit
		func(c *astutil.Cursor) bool {
			switch n := c.Node().(type) {
			case *ast.BinaryExpr:
				x, okx := Unparen(n.X).(*ast.BasicLit)
				y, oky := Unparen(n.Y).(*ast.BasicLit)
				if okx && oky {
					c.Replace(Eval(n.Op, x, y))
				}
			}
			return true
		},
	)
}

/*
*

	ForLoop Expand; First Round, Last Round of ForStmt

*
*/

func ExpandLoopPkg(fset *token.FileSet, info *types.Info, pkg *ast.Package) {
	for _, file := range pkg.Files {
		if ignoreTemplate(fset, file) {
			ExpandLoopFile(fset, info, file)
		}
	}
}

func ExpandLoopFile(fset *token.FileSet, info *types.Info, file *ast.File) {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			ExpandLoopFn(fset, info, fn)
		}
	}
}

func ExpandLoopFn(fset *token.FileSet, info *types.Info, fn *ast.FuncDecl) {
	astutil.Apply(fn,
		func(c *astutil.Cursor) bool {
			return true
		},
		func(c *astutil.Cursor) bool {
			switch c.Node().(type) {
			case *ast.ForStmt:
				ExpandLoopCore(fset, info, c)
			}
			return true
		},
	)
}

var label = 1

func nextSuffix() string {
	str := strconv.FormatInt(int64(label), 10)
	label++
	return "_" + str
}
func ExpandLoopCore(fset *token.FileSet, info *types.Info, c *astutil.Cursor) {

	cfor := c.Node().(*ast.ForStmt)
	cInit := cfor.Init.(*ast.AssignStmt)
	cInitExpr := cInit.Rhs[0]
	cCondExpr := cfor.Cond.(*ast.BinaryExpr).Y

	preBlock := &ast.BlockStmt{List: []ast.Stmt{}}
	preSuffix := nextSuffix()
	preInit := NewDeclFromAssign(fset, info, cInit, preSuffix)

	postBlock := &ast.BlockStmt{List: []ast.Stmt{}}
	postSuffix := nextSuffix()
	postInit := NewDeclFromAssign(fset, info, cInit, postSuffix)

	c.InsertBefore(preInit)
	renameMap := map[string]string{}
	renameMap[cInit.Lhs[0].(*ast.Ident).Name] = preInit.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Names[0].Name
	for _, v := range cfor.Body.List {
		newSt := NewStmt(fset, v)
		preBlock.List = append(preBlock.List, newSt)
	}
	replaceIdentInStmtByMap(fset, info, preBlock, preSuffix, renameMap)
	for _, v := range preBlock.List {
		c.InsertBefore(v)
	}

	renameMap = map[string]string{}
	renameMap[cInit.Lhs[0].(*ast.Ident).Name] = postInit.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Names[0].Name
	for _, v := range cfor.Body.List {
		newSt := NewStmt(fset, v)
		postBlock.List = append(postBlock.List, newSt)
	}
	replaceIdentInStmtByMap(fset, info, postBlock, postSuffix, renameMap)
	for i := len(postBlock.List) - 1; i >= 0; i-- {
		c.InsertAfter(postBlock.List[i])
	}
	postInit.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0] = &ast.BinaryExpr{
		X:  cCondExpr,
		Op: token.SUB,
		Y: &ast.BasicLit{
			Kind:  token.INT,
			Value: "1",
		},
	}
	c.InsertAfter(postInit)

	// change Init and Cond-Bound for original ForStmt
	cInitExpr = &ast.BinaryExpr{
		X:  cInitExpr,
		Op: token.ADD,
		Y: &ast.BasicLit{
			Kind:  token.INT,
			Value: "1",
		},
	}
	cCondExpr = &ast.BinaryExpr{
		X:  cCondExpr,
		Op: token.SUB,
		Y: &ast.BasicLit{
			Kind:  token.INT,
			Value: "1",
		},
	}
	cfor.Init.(*ast.AssignStmt).Rhs[0] = cInitExpr
	cfor.Cond.(*ast.BinaryExpr).Y = cCondExpr
	// Clone ForStmt for 1st round rename
	// preFor := NewStmt(fset, cfor).(*ast.ForStmt)
	// // Rename var Decl in ForStmt
	// // Insert 1st round before For
	// c.InsertBefore(preInit)
	// for _, stmt := range preFor.Body.List {
	// 	c.InsertBefore(stmt)
	// }

	// // Clone ForStmt for last round rename
	// postFor := NewStmt(fset, cfor).(*ast.ForStmt)
	// replaceIdentInStmtByRule(fset, info, postFor, "_pst", start)
	// postInit := Assign2Decl(fset, postFor.Init)
	// cexpr := postInit.(*ast.DeclStmt).Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0]
	// cexpr = &ast.BinaryExpr{
	// 	X:  cexpr,
	// 	Op: token.SUB,
	// 	Y: &ast.BasicLit{
	// 		Kind:  token.INT,
	// 		Value: "1",
	// 	},
	// }
	// for i := len(postFor.Body.List) - 1; i >= 0; i-- {
	// 	c.InsertAfter(postFor.Body.List[i])
	// }
	// c.InsertAfter(postInit)

	// // change init to init+1, change cond to cond-1
	// cInit := cfor.Init.(*ast.AssignStmt)
	// cInit.Rhs[0] = &ast.BinaryExpr{
	// 	X:  cInit.Rhs[0],
	// 	Op: token.ADD,
	// 	Y: &ast.BasicLit{
	// 		Kind:  token.INT,
	// 		Value: "1",
	// 	},
	// }
	// cCond := cfor.Cond.(*ast.BinaryExpr)
	// cCond.Y = &ast.BinaryExpr{
	// 	X:  cCond.Y,
	// 	Op: token.SUB,
	// 	Y: &ast.BasicLit{
	// 		Kind:  token.INT,
	// 		Value: "1",
	// 	},
	// }
}
func ignoreTemplate(fset *token.FileSet, file *ast.File) bool {
	filename := fset.Position(file.Pos()).Filename
	return !strings.HasSuffix(filename, "template.go")
}
func replaceIdentInStmtByMap(fset *token.FileSet, info *types.Info, stmt ast.Stmt, suffix string, nMap map[string]string) {
	astutil.Apply(stmt,
		func(c *astutil.Cursor) bool {
			return true
		},
		func(c *astutil.Cursor) bool {
			switch n := c.Node().(type) {
			case *ast.DeclStmt:
				valSpec := n.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				name := valSpec.Names[0].Name
				if nName, ok := nMap[name]; ok {
					log.Panicf("nMap find %v-%v", name, nName)
				}
				nMap[name] = name + suffix
				valSpec.Names[0].Name = nMap[name]
			case *ast.Ident:
				if nMap[n.Name] != "" {
					n.Name = nMap[n.Name]
				}
			}
			return true
		},
	)
}

/*
*
*
	Delate Unused VAR Decl
*
*
*/

func DelateUnusedGenDeclPkg(fset *token.FileSet, pkg *ast.Package) {
	for _, file := range pkg.Files {
		if ignoreTemplate(fset, file) {
			DelateUnusedGenDeclFile(fset, file)
		}

	}
}

func DelateUnusedGenDeclFile(fset *token.FileSet, file *ast.File) {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			DelateUnusedGenDeclFn(fset, fn)
		}
	}
}

func DelateUnusedGenDeclFn(fset *token.FileSet, fn *ast.FuncDecl) {
	decls := []*ast.DeclStmt{}
	ast.Inspect(fn,
		func(n ast.Node) bool {
			if n == nil {
				return true
			}
			if d, ok := n.(*ast.DeclStmt); ok {
				_ = d.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				decls = append(decls, d)
				return false
			}
			return true
		},
	)
	for _, decl := range decls {
		vs := decl.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
		ident := vs.Names[0]
		used := false
		ast.Inspect(fn, func(n ast.Node) bool {
			if n == nil {
				return true
			}
			switch x := n.(type) {
			case *ast.Ident:
				if x.Name == ident.Name && x.Obj == ident.Obj {
					if x.NamePos != ident.NamePos {
						used = true
						return false
					}
				}
			}
			return true
		})
		if !used {
			astutil.Apply(fn,
				func(c *astutil.Cursor) bool {
					return true
				},
				func(c *astutil.Cursor) bool {
					if d, ok := c.Node().(*ast.DeclStmt); ok {
						if d == decl {
							c.Delete()
							return true
						}
					}
					return true
				})
		}
	}
}

/*
*
*

	Inline All Func

*
*
*/
func InlinerPkg(fset *token.FileSet, pkg *ast.Package) {
	log.Printf("[Pass]: Inline Fn in Pkg [%v]\n", pkg.Name)
	for _, file := range pkg.Files {
		if ignoreTemplate(fset, file) {
			InlinerFile(fset, file)
		}

	}
}

func InlinerFile(fset *token.FileSet, file *ast.File) {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if fn.Name.Name == "main" {
				InlinerFn(fset, fn)
			}
		}
	}
	astutil.Apply(file, func(c *astutil.Cursor) bool {
		if c.Node() == nil {
			return true
		}
		if fn, ok := c.Node().(*ast.FuncDecl); ok {
			if fn.Name.Name != "main" {
				c.Delete()
			}
		}
		return true
	}, nil)
}

func InlinerFn(fset *token.FileSet, fn *ast.FuncDecl) {
	modified := true
	fnInlinePre := func(c *astutil.Cursor) bool {
		var rhs ast.Expr
		switch d := c.Node().(type) {
		case *ast.DeclStmt:
			rhs = d.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0]
		case *ast.AssignStmt:
			rhs = d.Rhs[0]
		case *ast.ExprStmt:
			rhs = d.X
		default:
			return true
		}
		if _, ok := rhs.(*ast.CallExpr); !ok {
			return true
		}
		d := rhs.(*ast.CallExpr)
		fnIdent := d.Fun.(*ast.Ident)
		if isBuiltin(fnIdent) {
			return true
		}
		modified = true
		callee := fnIdent.Obj.Decl.(*ast.FuncDecl)
		block := &ast.BlockStmt{
			List: []ast.Stmt{},
		}
		for _, st := range callee.Body.List {
			block.List = append(block.List, NewStmt(fset, st))
		}
		callerArgs := d.Args
		calleeArgNames := []string{}
		for _, field := range callee.Type.Params.List {
			if len(field.Names) != 1 {
				log.Printf("Fn Args share 1 Type\n")
			}
			calleeArgNames = append(calleeArgNames, field.Names[0].Name)
		}
		suffix := nextSuffix()
		astutil.Apply(block, func(cn *astutil.Cursor) bool {
			if cn.Node() == nil {
				return true
			}
			if ident, ok := cn.Node().(*ast.Ident); ok {
				for i, name := range calleeArgNames {
					if ident.Name == name {
						cn.Replace(callerArgs[i])
						return true
					}
				}
				if ident.Obj != nil && ident.Obj.Kind == ast.Var {
					cn.Replace(ast.NewIdent(fmt.Sprintf("%s_%s%s", ident.Name, callee.Name.Name, suffix)))
				}
			}
			return true
		}, nil)
		for _, st := range block.List {
			if ret, ok := st.(*ast.ReturnStmt); ok {
				if _, ok := c.Node().(*ast.DeclStmt); ok {
					c.Node().(*ast.DeclStmt).Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0] = ret.Results[0]
				} else {
					c.Node().(*ast.AssignStmt).Rhs[0] = ret.Results[0]
				}
			} else {
				c.InsertBefore(st)
			}
		}
		return true
	}
	for modified {
		modified = false
		astutil.Apply(fn, fnInlinePre, nil)
	}
}

/*
*
*
	Do Def Use Analysis (without SSA, record each Assign separately)
*
*
*/

type AnalysisResult struct {
	IdentsOf  map[string][]*ast.Ident
	DefuseMap map[*ast.Ident]UseInfo
}

func AnalyzeDefusePkg(fset *token.FileSet, info *types.Info, pkg *ast.Package) *AnalysisResult {
	log.Printf("Analyze Pkg [%s]...\n", pkg.Name)
	ainfo := new(AnalysisResult)
	ainfo.IdentsOf = make(map[string][]*ast.Ident)
	ainfo.DefuseMap = make(map[*ast.Ident]UseInfo)
	for _, file := range pkg.Files {
		if ignoreTemplate(fset, file) {
			fnmap, idmap := AnalyzeDefuseFile(fset, info, file)
			for i, v := range fnmap {
				ainfo.IdentsOf[i] = v
			}
			for i, v := range idmap {
				ainfo.DefuseMap[i] = v
			}
		}
	}
	log.Printf("...Done Analyze Pkg [%s]\n", pkg.Name)
	return ainfo
}

func AnalyzeDefuseFile(fset *token.FileSet, info *types.Info, file *ast.File) (map[string][]*ast.Ident, map[*ast.Ident]UseInfo) {
	log.Printf("Analyze File [%s]...\n", file.Name)
	usemap := make(map[*ast.Ident]UseInfo)
	identsOf := make(map[string][]*ast.Ident)
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			t := AnalyzeDefuseFn(fset, info, fn)
			for id, use := range t {
				usemap[id] = use
				identsOf[fn.Name.Name] = append(identsOf[fn.Name.Name], id)
			}
		}
	}
	log.Printf("...Done Analyze File [%s]\n", file.Name)
	return identsOf, usemap
}

type UseInfo interface {
	Weight() int
}
type normal struct {
	Tok token.Token
	X   []ast.Expr // [...]Ident or [1]BasicLit
	W   int
}
type fncall struct {
	Fn *ast.Ident
	X  []ast.Expr // [](Ident or BasicLit)
	W  int
}

func (ct *normal) Weight() int { return ct.W }
func (ct *fncall) Weight() int { return ct.W }

// types.Basic -> int, float, string, bool, etc
// types.Named -> sint32, etc
// types.Slice -> []int, []string, []sint32, etc
func AnalyzeDefuseFn(fset *token.FileSet, info *types.Info, fn *ast.FuncDecl) map[*ast.Ident]UseInfo {
	log.Printf("Analyze Fn [%s]...\n", fn.Name.Name)
	name2ident := map[string][]*ast.Ident{}
	ident2weight := map[*ast.Ident]int{}
	ident2expr := map[*ast.Ident]ast.Expr{}
	weight_outer := []int{}
	weight := 1
	astutil.Apply(fn,
		func(c *astutil.Cursor) bool {
			n := c.Node()
			if n == nil {
				return true
			}
			if n, ok := c.Node().(*ast.ForStmt); ok {
				begin := n.Init.(*ast.AssignStmt).Rhs[0].(*ast.BasicLit)
				end := n.Cond.(*ast.BinaryExpr).Y.(*ast.BasicLit)
				beginInt, _ := strconv.Atoi(begin.Value)
				endInt, _ := strconv.Atoi(end.Value)
				range_ := endInt - beginInt
				weight_outer = append(weight_outer, weight)
				weight = weight * range_
			}
			var ident *ast.Ident
			var typ types.Type
			var expr ast.Expr
			switch d := n.(type) {
			case *ast.DeclStmt:
				vs := d.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				ident = vs.Names[0]
				typ = info.TypeOf(ident)
				expr = vs.Values[0]
			case *ast.AssignStmt:
				ident = identOfExpr(d.Lhs[0])
				typ = info.TypeOf(ident)
				expr = d.Rhs[0]
			}
			if ident == nil {
				return true
			}
			switch t := typ.(type) {
			case *types.Basic:
				// All Basic types are public value
				return false
			case *types.Named, *types.Slice:
				// secret types are Named
				name := t.String()
				if isSecPrimType(name) {
					name2ident[ident.Name] = append(name2ident[ident.Name], ident)
					ident2expr[ident] = expr
					ident2weight[ident] = weight
				}
			default:
				fmt.Printf("typ: %v, %#v\n", typ, typ)
			}
			return true
		},
		func(c *astutil.Cursor) bool {
			if _, ok := c.Node().(*ast.ForStmt); ok {
				weight = weight_outer[len(weight_outer)-1]
				weight_outer = weight_outer[:len(weight_outer)-1]
			}
			return true
		})
	// config.DetailLog(func() {
	// 	log.Printf("[DefuseFn]fn.Name: %v\n", fn.Name)
	// 	names := make([]string, 0, len(name2ident))
	// 	for k := range name2ident {
	// 		names = append(names, k)
	// 	}
	// 	sort.Strings(names)
	// 	for _, name := range names {
	// 		str := fmt.Sprintf("name: %v. ", name)
	// 		for _, v := range name2ident[name] {
	// 			pos := fset.Position(v.NamePos)
	// 			str += fmt.Sprintf("%v:%v ", pos.Line, pos.Column)
	// 		}
	// 		log.Println(str)
	// 	}
	// })
	// idents := make([]*ast.Ident, 0, len(ident2wexpr))
	// for i := range ident2wexpr {
	// 	idents = append(idents, i)
	// }
	// sort.Slice(idents, func(i, j int) bool {
	// 	return idents[i].NamePos < idents[j].NamePos
	// })

	ident2use := map[*ast.Ident]UseInfo{}
	for ident, expr := range ident2expr {
		t := new(normal)
		ident2use[ident] = t
		t.W = ident2weight[ident]
		switch x := expr.(type) {
		case *ast.BasicLit:
			t.Tok = token.ASSIGN
			t.X = append(t.X, x)
		case *ast.BinaryExpr:
			t.Tok = x.Op
			t.X = append(t.X, defOfIdentInFn(fn, identOfExpr(x.X)))
			t.X = append(t.X, defOfIdentInFn(fn, identOfExpr(x.Y)))
		case *ast.UnaryExpr:
			t.Tok = x.Op
			if isBasicLit(x.X) {
				t.X = append(t.X, x.X)
			} else {
				t.X = append(t.X, defOfIdentInFn(fn, identOfExpr(x.X)))
			}
		case *ast.Ident:
			t.Tok = token.ASSIGN
			t.X = []ast.Expr{identOfExpr(x)}
		case *ast.IndexExpr:
			t.Tok = token.ASSIGN
			t.X = []ast.Expr{defOfIdentInFn(fn, identOfExpr(x.X))}
		case *ast.SliceExpr:
			t.Tok = token.ASSIGN
			t.X = []ast.Expr{defOfIdentInFn(fn, identOfExpr(x.X))}
		case *ast.CallExpr:
			t := new(fncall)
			ident2use[ident] = t
			t.W = ident2weight[ident]
			t.Fn = identOfExpr(x.Fun)
			if n := t.Fn.Name; n == "make" || n == "append" {
				break
			}
			for _, arg := range x.Args {
				if isBasicLit(arg) {
					t.X = append(t.X, arg)
				} else {
					t.X = append(t.X, defOfIdentInFn(fn, identOfExpr(arg)))
				}
			}
		default:
		}
	}
	// config.DetailLog(func() {

	// 	for _, ident := range idents {
	// 		str := ""
	// 		switch d := ident2use[ident].(type) {
	// 		case *normal:
	// 			if lit, ok := d.X[0].(*ast.BasicLit); ok {
	// 				config.DetailLog(func() {
	// 					log.Printf("%20s =%2v= %s %s\n", ident.Name, d.W, d.Tok, lit.Value)
	// 				})
	// 			} else {
	// 				for _, v := range d.X {
	// 					if ident, ok := v.(*ast.Ident); ok {
	// 						str += fmt.Sprintf(" %s<-%s,", fmtIdentOrLit(fset, v), fmtIdentOrLit(fset, defOfIdentInFn(fn, ident)))
	// 					} else if lit, ok := v.(*ast.BasicLit); ok {
	// 						str += " " + lit.Value
	// 					} else {
	// 						log.Panicf("%#v", v)
	// 					}
	// 				}
	// 				if str != "" {
	// 					str = str[:len(str)-1]
	// 				}
	// 				config.DetailLog(func() {
	// 					log.Printf("%20s =%2v= %s %s\n", ident.Name, d.W, d.Tok, str)
	// 				})
	// 			}
	// 		case *fncall:
	// 			for _, v := range d.X {
	// 				if ident, ok := v.(*ast.Ident); ok {
	// 					str += fmt.Sprintf(" %s<-%s,", fmtIdentOrLit(fset, v), fmtIdentOrLit(fset, defOfIdentInFn(fn, ident)))
	// 				} else if lit, ok := v.(*ast.BasicLit); ok {
	// 					str += " " + lit.Value
	// 				} else {
	// 					log.Panicf("%#v", v)
	// 				}
	// 			}
	// 			if str != "" {
	// 				str = str[:len(str)-1]
	// 			}
	// 			config.DetailLog(func() {
	// 				log.Printf("%20s =%2v= %s %s\n", ident.Name, d.W, d.Fn.Name, str)
	// 			})
	// 		default:
	// 			log.Panicf("unexpected ident2use: %#v, %#v", ident, ident2use[ident])
	// 		}
	// 	}
	// })
	log.Printf("...Done Analyze Fn [%s]\n", fn.Name.Name)
	return ident2use
}

var fn_id2def = make(map[string]map[*ast.Ident]*ast.Ident)

func defOfIdentInFn(fn *ast.FuncDecl, ident *ast.Ident) *ast.Ident {
	if fn_id2def[fn.Name.Name] == nil {
		lastDefine := make(map[string]*ast.Ident)
		id2def := make(map[*ast.Ident]*ast.Ident)
		// add inputs into map
		for _, v := range fn.Type.Params.List {
			for _, v := range v.Names {
				lastDefine[v.Name] = v
			}
		}
		// postorder traversal
		astutil.Apply(fn, nil, func(c *astutil.Cursor) bool {
			if declStmt, ok := c.Node().(*ast.DeclStmt); ok {
				valSpec := declStmt.Decl.(*ast.GenDecl).Specs[0].(*ast.ValueSpec)
				for _, v := range valSpec.Names {
					lastDefine[v.Name] = v
				}
			}
			if assignStmt, ok := c.Node().(*ast.AssignStmt); ok {
				lhs := assignStmt.Lhs[0]
				id := exprOfExpr(lhs).(*ast.Ident)
				lastDefine[id.Name] = id
			}

			if id, ok := c.Node().(*ast.Ident); ok {
				//fmt.Printf("record: %v, %v\n", id, lastDefine[id.Name])
				id2def[id] = lastDefine[id.Name]
			}
			return true
		})
		fn_id2def[fn.Name.Name] = id2def
	}
	id2def := fn_id2def[fn.Name.Name]
	if id, ok := id2def[ident]; ok {
		//fmt.Printf("ident %v, id: %v\n", ident, id)
		return id
	}
	log.Panicf("")
	return nil
}

/*
*
*
	Optimize conversion by Solver
*
*
*/
// Token except token.Token
const (
	NEW token.Token = 100 + iota
	NEWN
	CALL
)

func MP_Optimize(fset *token.FileSet, info *types.Info, pkg *ast.Package, result *AnalysisResult) (map[*ast.Ident]int, []int) {
	ident2idx := make(map[*ast.Ident]int)
	idx2ident := make(map[int]*ast.Ident)
	idx := 0
	for ident := range result.DefuseMap {
		ident2idx[ident] = idx
		idx2ident[idx] = ident
		idx++
	}
	n := idx
	log.Printf("MP_Optimize %v Variables\n", n)
	// a == 1 denote sec use A,
	// y == 1 denote sec use Y
	// s == 1 denote sec convert from A to Y
	// t == 1 denote sec convert from Y to A
	// base of a,y,s,t. all Zero
	base := make([]float64, 4*n)
	// coefficient of a,y is the decl/assign cost
	// coefficient of s,t is the conversion cost
	coef := make([]float64, len(base))
	for idx, ident := range idx2ident {
		use := result.DefuseMap[ident]
		switch d := use.(type) {
		case *normal:
			// new/newfrom
			if _, ok := d.X[0].(*ast.BasicLit); ok && len(d.X) == 1 {
				typ := info.TypeOf(ident)
				if d.Tok != token.ILLEGAL {
					coef[idx], coef[idx+n] = _CostOfTok(d.Tok, typ.String(), d.W)
				}
				coef[idx+2*n], coef[idx+3*n] = _CostOfConvert(typ.String(), d.W)
				// unary/binary expr
			} else {
				typ := info.TypeOf(d.X[0])
				if d.Tok != token.ILLEGAL {
					coef[idx], coef[idx+n] = _CostOfTok(d.Tok, typ.String(), d.W)
				}
				coef[idx+2*n], coef[idx+3*n] = _CostOfConvert(typ.String(), d.W)
			}
		case *fncall:
			switch d.Fn.Name {
			case "make", "append":
				typ := info.TypeOf(ident)
				coef[idx], coef[idx+n] = 0, 0
				coef[idx+2*n], coef[idx+3*n] = _CostOfConvert(typ.String(), d.W)
			case "b1n", "i8n", "i16n", "i32n", "i64n", "f32n", "f64n":
				typ := info.TypeOf(ident)
				coef[idx], coef[idx+n] = _CostOfTok(NEWN, typ.String(), d.W)
				coef[idx+2*n], coef[idx+3*n] = _CostOfConvert(typ.String(), d.W)
			default:
				fmt.Printf("fncall.Fn: %v\n", d)
			}
		}
	}

	// bound of ayst
	bound := make([][2]float64, len(base))
	for i := 0; i < len(base); i++ {
		bound[i][0] = 0.0
		bound[i][1] = 1.0
	}
	tolerate := 1e-5
	// constraint, matrix has linear constraints format L <= Ax <= R
	var matrix [][]float64
	// a[i]+y[i] = 1 for every Ident
	for idx := range idx2ident {
		i := idx
		var a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
		a[i] = 1.0
		y[i] = 1.0
		matrix = append(matrix, functional.Append[float64](1.0, a, y, s, t, 1.0))
	}
	// a[i] >= s[i] for every Ident
	for idx := range idx2ident {
		i := idx
		a, y, s, t := make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
		s[i] = -1
		a[i] = 1
		matrix = append(matrix, functional.Append[float64](0.0, a, y, s, t, 1.0))
	}
	// y[i] >= t[i] for every Ident
	for idx := range idx2ident {
		i := idx
		a, y, s, t := make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
		t[i] = -1
		y[i] = 1
		matrix = append(matrix, functional.Append[float64](0.0, a, y, s, t, 1.0))
	}
	//  a[j] == a[i], y[j] == y[i] for every Slice/Array
	//  s[j] == s[i], t[j] == t[i]
	name2ident := make(map[string]*ast.Ident)
	for idx, ident := range idx2ident {
		if _, ok := info.TypeOf(ident).(*types.Slice); ok {
			if _, ok := name2ident[ident.Name]; !ok {
				name2ident[ident.Name] = ident
			}
			last := name2ident[ident.Name]
			if last == ident {
				continue
			}
			if ident.Name == last.Name {
				a, y, s, t := make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				a[idx] = 1
				a[ident2idx[last]] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				y[idx] = 1
				y[ident2idx[last]] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				s[idx] = 1
				s[ident2idx[last]] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				t[idx] = 1
				t[ident2idx[last]] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				name2ident[ident.Name] = ident
			}
		}
	}

	// constraint conversion due to def-use info
	for idx, ident := range idx2ident {
		useinfo := result.DefuseMap[ident]
		if _, ok := useinfo.(*fncall); ok {
			continue
		}
		norma := useinfo.(*normal)
		for _, expr := range norma.X {
			// BasicLit is New(), Free Variable
			if _, ok := expr.(*ast.BasicLit); ok && len(norma.X) == 1 {
				continue
			}
			rident := expr.(*ast.Ident)
			ridx := ident2idx[rident]

			// Direct Assignment, They are the same one
			if norma.Tok == token.ASSIGN && len(norma.X) == 1 {
				var a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				a[idx] = 1
				a[ridx] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				y[idx] = 1
				y[ridx] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				s[idx] = 1
				s[ridx] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
				t[idx] = 1
				t[ridx] = -1
				matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, tolerate))
				continue
			}
			// 1: s[j] >= y[i] - y[j] ==> s[j] - y[i] + y[j] >=0
			var a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
			y[idx] = -1.0
			y[ridx] = 1.0
			s[ridx] = 1.0
			matrix = append(matrix, functional.Append[float64](-tolerate, a, y, s, t, 1.0+tolerate))
			// 4: t[j] >= a[i] - a[j] ==> t[j] - a[i] + a[j] >=0
			a, y, s, t = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
			a[idx] = -1.0
			a[ridx] = 1.0
			t[ridx] = 1.0
			matrix = append(matrix, functional.Append[float64](0.0, a, y, s, t, 1.0))
		}
	}
	prob := clp.NewSimplex()
	prob.EasyLoadDenseProblem(
		//         A    B    C
		coef,
		bound,
		matrix,
	)
	prob.SetOptimizationDirection(clp.Minimize)
	prob.SetPrimalTolerance(tolerate * 10)
	prob.Primal(clp.NoValuesPass, clp.NoStartFinishOptions)
	log.Printf("MP_Optimize Tolerance: %v\n", prob.PrimalTolerance())
	log.Printf("MP_Optimize Value: %v microseconds <==> %v seconds\n", prob.ObjectiveValue(), prob.ObjectiveValue()/1e6)

	soln := prob.PrimalColumnSolution()
	_VerifyTolerate(tolerate*100, soln)
	sol := make([]int, len(soln))
	for i := range soln {
		sol[i] = int(math.Round(soln[i]))
	}
	// for idx, ident := range idx2ident {
	// 	fmt.Printf("%50s: %v %v %v %v\n", ident.Name, sol[idx], sol[n+idx], sol[2*n+idx], sol[3*n+idx])
	// }
	_VerifyArrayType(result, info, ident2idx, sol)
	_VerifyNormal(result, ident2idx, sol)
	// mps := filepath.Join(config.Root, "config", "mpc.mps")
	// log.Printf("WriteMPS to [%s]", mps)
	// if !prob.WriteMPS(mps) {
	// 	log.Panicln("prob.WriteMPS(): falied")
	// }
	return ident2idx, sol
}

func _VerifyTolerate(tolerate float64, soln []float64) {
	log.Println("MP_Optimize Verify Tolerate")
	for _, v := range soln {
		if math.Abs(1-v) > tolerate && math.Abs(v) > tolerate {
			log.Panicf("MP_Optimize Tolerate Exceeded: %f", v)
		}
	}
}
func _VerifyArrayType(result *AnalysisResult, info *types.Info, ident2idx map[*ast.Ident]int, soln []int) {
	log.Println("MP_Optimize Verify ArrayType")
	name2type := make(map[string]int)
	for ident := range result.DefuseMap {
		if _, ok := info.TypeOf(ident).(*types.Slice); ok {
			if _, ok := name2type[ident.Name]; !ok {
				name2type[ident.Name] = soln[ident2idx[ident]]
			}
			if name2type[ident.Name] == soln[ident2idx[ident]] {
				continue
			} else {
				log.Panicf("VerifyArrayType: %s %d %d", ident.Name, name2type[ident.Name], soln[ident2idx[ident]])
			}
		}
	}
}

func _VerifyNormal(result *AnalysisResult, ident2idx map[*ast.Ident]int, soln []int) {
	log.Println("MP_Optimize Verify Normal")
	n := len(ident2idx)
	for ident, idx := range ident2idx {
		useinfo := result.DefuseMap[ident]
		if _, ok := useinfo.(*fncall); ok {
			continue
		}
		norma := useinfo.(*normal)
		if _, ok := norma.X[0].(*ast.BasicLit); ok && len(norma.X) == 1 {
			continue
		}
		config.DetailLog(func() {
			log.Printf("ident: %s %#v\n", ident, norma.X)
		})
		for _, expr := range norma.X {
			rident := expr.(*ast.Ident)
			ridx := identToIdx(ident2idx, rident)
			config.DetailLog(func() {
				log.Printf("rident: %#v\n", rident)
			})
			la, ly, ls, lt := soln[idx], soln[n+idx], soln[2*n+idx], soln[3*n+idx]
			ra, ry, rs, rt := soln[ridx], soln[n+ridx], soln[2*n+ridx], soln[3*n+ridx]
			if (la == 1 && ra == 1) || (la == 1 && rt == 1) || (ly == 1 && ry == 1) || (ly == 1 && rs == 1) {
				if ls == 1 || lt == 1 {
					log.Printf("Conversion Occurs at %s [%v %v %v %v]", rident.Name, la, ly, ls, lt)
				}
			} else {
				log.Panicf("VerifyNormal failed: \n%30s [%v,%v,%v,%v]\n%30s [%v,%v,%v,%v]",
					ident.Name, soln[idx], soln[n+idx], soln[2*n+idx], soln[3*n+idx],
					rident.Name, soln[ridx], soln[n+ridx], soln[2*n+ridx], soln[3*n+ridx])
			}
			config.DetailLog(func() {
				log.Printf("VerifyNormal: %4v [%v,%v,%v,%v] %4v [%v,%v,%v,%v]\n", idx, la, ly, ls, lt, ridx, ra, ry, rs, rt)
			})
		}
		config.DetailLog(func() {
			log.Println()
		})
	}
}

var opcostpath = filepath.Join(config.Root, "cmd", "calibration", "opcost.json")
var costpath = filepath.Join(config.Root, "cmd", "calibration", "cost.json")
var opcost = calib.LoadOpCostJson(opcostpath)
var cost = calib.LoadCostJson(costpath)

func _LongType2JsonKey(typ string) string {
	for strings.HasPrefix(typ, "[]") {
		typ = typ[2:]
	}
	splited := strings.Split(typ, ".")
	typ = splited[len(splited)-1]
	typ = strings.ReplaceAll(typ, "bool", "b1")
	typ = strings.ReplaceAll(typ, "int", "i")
	typ = strings.ReplaceAll(typ, "float", "f")
	return typ
}
func _CostOfTok(tok token.Token, typ string, weight int) (float64, float64) {
	opJsonKey := tok2JsonKey(tok)
	opcell := opcost[opJsonKey]
	singleTyp := _LongType2JsonKey(typ)
	var oprandType string
	if opJsonKey == "Not" {
		oprandType = "sb1"
	} else if opJsonKey == "Mux" {
		oprandType = fmt.Sprintf("sb1_%s_%s", singleTyp, singleTyp)
	} else {
		oprandType = fmt.Sprintf("%s_%s", singleTyp, singleTyp)
	}
	typedcell := opcell[oprandType]
	a := typedcell.Cost["A"] * float64(weight)
	y := typedcell.Cost["Y"] * float64(weight)
	config.DetailLog(func() {
		log.Printf("Cost [%s] of [%s] weighted [%v]: %f, %f\n", opJsonKey, oprandType, weight, a, y)
	})
	return a, y
}
func _CostOfConvert(typ string, weight int) (float64, float64) {
	typjkey := _LongType2JsonKey(typ)
	s := cost["A2Y"][typjkey] * float64(weight)
	t := cost["Y2A"][typjkey] * float64(weight)
	config.DetailLog(func() {
		log.Printf("Cost Conversion weighted [%v]: %f, %f\n", weight, s, t)
	})

	return s, t
}

/*
*

	Code Generation

*
*/

var (
	_net  = cg.Id("net")
	_role = 0
	_addr = "127.0.0.1:23344"
)

var (
	_pvt    = "s3l/mpcfgo/pkg/type/pvt"
	_pub    = "s3l/mpcfgo/pkg/type/pub"
	_triple = "s3l/mpcfgo/pkg/primitive/triple"
)
var (
	_in             []cg.Code = []cg.Code{}
	__in_cnt                  = 0
	_nextIn                   = func() string { __in_cnt++; return fmt.Sprintf("_in%d", __in_cnt) }
	_secLitImplType string    = "MissSecLitImplType"
)

func CodeGen(fset *token.FileSet, info *types.Info, pkg *ast.Package, ident2idx map[*ast.Ident]int, soln []int, dstDir string) {
	for _, file := range pkg.Files {
		if ignoreTemplate(fset, file) {
			for i := range make([]struct{}, 2) {
				_role = i
				f := CodeGenFn(fset, info, file, ident2idx, soln)
				if f != nil {
					subDir := filepath.Join(dstDir, strconv.FormatInt(int64(i), 10))
					if err := os.Mkdir(subDir, 0770); err != nil {
						if !os.IsExist(err) {
							log.Fatalf("Mkdir [%s] Error %s\n", subDir, err)
						}
					}
					target := filepath.Join(dstDir, strconv.FormatInt(int64(i), 10), file.Name.Name+".go")
					err := f.Save(target)
					if err != nil {
						config.DetailLog(func() {
							log.Panicf("CodeGen Save Error:\n%v\n######Error Info Above######", err)
						})
					}
					log.Printf("Write [%s] into [%s]\n", file.Name.Name, target)
				}
			}
		}
	}
}

func CodeGenFn(fset *token.FileSet, info *types.Info, file *ast.File, ident2idx map[*ast.Ident]int, soln []int) *cg.File {
	for _, v := range file.Decls {
		if fn, ok := v.(*ast.FuncDecl); ok {
			if fn.Name.Name == "main" {
				log.Printf("Gen [main] of File [%s]...\n", file.Name)
				return CodeGenFnCore(fset, info, fn, ident2idx, soln)
			}
		}
	}
	return nil
}
func CodeGenFnCore(fset *token.FileSet, info *types.Info, fn *ast.FuncDecl, ident2idx map[*ast.Ident]int, soln []int) *cg.File {
	ser := _role == 0
	l := len(ident2idx)
	f := cg.NewFile(fn.Name.Name)

	var codes = [][]cg.Code{{}}
	var pregen func(c *astutil.Cursor) bool
	var postgen func(c *astutil.Cursor) bool
	// general traversl function
	// rhs are solved in CgNromalExpr
	// thus, rhs need solver information
	pregen = func(c *astutil.Cursor) bool {
		code := codes[len(codes)-1]
		var cmd *cg.Statement
		node := c.Node()
		switch n := node.(type) {
		case *ast.GenDecl:
			if n.Tok != token.VAR {
				log.Panicf("%#v\n", n)
			}
			lhs := n.Specs[0].(*ast.ValueSpec).Names[0]
			rhs := n.Specs[0].(*ast.ValueSpec).Values[0]
			var curlsoln []int
			// if in ident2idx, it is secret, soln should be non-nil
			var idx int
			var ok bool
			if idx, ok = ident2idx[lhs]; ok {
				curlsoln = []int{soln[idx], soln[l+idx], soln[2*l+idx], soln[3*l+idx]}
				_secLitImplType = secImplTypeOfExpr(info, lhs)
			}
			cmd = _CgExpr(nil, lhs, curlsoln, ident2idx, soln)
			cmd = cmd.Op(":=")
			cmd = _CgExpr(cmd, rhs, curlsoln, ident2idx, soln)
			// append this stmt into code space
			code = append(code, cmd)
			// if need conversion, add a conversion stmt
			if ok {
				if curlsoln[0] == 1 && curlsoln[2] == 1 {
					cmd = cg.Id(lhs.Name+"_c").Op(":=").Qual(_pvt, "A2Y").Call(_net, cg.Id(lhs.Name))
					code = append(code, cmd)
				}
				if curlsoln[1] == 1 && curlsoln[3] == 1 {
					cmd = cg.Id(lhs.Name+"_c").Op(":=").Qual(_pvt, "Y2A").Call(_net, cg.Id(lhs.Name))
					code = append(code, cmd)
				}
			}
			codes[len(codes)-1] = code
		case *ast.AssignStmt:
			lhs := n.Lhs[0]
			rhs := n.Rhs[0]
			if n.Tok == token.DEFINE {
				log.Panicf("%#v\n", n)
			}
			lid := identOfExpr(lhs)
			var (
				curlsoln []int
				idx      int
				ok       bool
			)
			if idx, ok = ident2idx[lid]; ok {
				curlsoln = []int{soln[idx], soln[l+idx], soln[2*l+idx], soln[3*l+idx]}
			}
			cmd = _CgExpr(nil, lhs, curlsoln, ident2idx, soln)
			cmd = cmd.Op("=")
			cmd = _CgExpr(cmd, rhs, curlsoln, ident2idx, soln)
			code = append(code, cmd)
			if ok {
				cName := lz.CondInit("",
					lz.Case(curlsoln[0] == 1 && curlsoln[2] == 1, "A2Y"),
					lz.Case(curlsoln[1] == 1 && curlsoln[3] == 1, "Y2A"))
				if cName != "" {
					lhs_p := renameByFn(NewExpr(fset, lhs), func(s string) string { return s + "_c" })
					cmd = _CgExpr(nil, lhs_p, nil, nil, nil).Op("=").Qual(_pvt, cName).Call(_net, _CgExpr(nil, lhs, nil, nil, nil))
					code = append(code, cmd)
				}
			}
			codes[len(codes)-1] = code
		case *ast.ForStmt:
			init := n.Init.(*ast.AssignStmt)
			icmd := _CgExpr(nil, init.Lhs[0], nil, nil, nil)
			icmd = _CgExpr(icmd.Op(":="), init.Rhs[0], nil, nil, nil)
			cond := n.Cond
			ccmd := _CgExpr(nil, cond, nil, nil, nil)
			post := n.Post.(*ast.AssignStmt)
			pcmd := _CgExpr(nil, post.Lhs[0], nil, nil, nil)
			pcmd = pcmd.Op("++")
			codes = append(codes, []cg.Code{})
			astutil.Apply(n.Body, pregen, postgen)
			code = append(code, cg.For(icmd, ccmd, pcmd).Block(codes[len(codes)-1]...))
			codes = codes[:len(codes)-1]
			codes[len(codes)-1] = code
			return false
		case *ast.IfStmt:
		case *ast.ExprStmt:
			cmd = _CgExpr(nil, n.X, nil, nil, nil)
			code = append(code, cmd)
			codes[len(codes)-1] = code
			return false
		}
		return true
	}
	postgen = func(c *astutil.Cursor) bool {
		return true
	}
	// precode
	code := []cg.Code{}
	// Gen Network Phase
	code = append(code, cg.Id("time_net").Op(":=").Qual("time", "Now").Call())
	if ser {
		code = append(code, cg.Id("net").Op(":=").Qual("s3l/mpcfgo/internal/network", "NewServer").Call(cg.Lit(_addr)))
	} else {
		code = append(code, cg.Id("net").Op(":=").Qual("s3l/mpcfgo/internal/network", "NewClient").Call(cg.Lit(_addr)))
	}
	code = append(code, cg.Id("net").Dot("Connect").Call())
	code = append(code, cg.Qual("fmt", "Println").Call(cg.Lit("Net Connect:"), cg.Qual("time", "Since").Call(cg.Id("time_net"))))
	// Gen Precomputed Triple Generation Phase
	code = append(code, cg.Id("time_triple").Op(":=").Qual("time", "Now").Call())
	code = append(code, cg.Qual(_pvt, "TripleFactory").Call(cg.Id("net").Dot("Server")).Dot("SetTriples").Call(
		cg.Qual(_triple, "NewTriples").Call(cg.Id("net"), cg.Qual(_pub, "ZeroInt32"), cg.Lit(100)),
	))
	code = append(code, cg.Qual("fmt", "Println").Call(cg.Lit("Gen Triple:"), cg.Qual("time", "Since").Call(cg.Id("time_triple"))))
	// Gen Online Computation Phase
	code = append(code, cg.Id("time_in_exec").Op(":=").Qual("time", "Now").Call())
	astutil.Apply(fn.Body, pregen, postgen)
	codes[0] = append(_in, codes[0]...)
	codes[0] = append(code, codes[0]...)
	codes[0] = append(codes[0], cg.Qual("fmt", "Println").Call(cg.Lit("In Exec:"), cg.Qual("time", "Since").Call(cg.Id("time_in_exec"))))
	f.Func().Id("main").Params().Block(codes[0]...)
	_in = []cg.Code{}
	return f
}
func __soln2Protocol(soln []int) string {
	if soln[0] == 1 {
		return "AShare"
	} else if soln[1] == 1 {
		return "YShare"
	}
	log.Panicf("%#v\n", soln)
	return ""
}
func _CgExpr(cmd *cg.Statement, e ast.Expr, lsoln []int, ident2idx map[*ast.Ident]int, soln []int) *cg.Statement {
	if e == nil {
		log.Panicln(e)
	}
	switch n := e.(type) {
	case *ast.BasicLit:
		// nil means in a public statement
		if lsoln == nil {
			if cmd == nil {
				return cg.Lit(_CgLit(n))
			}
			return cmd.Lit(_CgLit(n))
		}
		// non-nil means in a secret statement
		if cmd == nil {
			cmd = cg.Qual(_pvt, __soln2Protocol(lsoln))
		} else {
			cmd = cmd.Qual(_pvt, __soln2Protocol(lsoln))
		}
		if _role == 0 {
			cmd = cmd.Dot("New").Call(_net, cg.Qual(_pub, "Zero"+_secLitImplType).Dot("From").Call(cg.Lit(_CgLit(n))))
		} else {
			cmd = cmd.Dot("NewFrom").Call(_net)
		}
		return cmd
	case *ast.Ident:
		if cmd == nil {
			return cg.Id(n.Name)
		}
		return cmd.Id(n.Name)
	case *ast.IndexExpr:
		cmd = _CgExpr(cmd, n.X, nil, nil, nil)
		return cmd.Index(_CgExpr(nil, n.Index, nil, nil, nil))
	case *ast.SliceExpr:
		cmd = _CgExpr(cmd, n.X, lsoln, ident2idx, soln)
		if n.Low == nil {
			return cmd.Index(cg.Empty(), _CgExpr(nil, n.High, nil, nil, nil))
		} else if n.High == nil {
			return cmd.Index(_CgExpr(nil, n.Low, nil, nil, nil), cg.Empty())
		} else {
			return cmd.Index(_CgExpr(nil, n.Low, nil, nil, nil), _CgExpr(nil, n.High, nil, nil, nil))
		}
	case *ast.ArrayType:
		if cmd == nil {
			return _CgExpr(cg.Index(), n.Elt, nil, nil, nil)
		}
		cmd = cmd.Index()
		return _CgExpr(cmd, n.Elt, nil, nil, nil)
	case *ast.ParenExpr:
		if cmd == nil {
			return cg.Parens(_CgExpr(nil, n.X, lsoln, ident2idx, soln))
		}
		return cmd.Parens(_CgExpr(nil, n.X, lsoln, ident2idx, soln))
	case *ast.UnaryExpr:
		// in Sec env
		if lsoln != nil {
			cmd = _CgExpr(cmd, n.X, lsoln, ident2idx, soln)
			return cmd.Dot(token2ImplFn[n.Op]).Call(_net)
		}
		return _CgExpr(cmd.Op(n.Op.String()), n.X, lsoln, ident2idx, soln)
	case *ast.BinaryExpr:
		if lsoln != nil {
			cmd = _CgExpr(cmd, n.X, lsoln, ident2idx, soln).Dot(token2ImplFn[n.Op])
			return cmd.Call(_net, _CgExpr(nil, n.Y, lsoln, ident2idx, soln))
		}
		cmd = _CgExpr(cmd, n.X, lsoln, ident2idx, soln)
		return _CgExpr(cmd.Op(n.Op.String()), n.Y, lsoln, ident2idx, soln)
	case *ast.CallExpr:
		return _CgCallExpr(cmd, n, lsoln, ident2idx, soln)
	default:
		log.Panicf("%#v\n", e)
		return nil
	}
}

func _CgCallExpr(cmd *cg.Statement, e *ast.CallExpr, lsoln []int, ident2idx map[*ast.Ident]int, soln []int) *cg.Statement {
	fnName := e.Fun.(*ast.Ident).Name
	switch fnName {
	// built-in func
	case "make", "copy", "append", "print", "len", "delete":
		// make slice of secret primitive type
		if typArg := e.Args[0]; isArrayType(typArg) && fnName == "make" && isSecArrayType(typArg) {
			index := cg.Index()
			isarr := false
			for typArg, isarr = isEltArrayType(typArg.(*ast.ArrayType)); isarr; typArg, isarr = isEltArrayType(typArg.(*ast.ArrayType)) {
				index = index.Index()
			}
			index.Qual(_pvt, "PvtNum")
			cmd = cmd.Make(index, cg.Lit(_CgLit(e.Args[1].(*ast.BasicLit))))
			return cmd
		}
		// other built-in function call
		cmd = _CgExpr(cmd, e.Fun, lsoln, ident2idx, soln)
		args := make([]cg.Code, len(e.Args))
		for i, v := range e.Args {
			args[i] = _CgExpr(nil, v, nil, nil, nil)
		}
		return cmd.Call(args...)
	// api group 0
	case "b1", "i8", "i16", "i32", "i64", "f32", "f64":
		// scan input from stdin
		inName := _nextIn()
		in := cg.Var().Id(inName).Int()
		_in = append(_in, in)
		_in = append(_in, cg.Qual("fmt", "Scanf").Call(cg.Lit("%d"), cg.Op("&").Id(inName)))
		// distribute secret variable
		cmd = cmd.Qual(_pvt, __soln2Protocol(lsoln)).Dot("New").Call(_net, cg.Qual(_pub, "Zero"+shortType2NormalType(fnName)).Dot("From").Call(cg.Id(inName)))
		return cmd
	// api group 1
	case "b1n", "i8n", "i16n", "i32n", "i64n", "f32n", "f64n":
		inParty := _CgLit(e.Args[0].(*ast.BasicLit)).(int)
		inVarName := _nextIn()
		length := _CgLit(e.Args[1].(*ast.BasicLit)).(int)
		config.DetailLog(func() {
			log.Printf("Party [%#v] %s(%#v, %#v)\n", _role, fnName, inParty, length)
			log.Println(inParty == _role)
		})
		inVarCmd := cg.Id(inVarName).Op(":=").Id("make").Call(cg.Index().Qual(_pvt, "PvtNum"), cg.Lit(int(length)))
		_in = append(_in, inVarCmd)
		iName := _nextIn()
		init := lz.CondInit(nil,
			lz.Case(inParty == _role, cg.For(cg.Id("i").Op(":=").Lit(0), cg.Id("i").Op("<").Lit(int(length)), cg.Id("i").Op("++")).Block(
				cg.Var().Id(iName).Int(),
				cg.Qual("fmt", "Scanf").Call(cg.Lit("%d"), cg.Op("&").Id(iName)),
				cg.Id(inVarName).Index(cg.Id("i")).Op("=").Qual(_pvt, __soln2Protocol(lsoln)).Dot("New").Call(_net, cg.Qual(_pub, "Zero"+shortType2NormalType(fnName[:len(fnName)-1])).Dot("From").Call(cg.Id(iName))),
			)),
			lz.Case(inParty != _role, cg.For(cg.Id("i").Op(":=").Lit(0), cg.Id("i").Op("<").Lit(int(length)), cg.Id("i").Op("++")).Block(
				cg.Id(inVarName).Index(cg.Id("i")).Op("=").Qual(_pvt, __soln2Protocol(lsoln)).Dot("NewFrom").Call(_net),
			)),
		)
		_in = append(_in, init)
		// if inParty == _role {
		// 	_in = append(_in,
		// 		cg.For(cg.Id("i").Op(":=").Lit(0), cg.Id("i").Op("<").Lit(int(length)), cg.Id("i").Op("++")).Block(
		// 			cg.Id(inVarName).Index(cg.Id("i")).Op("=").Qual(_pvt, __soln2Protocol(lsoln)).Dot("NewFrom").Call(_net),
		// 		))
		// } else {
		// 	_in = append(_in,
		// 		cg.For(cg.Id("i").Op(":=").Lit(0), cg.Id("i").Op("<").Lit(int(length)), cg.Id("i").Op("++")).Block(
		// 			cg.Id(inVarName).Index(cg.Id("i")).Op("=").Qual(_pvt, __soln2Protocol(lsoln)).Dot("NewFrom").Call(_net),
		// 		))
		// }

		return cmd.Id(inVarName)
	// api group 2
	case "openb1", "openi8", "openi16", "openi32", "openi64", "openf32", "openf64",
		"openb1n", "openi8n", "openi16n", "openi32n", "openi64n", "openf32n", "openf64n":
		opened := _CgExpr(nil, e.Args[0], nil, nil, nil)
		if strings.HasSuffix(fnName, "n") {
			return cmd.Qual(_pvt, "DeclassifyN").Call(_net, opened)
		}
		return cmd.Qual(_pvt, "Declassify").Call(_net, opened)
	// api group 4
	case "Addi8", "Addi16", "Addi32", "Addi64", "Addf32", "Addf64",
		"Subi8", "Subi16", "Subi32", "Subi64", "Subf32", "Subf64",
		"Muli8", "Muli16", "Muli32", "Muli64", "Mulf32", "Mulf64",
		"Divi8", "Divi16", "Divi32", "Divi64", "Divf32", "Divf64",
		"Gti8", "Gti16", "Gti32", "Gti64", "Gtf32", "Gtf64",
		"Lti8", "Lti16", "Lti32", "Lti64", "Ltf32", "Ltf64",
		"Eqi8", "Eqi16", "Eqi32", "Eqi64", "Eqf32", "Eqf64",
		"Shri8", "Shri16", "Shri32", "Shri64",
		"Shli8", "Shli16", "Shli32", "Shli64",
		"Eqb1", "Andb1", "Orb1":
		return cmd.Id("apifn")
	// api group 5
	case "Muxb1", "Muxi8", "Muxi16", "Muxi32", "Muxi64", "Muxf32", "Muxf64":
		return cmd.Id("muxfn")
	default:
		return cmd.Id("ThereIsCall0")
	}
}

func _CgLit(lit *ast.BasicLit) interface{} {
	switch lit.Kind {
	case token.INT:
		i, err := strconv.ParseInt(lit.Value, 10, 64)
		if err != nil {
			log.Panicf("%#v, %#v", lit, err)
		}
		return int(i)
	case token.FLOAT:
		i, err := strconv.ParseFloat(lit.Value, 64)
		if err != nil {
			log.Panicf("%#v, %#v", lit, err)
		}
		return float32(i)
	}
	return nil
}

/*
*

	Help Function

*
*/
var secPrimTypeSet = []string{
	"sbool",
	"sint8",
	"sint16",
	"sint32",
	"sint64",
	"sfloat32",
	"sfloat64",
}

var secImplTypeSet = []string{
	"Bool",
	"Int8",
	"Int16",
	"Int32",
	"Int64",
	"Float32",
	"Float64",
}

var token2ImplFn = map[token.Token]string{
	token.ADD:  "Add",
	token.SUB:  "Sub",
	token.MUL:  "Mul",
	token.QUO:  "Div",
	token.LAND: "And",
	token.LOR:  "Or",
	token.NOT:  "Not",
	token.EQL:  "Eq",
	token.GTR:  "Gt",
	token.LSS:  "Lt",
}

func isSecPrimType(s string) bool {
	for _, v := range secPrimTypeSet {
		if v == s {
			return true
		}
	}
	for _, v := range secPrimTypeSet {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}

func secPrim2Impl(s string) string {
	for i, v := range secPrimTypeSet {
		if v == s {
			return secImplTypeSet[i]
		}
	}
	log.Panicf("%#v\n", s)
	return ""
}
func isSecArrayType(expr ast.Expr) bool {
	switch n := expr.(type) {
	case *ast.ArrayType:
		return isSecArrayType(n.Elt)
	case *ast.Ident:
		return isSecPrimType(n.Name)
	}
	log.Panicf("%#v\n", expr)
	return false
}

func isArrayType(expr ast.Expr) bool {
	_, ok := expr.(*ast.ArrayType)
	return ok
}

func isEltArrayType(expr *ast.ArrayType) (elt ast.Expr, isArray bool) {
	elt, isArray = expr.Elt.(*ast.ArrayType)
	return
}

func pruneInfoType(s string) string {
	sp := strings.Split(s, ".")
	s = sp[len(sp)-1]
	return s
}
func secImplTypeOfExpr(info *types.Info, expr ast.Expr) string {
	id := identOfExpr(expr)
	typ := info.TypeOf(id)
	if typ == nil {
		log.Panicf("%#v\n", expr)
	}
	secPrimType := pruneInfoType(typ.String())
	return secPrim2Impl(secPrimType)
}

// Solve Ident, IndexExpr, SliceExpr
func exprOfExpr(expr ast.Expr) ast.Expr {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return expr
	case *ast.Ident:
		return expr
	case *ast.IndexExpr:
		return exprOfExpr(x.X)
	case *ast.SliceExpr:
		return exprOfExpr(x.X)
	default:
		log.Panicf("unexpected expr: %#v", expr)
		return nil
	}
}

func identOfExpr(expr ast.Expr) *ast.Ident {
	switch x := expr.(type) {
	case *ast.Ident:
		return x
	case *ast.IndexExpr:
		return identOfExpr(x.X)
	case *ast.SliceExpr:
		return identOfExpr(x.X)
	default:
		log.Panicf("unexpected expr: %#v", expr)
		return nil
	}
}
func fmtIdentOrLit(fset *token.FileSet, atomic ast.Expr) string {
	switch d := atomic.(type) {
	case *ast.Ident:
		pos := fset.Position(d.Pos())
		return fmt.Sprintf("%s.%v:%v", d.Name, pos.Line, pos.Column)
	case *ast.BasicLit:
		return d.Value
	}
	log.Panicf("%#v\n", atomic)
	return ""
}

func isBasicLit(m ast.Expr) bool {
	if _, ok := m.(*ast.BasicLit); ok {
		return ok
	}
	return false
}

func isBuiltin(i *ast.Ident) bool {
	return i.Obj == nil
}

func tok2JsonKey(tok token.Token) string {
	switch tok {
	case token.ADD:
		return "Add"
	case token.SUB:
		return "Sub"
	case token.MUL:
		return "Mul"
	case token.QUO:
		return "Div"
	case token.LAND:
		return "And"
	case token.LOR:
		return "Or"
	case token.NOT:
		return "Not"
	case token.EQL:
		return "Eq"
	case token.GTR:
		return "Gt"
	case token.LSS:
		return "Lt"
	case token.ASSIGN:
		return "New"
	case NEW:
		return "New"
	case NEWN:
		return "NewN"
	default:
		log.Panicf("%#v", tok)
		return ""
	}
}

func shortType2NormalType(s string) string {
	switch s {
	case "b1":
		return "Bool"
	case "i8":
		return "Int8"
	case "i16":
		return "Int16"
	case "i32":
		return "Int32"
	case "i64":
		return "Int64"
	case "f32":
		return "Float32"
	case "f64":
		return "Float64"
	default:
		log.Panic(s)
		return ""
	}
}

func renameByFn(e ast.Expr, f func(string) string) ast.Expr {
	switch d := e.(type) {
	case *ast.Ident:
		d.Name = f(d.Name)
		return d
	case *ast.IndexExpr:
		d.X = renameByFn(d.X, f)
		return d
	case *ast.SliceExpr:
		d.X = renameByFn(d.X, f)
		return d
	default:
		log.Panicf("%#v\n", e)
	}
	return nil
}

func identToIdx(ident2idx map[*ast.Ident]int, ident *ast.Ident) int {
	idx, ok := ident2idx[ident]
	if !ok {
		for i, v := range ident2idx {
			if ident.Name == i.Name && ident.NamePos == i.NamePos {
				idx = v
				break
			}
		}
	}
	return idx
}
