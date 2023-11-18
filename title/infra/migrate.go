package infra

import (
	"embed"

	"github.com/amjadjibon/microservices/pkg/db"
)

//go:embed migration/*.sql
var fs embed.FS

func MigrationUp(url string) error {
	return db.MigrateUp(url, "migration", fs)
}

func MigrationDown(url string) error {
	return db.MigrateDown(url, "migration", fs)
}
