package main

import (
	"flag"
	"os"
	"log"
)

var BUILD = ""

func main() {
	flag.Parse()

	dsn := flag.Arg(0)
	scripts := flag.Arg(1)

	if dsn == "" || scripts == "" {
		log.Println(`Usage: migrator "dsn" "scriptsDirectory"`)
		log.Println(`DSN Format: un:pw@tcp(server:port)/db_name?multiStatements=true&collation=utf8mb4_general_ci`)
		os.Exit(1)
	}

	err := Migrate(dsn, scripts)

	if err != nil {
		log.Println("build: ", BUILD)
		log.Println("supplied dsn: ", dsn)
		log.Println("supplied scriptsDirectory: ", scripts)
		log.Fatal(err.Error())
	}
}
