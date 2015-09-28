# logger

- Debug
- Debugf
- Debugln
- Info
- Infof
- Infoln
- Print
- Printf
- Println
- Warn
- Warnf
- Warnln
- Error
- Errorf
- Errorln
- Panic
- Panicf
- Panicln
- Fatal
- Fatalf
- Fatalln

```go
package main

input "github.com/crit/critical-go/logger"

var Log = logger.Basic()

func main() {
    Log.Infoln("Logging to stdout")
    Log.Warnln("Warning to stdout")
    Log.Errorln("Error to stdout")
}
```
