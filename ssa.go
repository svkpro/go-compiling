package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"io/ioutil"
	"os"
)

func main() {
	src, err := ioutil.ReadFile("main.go")

	// Parse the source files.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		fmt.Print(err)
		return
	}
	files := []*ast.File{f}

	// Create the type-checker's package.
	pkg := types.NewPackage("main", "")

	// Type-check the package, load dependencies.
	// Create and build the SSA program.
	main, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		fmt.Print(err) // type error in some package
		return
	}

	// Print out the package.
	main.WriteTo(os.Stdout)

	// Print out the package-level functions.
	main.Func("init").WriteTo(os.Stdout)
	main.Func("main").WriteTo(os.Stdout)
}