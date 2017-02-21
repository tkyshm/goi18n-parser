package goi18np

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

const (
	DefaultAnalyzerFuncName = "T"
	MaxDepth                = 3
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
			ident := traversalToIdent(x.Fun, 0)
			if ident == nil {
				return true
			}

			fn := ident.Name
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
		return true
	})

	return da.Records
}

func traversalToIdent(n interface{}, depth int) *ast.Ident {
	switch x := n.(type) {
	case *ast.Ident:
		return x
	case *ast.SelectorExpr:
		if depth < MaxDepth {
			return traversalToIdent(x.Sel, depth+1)
		}
	default:
		return nil
	}
	return nil
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
