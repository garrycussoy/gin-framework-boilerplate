package routes

import (
	Repository "gin-framework-boilerplate/internal/adapters/repository/postgresql"
	Usecase "gin-framework-boilerplate/internal/business/usecases"
	Handler "gin-framework-boilerplate/internal/http/handlers"
	"gin-framework-boilerplate/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type authRoutes struct {
	Handler Handler.AuthHandler
	router  *gin.RouterGroup
	db      *sqlx.DB
}

func NewAuthRoute(router *gin.RouterGroup, db *sqlx.DB, jwtService jwt.JWTService) *authRoutes {
	UserRepository := Repository.NewUserRepository(db)
	AuthUsecase := Usecase.NewAuthUsecase(jwtService, UserRepository)
	AuthHandler := Handler.NewAuthHandler(AuthUsecase)

	return &authRoutes{Handler: AuthHandler, router: router, db: db}
}

func (r *authRoutes) Routes() {
	// Routing
	r.router.POST("/login", r.Handler.UserLogin)
}
