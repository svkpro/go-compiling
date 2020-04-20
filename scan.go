package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"log"
)

func main() {
	src, err := ioutil.ReadFile("main.go")

	if err != nil{
		log.Fatal(err)
	}

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("main.go", fset.Base(), len(src))
	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}