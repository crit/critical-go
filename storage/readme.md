# storage

- Local (in memory)
- Folder (file system folder)
- S3 (Amazon's S3)
- Redis

```golang
package main

import (
    "fmt"
    "github.com/crit/critical-go/storage"
)

func main() {
    store := storage.Local()

    store.Put("test", []byte("testing value"))

    fmt.Printf("Test: %s" store.Get("test"))
}
```
