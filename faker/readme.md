# Faker

- Words
- Sentences
- Date
- Number
- Email
- Name

```golang
package main

import (
    "fmt"
    "github.com/crit/critical-go/faker"
)

func main() {
    words := faker.Words()

    fmt.Printf("Words: %+v", words)
}
```

