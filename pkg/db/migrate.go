package db

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Action int

const (
	ActionUp   Action = iota + 1 // Migrate up
	ActionDown                   // Migrate down
)

func Migrate(url, path string, action Action, fs embed.FS) error {
	d, err := iofs.New(fs, path)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, url)
	if err != nil {
		return err
	}

	defer func() {
		_, _ = m.Close()
	}()

	currentVersion, _, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return err
	}

	if action == ActionUp {
		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	} else if action == ActionDown {
		err = m.Steps(-1)
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}

	} else {
		return errors.New("invalid action")
	}

	printFile(fs, currentVersion, action)

	return nil
}

func MigrateUp(url, path string, fs embed.FS) error {
	return Migrate(url, path, ActionUp, fs)
}

func MigrateDown(url, path string, fs embed.FS) error {
	return Migrate(url, path, ActionDown, fs)
}

func printFile(fs embed.FS, currentVersion uint, action Action) {
	readDir, err := fs.ReadDir("migration")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if action == ActionUp {
		count := 0
		for idx, file := range readDir {
			if strings.Contains(file.Name(), "up") && (idx+1)/2 > int(currentVersion) {
				count++
				fmt.Println(file.Name())
			}
		}

		if count == 0 {
			fmt.Println("No migration file found")
		} else if count == 1 {
			fmt.Printf("1 file found\n")
		} else {
			fmt.Printf("%d files found\n", count)
		}
	}

	if action == ActionDown {
		for idx, file := range readDir {
			if strings.Contains(file.Name(), "down") &&
				idx/2 == int(currentVersion-1) {
				fmt.Println(file.Name())
			}
		}
		fmt.Println("Last migration down successfully, Current Version:", currentVersion-1)
	}
}
