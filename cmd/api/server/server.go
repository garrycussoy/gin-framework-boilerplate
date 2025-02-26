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

	ESBAdapters "gin-framework-boilerplate/internal/adapters/clients/esb"
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

	docs "gin-framework-boilerplate/cmd/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	httpClient := setupHttpClient()

	// Setup middleware
	router = setupMiddleware(router, httpClient)

	// Define clients service
	esbClient := ESBAdapters.NewESBClient(httpClient)

	// JWT service
	jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)

	// Notification service
	// emailNotificationService := notifications.NewSendEmailNotificationService(config.AppConfig.EmailSender, config.AppConfig.EmailPassword)

	// Initialize auth middleware
	authMiddleware := middlewares.NewAuthMiddleware(jwtService)

	// API routes
	api := router.Group("bo-api")
	routes.NewGeneralsRoute(api).Routes()
	routes.NewAuthRoute(api, conn, jwtService).Routes()
	routes.NewUsersRoute(api, conn, authMiddleware, esbClient).Routes()

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

	// Setup swagger
	docs.SwaggerInfo.BasePath = "/bo-api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// A function to setup middlewares
func setupMiddleware(r *gin.Engine, httpClient *resty.Client) *gin.Engine {
	// Set up middlewares
	r.Use(middlewares.CORSMiddleware())                      // Setup CORS
	r.Use(middlewares.CorrelationIdMiddleware())             // Setup Correlation-ID
	r.Use(gin.LoggerWithFormatter(logger.HTTPLogger))        // Log some basic data of incoming HTTP request
	r.Use(logger.RequestPayloadLogger())                     // Log incoming HTTP request payload
	r.Use(logger.ResponsePayloadLogger())                    // Log response body of incoming HTTP request
	r.Use(logger.ExternalHTTPRequestMiddleware(httpClient))  // Log external HTTP request payload
	r.Use(logger.ExternalHTTPResponseMiddleware(httpClient)) // Log response body of external HTTP request
	r.Use(middlewares.TimeoutMiddleware())                   // Setup general endpoint timeout
	r.Use(gin.Recovery())                                    // Recover any panic

	return r
}

// A function to setup HTTP call system
func setupHttpClient() *resty.Client {
	// Create a Resty client
	client := resty.New()

	// Setup retry policy
	client.SetRetryCount(5).
		SetRetryWaitTime(3 * time.Second).
		SetRetryMaxWaitTime(30 * time.Second)

	return client
}
