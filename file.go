package main

import (
"go/ast"
"go/parser"
"go/token"
"log"
"fmt"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "hello.go", nil, 0)
	if err != nil {
	log.Fatal(err)
	}
	fmt.Print("Imports:")
	for _, i := range node.Imports {
	    fmt.Println(i.Path.Value)
	}

	fmt.Println("Comments:")
	for _, c := range node.Comments {
	    fmt.Print(c.Text())
	}

	fmt.Print("Functions:")
	for _, f := range node.Decls {
	    fn, ok := f.(*ast.FuncDecl)
	    if !ok {
	        continue
	    }
	    fmt.Println(fn.Name.Name)
	}


	//ast.Print(fset, f)
}
