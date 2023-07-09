package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	//fnValue := reflect.MakeFunc(reflect.TypeOf(true), func(args []reflect.Value) (results []reflect.Value) { return nil })
	//fnPtr := fnValue.Pointer()
	reflect.ValueOf(1)
	reflect.ValueOf(1)
	//fnImpl := unsafe.Pointer(fnPtr)
	reflect.ValueOf(1)
	// expr, err := parser.ParseExpr()
	// fmt.Printf("expr: %#v\n%v", expr, err)
}

func TestParser(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "/home/fanyx/mine/FeMPC/example/biometric/native/biometric.go", nil, parser.AllErrors)
	fmt.Printf("f: %v\n", f)
	for _, v := range f.Decls {
		if st, ok := v.(*ast.GenDecl); ok {
			fmt.Printf("st.Tok: %v\n", st.Tok)
			for _, vj := range st.Specs {
				switch spec := vj.(type) {
				case *ast.ImportSpec:
					continue
				case *ast.ValueSpec:
					fmt.Printf("spec.Names: %v\n", spec.Names)
					fmt.Printf("spec.Type: %v\n", spec.Type)
					fmt.Printf("spec.Values: %v\n", spec.Values)
				case *ast.TypeSpec:
				}
			}
		} else if st, ok := v.(*ast.FuncDecl); ok {
			fmt.Printf("FuncDecl: %#v\n", st)
			for _, vj := range st.Body.List {
				fmt.Printf("vj: %#v\n", vj)
			}
		}
	}
	fmt.Printf("err: %v\n", err)
}
