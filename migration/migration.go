package main

import (
	"flag"
	"os"
	"log"
)

func main() {
	flag.Parse()

	dsn := flag.Arg(0)
	scripts := flag.Arg(1)

	if dsn == "" || scripts == "" {
		log.Println(`Usage: migrator "dsn" "scriptsDirectory"`)
		os.Exit(1)
	}

	err := Migrate(dsn, scripts)

	if err != nil {
		log.Println("supplied dsn: ", dsn)
		log.Println("supplied scriptsDirectory: ", scripts)
		log.Fatal(err.Error())
	}
}
