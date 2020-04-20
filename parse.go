package main

import (
"go/ast"
"go/parser"
"go/token"
	"io/ioutil"
	"log"
)

func main() {
	src, err := ioutil.ReadFile("main.go")

	if err != nil{
		log.Fatal(err)
	}

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}
