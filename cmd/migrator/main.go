package main

import (
	"flag"

	"github.com/golang-migrate/migrate"
)

func main() {
	var storagePath, migrationsPath string

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations folder")
	flag.Parse()

	if storagePath == "" {
		panic("storage path is required")
	}
	if migrationsPath == "" {
		panic("migrations path is requiered")
	}

	m, err := migrate.New("file://" + mirgramigrationsPath)

	if err != nil {
		panic(err)
	}

}
