# Critical Go

My support libraries for Go programming language projects.

## Env

```go
package main

import "github.com/crit/critical-go/env"

func main() {
    host := env.GetString("DB_HOST", "localhost")
    
    // if DB_HOST was set in the environment, then host would be set to its
    // value. Otherwise host will be set to "localhost".
}
```

## Errors

[Code Walk-through](http://critrussell.com/posts/application-specific-error-handling/)

```go
package users

import "github.com/crit/critical-go/errors"

// imports, types, other funcs, etc ...

func Get(id string) (user User, err error) {
	user.ID, err = tokens.ToID(id) // decode id token

	if err != nil {
		// ok, this is a client error. cool
		return user, errors.Error(http.StatusBadRequest, err)
		// or we could log `err` and use the following to keep
		// info from the client:
		// return user, errors.ErrorString(http.StatusNotFound, "user not found")
	}

	err = db.Find(&user) // look up in repo
	// assuming a project specific implementation of db.Find
	// we may be getting a errors.Err struct so we can just return it.
	// Even if we are not, our implementation should handle being
	// passed native errors.

	return user, err
}
```

## Logger

[Code Walk-through](http://critrussell.com/posts/go-structured-logger/)

Standard usage:

```go
// project/cmd/product/main.go
package main

import (
  "github.com/crit/critical-go/logger"  
)

func main() {
  log := logger.New("product", logger.Writer{})
  
  log.With(logger.Data{"hello": "world"}).Info("example")
  // running the program after building (or possibly from `go run main.go`) will yield 
  // this in the cli:
  // {..., "app":"product","level":"info","msg":"example","data":{"hello":"world"}, ...}
}
```

Post to remote server usage:

```go
// project/cmd/product/main.go
package main

import (
  "github.com/crit/critical-go/logger"  
)

func main() {
  log := logger.New("product", logger.NewPostWriter("https://example/injest", logger.Writer{}))
  
  log.With(logger.Data{"hello": "world"}).Info("example")
  // running the program after building (or possibly from `go run main.go`) will yield 
  // this in the cli:
  // {..., "app":"product","level":"info","msg":"example","data":{"hello":"world"}, ...}
  // and this log value will be posted to the remote server (if the network allows).
}
```

## Password

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
