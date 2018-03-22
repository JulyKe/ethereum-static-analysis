package main

import (
"go/ast"
"go/parser"
"go/token"
"log"
"fmt"
"strings"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "hello.go", nil, 0)
	if err != nil {
	log.Fatal(err)
	}
	var v visitor
	ast.Walk(v, node)
}

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

