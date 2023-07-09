package iast

import (
	"go/ast"
	"go/token"
	"go/types"
	"log"
)

func NewIdent(ident *ast.Ident) *ast.Ident {
	if ident == nil {
		log.Panicf("NewIdent: ident is nil")
	}
	nIdent := ast.NewIdent(ident.Name)
	nIdent.Obj = ident.Obj
	return nIdent
}

func NewDeclFromAssign(fset *token.FileSet, info *types.Info, m *ast.AssignStmt, suffix string) *ast.DeclStmt {
	if len(m.Lhs) != 1 || len(m.Rhs) != 1 {
		log.Panicf("NewDeclFromAssign: len(m.Lhs) != 1 || len(m.Rhs) != 1")
	}
	ident := m.Lhs[0].(*ast.Ident)
	nIdent := ast.NewIdent(ident.Name + suffix)
	newDecl := &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names:  []*ast.Ident{nIdent},
					Type:   ast.NewIdent("int"),
					Values: []ast.Expr{NewExpr(fset, m.Rhs[0])},
				},
			},
		},
	}
	return newDecl
}
func NewDeclStmt(fset *token.FileSet, scope *ast.Scope, m ast.Stmt, name string) ast.Stmt {
	declStmt := m.(*ast.DeclStmt)
	genDecl := declStmt.Decl.(*ast.GenDecl)
	valSpec := genDecl.Specs[0].(*ast.ValueSpec)
	ident := valSpec.Names[0]
	obj := scope.Lookup(ident.Name)

	nIdent := ast.NewIdent(name)
	nObj := ast.NewObj(obj.Kind, name)
	nIdent.Obj = nObj
	if scope.Insert(nObj) != nil {
		log.Panicf("duplicate object %s", name)
	}
	newDecl := &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names:  []*ast.Ident{nIdent},
					Type:   nil,
					Values: []ast.Expr{NewExpr(fset, valSpec.Values[0])},
				},
			},
		},
	}
	nObj.Decl = newDecl
	return newDecl
}
func NewExpr(fset *token.FileSet, m ast.Expr) ast.Expr {
	if m == nil {
		return nil
	}
	switch d := m.(type) {
	case *ast.Ident:
		return NewIdent(d)
	case *ast.BasicLit:
		return &ast.BasicLit{
			Kind:  d.Kind,
			Value: d.Value,
		}
	case *ast.ParenExpr:
		return &ast.ParenExpr{
			X: NewExpr(fset, d.X),
		}
	case *ast.UnaryExpr:
		return &ast.UnaryExpr{
			Op: d.Op,
			X:  NewExpr(fset, d.X),
		}
	case *ast.BinaryExpr:
		return &ast.BinaryExpr{
			X:  NewExpr(fset, d.X),
			Op: d.Op,
			Y:  NewExpr(fset, d.Y),
		}
	case *ast.IndexExpr:
		return &ast.IndexExpr{
			X:     NewExpr(fset, d.X),
			Index: NewExpr(fset, d.Index),
		}
	case *ast.SliceExpr:
		return &ast.SliceExpr{
			X:      NewExpr(fset, d.X),
			Low:    NewExpr(fset, d.Low),
			High:   NewExpr(fset, d.High),
			Max:    NewExpr(fset, d.Max),
			Slice3: d.Slice3,
		}
	case *ast.CallExpr:
		args := []ast.Expr{}
		for _, v := range d.Args {
			args = append(args, NewExpr(fset, v))
		}
		return &ast.CallExpr{
			Fun:  NewExpr(fset, d.Fun),
			Args: args,
		}
	case *ast.ArrayType:
		return &ast.ArrayType{
			Elt: NewExpr(fset, d.Elt),
		}
	default:
		log.Panicf("%#v", d)
		return nil
	}
}

func NewStmt(fset *token.FileSet, m ast.Stmt) ast.Stmt {
	switch d := m.(type) {
	case *ast.DeclStmt:
		decl := d.Decl.(*ast.GenDecl)
		valSpec := decl.Specs[0].(*ast.ValueSpec)
		return &ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: decl.Tok,
				Specs: []ast.Spec{&ast.ValueSpec{
					Names:  []*ast.Ident{NewIdent(valSpec.Names[0])},
					Type:   valSpec.Type,
					Values: []ast.Expr{NewExpr(fset, valSpec.Values[0])},
				}},
			},
		}
	case *ast.AssignStmt:
		lhs := []ast.Expr{}
		for _, v := range d.Lhs {
			lhs = append(lhs, NewExpr(fset, v))

		}
		rhs := []ast.Expr{}
		for _, v := range d.Rhs {
			rhs = append(rhs, NewExpr(fset, v))
		}
		return &ast.AssignStmt{
			Lhs:    lhs,
			Tok:    d.Tok,
			TokPos: d.TokPos,
			Rhs:    rhs,
		}
	case *ast.ExprStmt:
		return &ast.ExprStmt{
			X: NewExpr(fset, d.X),
		}
	case *ast.ForStmt:
		return &ast.ForStmt{
			Init: NewStmt(fset, d.Init),
			Cond: NewExpr(fset, d.Cond),
			Post: NewStmt(fset, d.Post),
			Body: NewStmt(fset, d.Body).(*ast.BlockStmt),
		}
	case *ast.BlockStmt:
		list := []ast.Stmt{}
		for _, v := range d.List {
			list = append(list, NewStmt(fset, v))
		}
		return &ast.BlockStmt{
			List: list,
		}
	case *ast.ReturnStmt:
		results := []ast.Expr{}
		for _, v := range d.Results {
			results = append(results, NewExpr(fset, v))
		}
		return &ast.ReturnStmt{
			Results: results,
		}
	default:
		log.Panicf("%#v", d)
	}
	return nil
}

func Assign2Decl(fset *token.FileSet, m ast.Stmt) ast.Stmt {
	switch d := m.(type) {
	case *ast.AssignStmt:
		lhs := []*ast.Ident{}
		for _, v := range d.Lhs {
			lhs = append(lhs, NewExpr(fset, v).(*ast.Ident))

		}
		rhs := []ast.Expr{}
		for _, v := range d.Rhs {
			rhs = append(rhs, NewExpr(fset, v))
		}
		return &ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{&ast.ValueSpec{
					Names:  lhs,
					Type:   ast.NewIdent("int"),
					Values: rhs,
				}},
			},
		}
	default:
		log.Panicf("%#v", d)
	}
	return nil
}
