# database

- Mock
- MSSQL
- MySQL

```go
package main

import (
    "fmt"
    "github.com/crit/critical-go/database"
)

var DB = database.New(database.Config{
    Driver: "mysql",
    DSN:    "username:password@tcp(domain:port)/database_name?charset=utf8&parseTime=true",
    Idle:   "10",
    Max:    "100",
})

type User struct {
    ID int
    Name string
}

func main() {
    var user User

    db := DB.Connection()

    db.Where("id = ?", 1).First(&user)

    fmt.Printf("Name: %s" user.Name)
}
```
