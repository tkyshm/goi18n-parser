[![Build Status](https://travis-ci.org/tkyshm/goi18n-parser.svg?branch=master)](https://travis-ci.org/tkyshm/goi18n-parser)

# goi18n-parser

Note: Please check [go-i18n](https://github.com/nicksnyder/go-i18n) before to use this package.

## Usage

very simple!
```go
package main

import (
    "fmt"
    "goi18np"
    "encoding/json"
)

func main() {
    a := goi18np.Analyzer{
        Debug: true,
    }
    result := a.AnalyzeFromFile("path/to/code.go")

    out, err := json.Marshal(result)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(out))
}

```

## Analyzer

Analyzer has fields:

name     | type            |
-------- | --------------- | ----------------------------------------------
FuncName | string          | Analyze target function name
Debug    | bool            | If true, exec `ast.Print` for go source codes
Records  | []I18NRecord    | Analysis result with `FuncName`

you can analyze such as the following code:

```go
fmt.Println(T("message_uniq_key"))
fmt.Println(SomeStruct.T("message_uniq_key"))
```

If you want to parse other function names, please fill `FuncName`:

```go
a := Analyzer{
    FuncName: "Translation",
}
```
