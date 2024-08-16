package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "gobreakselectinfor",
	Doc:  "Checks that using break statement inside select inside for loop",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		ast.Inspect(funcDecl.Body, func(stmt ast.Node) bool {
			if forStmt, ok := stmt.(*ast.ForStmt); ok {
				ast.Inspect(forStmt.Body, func(stmt ast.Node) bool {
					if selStmt, ok := stmt.(*ast.SelectStmt); ok {
						ast.Inspect(selStmt.Body, func(stmt ast.Node) bool {
							if brkStmt, ok := stmt.(*ast.BranchStmt); ok && brkStmt.Tok == token.BREAK {
								pass.Reportf(stmt.Pos(), "break statement inside select statement inside for loop")
								return true
							}
							return true
						})
					}
					return true
				})
			}
			return true
		})

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
