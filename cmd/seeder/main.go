package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gin-framework-boilerplate/internal/adapters/repository/postgresql"
	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/internal/ports/repository/records"
	"gin-framework-boilerplate/pkg/helpers"
	"gin-framework-boilerplate/pkg/logger"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Define work directory
const (
	dir = "cmd/seeder/seeders"
)

func init() {
	// Initialize app config
	if err := config.InitializeAppConfig(false); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("Configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	// Setup database connection
	db, err := postgresql.SetupPostgresqlConnection()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}
	defer db.Close()

	// Run seeding process
	err = seeding(db)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}
}

// Function to run seeding process
func seeding(db *sqlx.DB) (err error) {
	logger.InfoF("Running seeding process", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Define appsettings object
	appsettingObj := records.Appsettings{}

	// Get current Latest Seeder value
	latestSeeder := "0000000000"
	err = db.QueryRow(`SELECT * FROM appsettings WHERE key = 'Latest Seeder'`).Scan(
		&appsettingObj.Id,
		&appsettingObj.Key,
		&appsettingObj.Value,
	)
	if err != nil {
		// This means that no seeder has been executed
		// So, the first things to do is inserting "Latest Seeder" value to appsettings table

		// Generate UUID
		id, err := helpers.GenerateUUID()
		if err != nil {
			logger.ErrorF("error generating UUID", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
			return err
		}

		// Insert new appsettings record
		execStatement := fmt.Sprintf("INSERT INTO appsettings (id, key, value) VALUES ('%s', 'Latest Seeder', '0000000000')", id)
		_, err = db.Exec(execStatement)
		if err != nil {
			logger.ErrorF("error insertng initial data to appsettings table", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
			return err
		}
	} else {
		// Assign latest seeder value
		latestSeeder = appsettingObj.Value
	}

	// Retrieves all seeder files
	files, err := filepath.Glob(filepath.Join(cwd, dir, "*.seeder.sql"))
	if err != nil {
		return errors.New("error when get files name")
	}

	// Loop through each file
	for _, file := range files {
		// We will only execute following command only if the version is greater than latest seeder
		pathSlice := strings.Split(file, "\\")
		seederVersion := strings.Split(pathSlice[len(pathSlice)-1], "_")[0]
		if strings.Compare(seederVersion, latestSeeder) > 0 {
			logger.Info("Generating seeder", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder, constants.LoggerFile: file})
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return errors.New("error when read file")
			}

			// Executing the command in SQL file
			_, err = db.Exec(string(data))
			if err != nil {
				fmt.Println(err)
				return fmt.Errorf("error when exec query in file:%v", file)
			}

			// Update latest seeder
			latestSeeder = seederVersion
		}
	}

	// Update latest seeder in the database
	newLatestSeederObj := records.Appsettings{
		Key:   "Latest Seeder",
		Value: latestSeeder,
	}
	_, err = db.NamedExec("UPDATE appsettings SET value = :value WHERE key = :key", newLatestSeederObj)
	if err != nil {
		return errors.New("error when updating Last Seeder value")
	}

	logger.InfoF("Seeding process success", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})

	return
}
