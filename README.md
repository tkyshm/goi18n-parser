[![Build Status](https://travis-ci.org/tkyshm/goi18n-parser.svg?branch=master)](https://travis-ci.org/tkyshm/goi18n-parser)

# goi18n-parser
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
    a := goi18np.DefaultAnalyzer{
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
