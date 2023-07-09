package iast

import (
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DoMain(dirpath string) error {
	// const propagate
	fset, pkg := ParseDir(dirpath)
	ConstPropagate(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "0_constp")

	// expand forloop's 1st and last round
	cpdir := SiblingDirOf(dirpath, "0_constp")
	fset, pkg = ParseDir(cpdir)
	info, err := CheckInfo(fset, FilesOfPkg(pkg)...)
	if err != nil {
		log.Printf("InfoCheck %v", err)
	}
	ExpandLoopPkg(fset, info, pkg)
	ConstPropagate(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "1_expand")

	// delete unused var (generated due to expand)
	xpd := SiblingDirOf(cpdir, "1_expand")
	fset, pkg = ParseDir(xpd)
	DelateUnusedGenDeclPkg(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "2_unused")

	// inline, make main decl has no fucall, delete all other func decl
	uusd := SiblingDirOf(xpd, "2_unused")
	fset, pkg = ParseDir(uusd)
	InlinerPkg(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "3_inlined")

	// const var propagation
	inlined := SiblingDirOf(uusd, "3_inlined")
	fset, pkg = ParseDir(inlined)
	VarPropagate(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "5_varp")
	inlined = SiblingDirOf(cpdir, "5_varp")
	fset, pkg = ParseDir(inlined)
	DelateUnusedGenDeclPkg(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "5_varp")
	inlined = SiblingDirOf(uusd, "5_varp")
	fset, pkg = ParseDir(inlined)
	VarPropagate(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "5_varp")
	inlined = SiblingDirOf(uusd, "5_varp")
	fset, pkg = ParseDir(inlined)
	VarPropagate(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "5_varp")
	inlined = SiblingDirOf(cpdir, "5_varp")
	fset, pkg = ParseDir(inlined)
	DelateUnusedGenDeclPkg(fset, pkg)
	WriteTo(fset, FilesOfPkg(pkg), "5_varp")
	// do def-use analysis
	varped := SiblingDirOf(uusd, "5_varp")
	fset, pkg = ParseDir(varped)
	info, err = CheckInfo(fset, FilesOfPkg(pkg)...)
	if err != nil {
		log.Printf("InfoCheck %v", err)
	}
	ar := AnalyzeDefusePkg(fset, info, pkg)

	// do optimize of mixed protocols
	ident2idx, soln := MP_Optimize(fset, info, pkg, ar)
	//WriteTo(fset, FilesOfPkg(pkg), "3_defuse")

	// do code gen
	gen := SiblingDirOf(uusd, "4_codegen")
	CodeGen(fset, info, pkg, ident2idx, soln, gen)
	
	// fs := []*ast.File{}
	// for fname, file := range pkg.Files {
	// 	fs = append(fs, file)
	// 	log.Printf("File [%v] in Pkg [%v]\n", fname, pkg.Name)
	// }
	// log.Printf("Analyze Pkg [%v]\n", pkg.Name)
	// if iast.DoMain(fset, fs) != nil {
	// 	log.Fatal("DoMain", err)
	// }
	// const propagate
	//doTranslate(f)
	// Get TypeInfo
	// info, err := CheckInfo(fset, fs...)
	// if err != nil {
	// 	log.Panicf("CheckInfo %v\n", err)
	// }
	// // Expand Loop in 1 depth
	// for _, file := range fs {
	// 	if !(file.Name.Name == "template.go") {
	// 		ExpandLoopFile(fset, info, file)
	// 	}
	// }
	// // Write to unroll
	// WriteTo(fset, fs, "unroll")
	// for _, file := range fs {
	// 	if !(file.Name.Name == "template.go") {
	// 		ToSsaFile(fset, file)
	// 	}
	// }
	// // Make To-SSA translation
	// for _, file := range fs {
	// 	if !(file.Name.Name == "template.go") {
	// 		ToSsaFile(fset, file)
	// 	}
	// }

	// Make Def-Use Analysis
	// info, err := CheckInfo(fset, fs...)
	// if err != nil {
	// 	log.Fatal("CheckInfo ", err)
	// }
	// for _, file := range fs {
	// 	if !(file.Name.Name == "template.go") {
	// 		doDefuseAnalyzeFile(fset, file, info)
	// 	}
	// }

	// Write out new file

	return nil
}

func FilesOfPkg(pkg *ast.Package) []*ast.File {
	fs := []*ast.File{}
	for _, v := range pkg.Files {
		fs = append(fs, v)
	}
	return fs
}
func ParseDir(dirpath string) (*token.FileSet, *ast.Package) {
	log.Printf("Parse Dir [%s]\n", dirpath)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dirpath, nil, 0)
	if err != nil {
		log.Fatal("ParseDir ", err)
	}
	var pkg *ast.Package
	for _, v := range pkgs {
		pkg = v
		break
	}
	if pkg == nil {
		log.Panicf("Zero Pkg in %v\n", dirpath)
	}
	return fset, pkg
}

func WriteTo(fset *token.FileSet, files []*ast.File, dirName string) {
	nfset := token.NewFileSet()
	fullFileName := fset.Position(files[0].Pos()).Filename
	dirPath, _ := filepath.Split(fullFileName)
	if strings.HasSuffix(dirPath, string(filepath.Separator)) {
		dirPath = dirPath[:len(dirPath)-1]
	}
	parentDirPath := filepath.Dir(dirPath)
	newDirPath := filepath.Join(parentDirPath, dirName)
	err := os.Mkdir(newDirPath, 0777)
	if err != nil && !os.IsExist(err) {
		log.Panicf("Mkdir [%v] Failed\n", err)
	} else {
		log.Printf("Mkdir [%v] Success\n", filepath.Join(parentDirPath, dirName))
	}
	for _, file := range files {
		_, fileName := filepath.Split(fset.Position(file.Pos()).Filename)
		log.Printf("Write [%s] into [%s]\n", fileName, newDirPath)
		filepath := filepath.Join(newDirPath, fileName)
		os.Remove(filepath)
		dst, err := os.Create(filepath)
		if err != nil {
			log.Panicf("Touch [%v] Failed\n", fileName)
		}
		err = format.Node(dst, nfset, file)
		if err != nil {
			log.Panicf("Format [%v] Failed\n", fileName)
		}
	}
}

func SiblingDirOf(cur string, dst string) string {
	if strings.HasSuffix(cur, string(filepath.Separator)) {
		cur = cur[:len(cur)-1]
	}
	parentDir := filepath.Dir(cur)
	return filepath.Join(parentDir, dst)
}

func CheckInfo(fset *token.FileSet, files ...*ast.File) (*types.Info, error) {
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types:     map[ast.Expr]types.TypeAndValue{},
		Defs:      make(map[*ast.Ident]types.Object),
		Uses:      make(map[*ast.Ident]types.Object),
		Scopes:    map[ast.Node]*types.Scope{},
		InitOrder: []*types.Initializer{},
	}
	_, err := conf.Check("cryptonet", fset, files, info)
	if err != nil {
		return nil, err // type error
	}
	// for id, obj := range info.Uses {
	// 	pos := fset.Position(id.Pos())
	// 	if !strings.HasSuffix(pos.Filename, "template.go") {
	// 		fmt.Printf("%s: %q uses %v\n",
	// 			fset.Position(id.Pos()).Filename, id.Name, obj)
	// 	}
	// }
	return info, nil
}
