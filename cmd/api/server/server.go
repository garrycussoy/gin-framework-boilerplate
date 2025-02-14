package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-framework-boilerplate/internal/adapters/repository/postgresql"
	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/constants"

	"gin-framework-boilerplate/internal/http/middlewares"
	"gin-framework-boilerplate/internal/http/routes"

	"gin-framework-boilerplate/pkg/jwt"
	"gin-framework-boilerplate/pkg/logger"

	// "gin-framework-boilerplate/pkg/notifications"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	// Setup database connection
	conn, err := postgresql.SetupPostgresqlConnection()
	if err != nil {
		return nil, err
	}

	// Setup router
	router := setupRouter()

	// Setup HTTP client
	// httpClient := setupHttpClient()

	// JWT service
	jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)

	// Notification service
	// emailNotificationService := notifications.NewSendEmailNotificationService(config.AppConfig.EmailSender, config.AppConfig.EmailPassword)

	// User middleware
	// User with valid basic token can access endpoint
	// authMiddleware := middlewares.NewAuthMiddleware(jwtService, false)

	// Timeout middleware
	// timeoutMiddleware := middlewares.NewTimeoutMiddleware()

	// API routes
	api := router.Group("bo-api")
	routes.NewGeneralsRoute(api).Routes()
	routes.NewAuthRoute(api, conn, jwtService).Routes()
	// routes.NewPatientsRoute(api, conn, authMiddleware, timeoutMiddleware, wiproClient, omnicareClient).Routes()
	// routes.NewMasterDataRoute(api, conn, authMiddleware, timeoutMiddleware, omnicareClient, wiproClient).Routes()
	// routes.NewBillingsRoute(api, conn, authMiddleware, timeoutMiddleware, omnicareClient, wiproClient).Routes()
	// routes.NewOrderRoute(api, conn, authMiddleware, timeoutMiddleware, omnicareClient, wiproClient).Routes()

	// Setup HTTP server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(config.AppConfig.ServerReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.AppConfig.ServerWriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() (err error) {
	// Graceful shutdown
	go func() {
		logger.InfoF("success to listen and serve on :%d", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, config.AppConfig.Port)
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Make blocking channel and waiting for a signal
	<-quit
	logger.Info("shutdown server ...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// Catching ctx.Done()
	// Timeout of 5 seconds
	<-ctx.Done()
	logger.Info("timeout of 5 seconds.", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	logger.Info("server exiting", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	return
}

// A function to adjust routing system
func setupRouter() *gin.Engine {
	// Set the runtime mode
	var mode = gin.ReleaseMode
	if config.AppConfig.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// Create a new router instance
	router := gin.New()

	// Set up middlewares
	router.Use(middlewares.CORSMiddleware())               // Setup CORS
	router.Use(middlewares.CorrelationIdMiddleware())      // Setup Correlation-ID
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger)) // Log some basic data of incoming HTTP request
	router.Use(logger.PayloadRequestLogger())              // Log incoming HTTP request payload
	router.Use(gin.Recovery())                             // Recover any panic

	return router
}

// A function to setup HTTP call system
func setupHttpClient() *resty.Client {
	// Create a Resty client
	client := resty.New()

	// Setup logging system
	// client.OnBeforeRequest(logger.ExternalHTTPRequestLogger)
	// client.OnAfterResponse(logger.ExternalHTTPResponseLogger)

	// Setup retry policy
	client.SetRetryCount(5).
		SetRetryWaitTime(3 * time.Second).
		SetRetryMaxWaitTime(30 * time.Second)

	return client
}
