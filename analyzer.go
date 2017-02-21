package goi18np

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

const (
	DefaultAnalyzerFuncName = "T"
)

type ExecMode int

type Analyzer interface {
	Name() string
}

type DefaultAnalyzer struct {
	Fname   string // Function name
	Debug   bool
	Records []I18NRecord
}

func (da DefaultAnalyzer) Name() string {
	if da.Fname != "" {
		return da.Fname
	}
	return DefaultAnalyzerFuncName
}

type I18NRecord struct {
	ID          string `json:"id"`
	Translation string `json:"translation"`
}

func (da *DefaultAnalyzer) AnalyzeFromFile(filename string) []I18NRecord {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}

	if da.Debug {
		ast.Print(fset, f)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			switch y := x.Fun.(type) {
			case *ast.Ident:
				fn := y.Name
				if fn == "T" {
					key := strings.Trim(x.Args[0].(*ast.BasicLit).Value, "\"")
					r := I18NRecord{
						ID: key,
					}
					if !containsID(key, da.Records) {
						da.Records = append(da.Records, r)
					}
				}
			}
		}
		return true
	})

	return da.Records
}

func (da *DefaultAnalyzer) AnalyzeFromFiles(files []string) []I18NRecord {
	for _, filename := range files {
		da.AnalyzeFromFile(filename)
	}
	return da.Records
}

func containsID(id string, rs []I18NRecord) bool {
	for _, r := range rs {
		if r.ID == id {
			return true
		}
	}
	return false
}
