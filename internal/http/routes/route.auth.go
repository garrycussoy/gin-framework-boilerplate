package routes

import (
	Usecase "gin-framework-boilerplate/internal/business/usecases"
	// PostgresRepository "gin-framework-boilerplate/internal/datasources/repositories/postgres"
	Handler "gin-framework-boilerplate/internal/http/handlers"
	// "gin-framework-boilerplate/internal/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type authRoutes struct {
	Handler Handler.AuthHandler
	router  *gin.RouterGroup
	db      *sqlx.DB
	// authMiddleware    gin.HandlerFunc
	// timeoutMiddleware middlewares.TimeoutMiddleware
}

func NewAuthRoute(router *gin.RouterGroup, db *sqlx.DB) *authRoutes {
	// PatientRepository := PostgresRepository.NewPatientRepository(db)
	// MasterDataRepository := PostgresRepository.NewMasterDataRepository(db)
	AuthUsecase := Usecase.NewAuthUsecase()
	AuthHandler := Handler.NewAuthHandler(AuthUsecase)

	return &authRoutes{Handler: AuthHandler, router: router, db: db}
}

func (r *authRoutes) Routes() {
	// Setup auth middleware
	// r.router.Use(r.authMiddleware)

	// Routing
	// r.router.GET("/appointment", r.timeoutMiddleware.ShortTimeoutMiddleware(), r.Handler.GetOmnicareAppointments)
	r.router.POST("/login", r.Handler.UserLogin)
}
