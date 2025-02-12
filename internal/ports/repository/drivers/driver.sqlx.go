package drivers

import (
	"fmt"
	"os"
	"time"

	"database/sql"
	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	zerologadapter "github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	"github.com/sirupsen/logrus"
)

// SQLXConfig holds the configuration for the database instance
type SQLXConfig struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	MaxLifetime    time.Duration
	Debug          bool
}

// InitializeSQLXDatabase returns a new DBInstance
func (config *SQLXConfig) InitializeSQLXDatabase() (*sqlx.DB, error) {
	// Define DB
	sqlDb, err := sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if config.Debug {
		// Define logger adapter
		// Reference : https://github.com/simukti/sqldb-logger
		loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))

		// Configure logger options
		loggerOptions := []sqldblogger.Option{
			sqldblogger.WithSQLQueryFieldname("sql"),
			sqldblogger.WithWrapResult(false),
			sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
			sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
			sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
		}
		sqlDb = sqldblogger.OpenDriver(config.DataSourceName, sqlDb.Driver(), loggerAdapter, loggerOptions...)
	}

	// Pass it to sqlx
	db := sqlx.NewDb(sqlDb, config.DriverName)

	// Set maximum number of open connections to database
	logger.Info(fmt.Sprintf("setting maximum number of open connections to %d", config.MaxOpenConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	db.SetMaxOpenConns(config.MaxOpenConns)

	// Set maximum number of idle connections in the pool
	logger.Info(fmt.Sprintf("setting maximum number of idle connections to %d", config.MaxIdleConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	db.SetMaxIdleConns(config.MaxIdleConns)

	// Set maximum time to wait for new connection
	logger.Info(fmt.Sprintf("setting maximum lifetime for a connection to %s", config.MaxLifetime), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	db.SetConnMaxLifetime(config.MaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}
