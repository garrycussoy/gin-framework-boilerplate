package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gin-framework-boilerplate/internal/adapters/repository/postgresql"
	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/pkg/logger"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	dir = "cmd/migration/migrations"
)

var (
	up   bool
	down bool
)

func init() {
	// Initialize app config
	if err := config.InitializeAppConfig(false); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("Configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	// Creates two flag: up and down
	flag.BoolVar(&up, "up", false, "involves creating new tables, columns, or other database structures")
	flag.BoolVar(&down, "down", false, "involves dropping tables, columns, or other structures")
	flag.Parse()

	// Setup database connection
	db, err := postgresql.SetupPostgresqlConnection()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration})
	}
	defer db.Close()

	// Process the migration for "up" command
	if up {
		err = migrate(db, "up")
		if err != nil {
			logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration})
		}
	}

	// Process the migration for "down" command
	if down {
		err = migrate(db, "down")
		if err != nil {
			logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration})
		}
	}
}

// Function to do the migration
func migrate(db *sqlx.DB, action string) (err error) {
	logger.InfoF("Running migration [%s]", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration}, action)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Retrieves migration file
	files, err := filepath.Glob(filepath.Join(cwd, dir, fmt.Sprintf("*.%s.sql", action)))
	if err != nil {
		return errors.New("error when get files name")
	}

	// Loop through each file, then execute them in order
	for _, file := range files {
		logger.Info("Executing migration", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration, constants.LoggerFile: file})
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return errors.New("error when read file")
		}

		_, err = db.Exec(string(data))
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("error when exec query in file:%v", file)
		}
	}

	logger.InfoF("Migration [%s] success", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration}, action)

	return
}
