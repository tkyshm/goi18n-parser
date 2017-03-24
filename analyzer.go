package goi18np

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"sort"
	"strings"
)

// Const values
const (
	AnalyzerFuncName = "T"
	MaxDepth         = 5
)

// Analyzer struct
type Analyzer struct {
	FuncName string // Function name
	Debug    bool
	Records  []I18NRecord
}

// Name returns the function name that is used to analyze go source code
func (da Analyzer) Name() string {
	if da.FuncName != "" {
		return da.FuncName
	}
	return AnalyzerFuncName
}

// I18NRecord has `id` and `translation` field
type I18NRecord struct {
	ID          string `json:"id"`
	Translation string `json:"translation"`
}

// I18NRecords is slice of I18NRecord
type I18NRecords []I18NRecord

// SortByID sorts I18NRecords by ID
func (rs I18NRecords) SortByID() {
	sort.Slice(rs, func(i, j int) bool {
		if strings.Compare(rs[i].ID, rs[j].ID) == -1 {
			return true
		}
		return false
	})
}

// AnalyzeFromFile analyzes a go file
func (da *Analyzer) AnalyzeFromFile(filename string) []I18NRecord {
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

// AnalyzeFromFiles analyzes multiple go source files
func (da *Analyzer) AnalyzeFromFiles(files []string) []I18NRecord {
	for _, filename := range files {
		da.AnalyzeFromFile(filename)
	}
	return da.Records
}

// SaveJSON saves JSON based on go-i18np format
func (da Analyzer) SaveJSON(path string) error {
	out, err := json.Marshal(da.Records)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, out, 0644)
	if err != nil {
		return err
	}

	return nil
}

// SaveJSONIndent saves JSON based on go-i18np format
func (da Analyzer) SaveJSONIndent(path, prefix, indent string) error {
	out, err := json.MarshalIndent(da.Records, prefix, indent)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, out, 0644)
	if err != nil {
		return err
	}

	return nil
}

func containsID(id string, rs []I18NRecord) bool {
	for _, r := range rs {
		if r.ID == id {
			return true
		}
	}
	return false
}
