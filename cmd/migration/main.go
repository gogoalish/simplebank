package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/mattes/migrate/source/file"
)

const (
	databaseURL   = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	migrationPath = "file:///home/student/simplebank/db/migration"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main.go [up|down]")
	}
	m, err := migrate.New(
		migrationPath,
		databaseURL)
	if err != nil {
		log.Fatal("failed to migrate: ", err)
	}

	verbose := os.Args[1]
	switch verbose {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "up1":
		err = m.Steps(1)
	case "down1":
		err = m.Steps(-1)
	default:
		log.Fatal("Usage: main.go [up|down]")
	}
	if err != nil {
		log.Fatal("failed migration: ", err)
	}
}
