# password

```go
package main

import (
    "fmt"
    "github.com/crit/critical-go/password"
)

func main() {
    hashword := password.Hash("TestPassword3!")

    fmt.Println("Hashed: " + hashword)
}
```
