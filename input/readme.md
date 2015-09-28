# input

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/crit/critical-go/input"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        in := input.New(req)

        fmt.Fprintf(w, "Input: %+v", in.All())
    })

    http.ListenAndServe(":8080", nil)
}
```
