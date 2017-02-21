# goi18n-parser
This package is to create go-i18n JSON.

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
    a := goi18np.Analyzer{}
    result := AnalyzeFromFile
    out, err := json.Marshal(result)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(out))
}

```
