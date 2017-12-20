package main

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"testing"
)

func createTempTestingDir(token string) string {
	dir := fmt.Sprintf("%s/migration-test-%s", os.TempDir(), token)

	err := os.Mkdir(dir, 0777)

	if err != nil {
		log.Fatal("Fatal os.Mkdir ", err.Error())
	}

	return dir
}

func createSQLFile(dir, token string, num int) {
	file, err := os.Create(fmt.Sprintf(`%s/00%d-%s.sql`, dir, num, token))

	if err != nil {
		log.Fatal("Fatal os.Create ", err.Error())
		return
	}

	_, err = file.WriteString(fmt.Sprintf(`CREATE TABLE test_%s_%d (id INT);`, token, num))

	if err != nil {
		log.Fatal("Fatal file.WriteString ", err.Error())
		return
	}
}

func TestMigrate(t *testing.T) {
	gotenv.Load(".env")

	token := xid.New().String()

	dsn := os.Getenv("DB_DSN")

	dir := createTempTestingDir(token)

	createSQLFile(dir, token, 1)
	createSQLFile(dir, token, 2)

	err := Migrate(dsn, dir)

	assert.Nil(t, err)
}
