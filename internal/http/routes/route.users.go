package routes

import (
	Repository "gin-framework-boilerplate/internal/adapters/repository/postgresql"
	Usecase "gin-framework-boilerplate/internal/business/usecases"
	Handler "gin-framework-boilerplate/internal/http/handlers"
	ESBPorts "gin-framework-boilerplate/internal/ports/clients/esb"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type usersRoutes struct {
	Handler        Handler.UserHandler
	router         *gin.RouterGroup
	db             *sqlx.DB
	authMiddleware gin.HandlerFunc
	esbC           ESBPorts.ESBClient
}

func NewUsersRoute(router *gin.RouterGroup, db *sqlx.DB, authMiddleware gin.HandlerFunc, esbC ESBPorts.ESBClient) *usersRoutes {
	UserRepository := Repository.NewUserRepository(db)
	UserUsecase := Usecase.NewUserUsecase(UserRepository, esbC)
	UserHandler := Handler.NewUserHandler(UserUsecase)

	return &usersRoutes{Handler: UserHandler, router: router, db: db, authMiddleware: authMiddleware, esbC: esbC}
}

func (r *usersRoutes) Routes() {
	// Routing
	r.router.GET("/user/:email", r.authMiddleware, r.Handler.GetUserByEmail)
	r.router.GET("/users", r.authMiddleware, r.Handler.GetUsers)
}
