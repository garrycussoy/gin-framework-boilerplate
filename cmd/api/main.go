package main

import (
	"runtime"

	"gin-framework-boilerplate/cmd/api/server"
	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/pkg/logger"

	"github.com/sirupsen/logrus"
)

// Initialize config
func init() {
	if err := config.InitializeAppConfig(false); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	// Setting up CPU-related things
	numCPU := runtime.NumCPU()
	logger.InfoF("the project is running on %d CPU(s)", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig}, numCPU)

	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(numCPU / 2)
	}

	// Initialize the server
	app, err := server.NewApp()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
	if err := app.Run(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
}
