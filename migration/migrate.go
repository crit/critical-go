package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db                 *sql.DB
	dbNameRegex        = regexp.MustCompile(`.+\/(.+)\?.+`)
)

// Migrate runs the migrations in scriptDire
func Migrate(dsn string, scriptsDirectory string) error {
	// Make sure we support multiStatements
	if !strings.Contains(dsn, "multiStatements=true") {
		return fmt.Errorf("you must include multiStatements=true in your dsn")
	}

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("First ping to database failed: ", err.Error())
		// Database might not exist. Try creating it.
		err = createDatabase(dsn)
		if err != nil {
			return fmt.Errorf("unable to ping database: %s", err.Error())
		}
	}

	err = createMigrationTableIfNotExists()
	if err != nil {
		return fmt.Errorf("failed to create table: %s", err.Error())
	}

	ranScripts, err := getRanScripts()
	if err != nil {
		return fmt.Errorf("unable to get ran scripts: %s", err.Error())
	}

	files, err := getScriptFiles(scriptsDirectory)
	if err != nil {
		return fmt.Errorf("unable to find scripts: %s", err.Error())
	}

	err = runNewMigrationScripts(ranScripts, scriptsDirectory, files)
	if err != nil {
		return fmt.Errorf("unable to run all scripts: %s", err.Error())
	}

	return nil
}

func createDatabase(dsn string) error {
	matches := dbNameRegex.FindStringSubmatch(dsn)
	if len(matches) < 2 {
		return fmt.Errorf("unable to find database name in migrations")
	}

	dbName := matches[1]
	masterDsn := strings.Replace(dsn, dbName, "", -1)

	masterDB, err := sql.Open("mysql", masterDsn)
	if err != nil {
		return err
	}
	defer masterDB.Close()

	_, err = masterDB.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		return fmt.Errorf("unable to run create database statement")
	}

	return nil
}

func createMigrationTableIfNotExists() error {
	rows, err := db.Query("SHOW TABLES LIKE '_migrations'")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Table doesn't exist
	if !rows.Next() {
		_, err := db.Exec(createMigrationTableSQL)
		if err != nil {
			return err
		}
	}

	return nil
}

func getRanScripts() (scripts map[string]bool, err error) {
	scripts = make(map[string]bool)

	rows, err := db.Query(getRanMigrationsSQL)
	if err != nil {
		return scripts, err
	}

	for rows.Next() {
		var scriptName string
		err = rows.Scan(&scriptName)
		if err != nil {
			return scripts, err
		}

		scripts[scriptName] = true
	}

	return scripts, nil
}

func getScriptFiles(directory string) (scripts []string, err error) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		return scripts, err
	}

	for _, file := range files {
		if !file.IsDir() {
			scripts = append(scripts, file.Name())
		}
	}

	return scripts, err
}

func runNewMigrationScripts(ranScripts map[string]bool, directory string, files []string) error {
	for _, script := range files {
		if ranScripts[script] {
			log.Printf("Skipping: %s, already ran", script)
			continue
		}

		log.Printf("Running: %s", script)
		err := runScript(filepath.Join(directory, script))
		if err != nil {
			return err
		}

		err = recordScriptRun(script)
		if err != nil {
			return err
		}
	}

	return nil
}

func recordScriptRun(file string) error {
	_, err := db.Exec("INSERT INTO _migrations (script) VALUES (?)", file)
	return err
}

func runScript(path string) error {
	log.Printf("Reading file: %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read file %s, error: %s", path, err.Error())
	}

	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("unable to execute script: %s, error: %s", path, err.Error())
	}

	return nil
}

var createMigrationTableSQL = `
	CREATE TABLE _migrations (
		script VARCHAR(255) NOT NULL,
		date_ran DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
`

var getRanMigrationsSQL = `
	SELECT
		script
	FROM _migrations
`
