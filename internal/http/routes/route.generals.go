package routes

import (
	Handler "gin-framework-boilerplate/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

type generalsRoutes struct {
	Handler Handler.GeneralHandler
	router  *gin.RouterGroup
}

func NewGeneralsRoute(router *gin.RouterGroup) *generalsRoutes {
	GeneralHandler := Handler.NewGeneralHandler()

	return &generalsRoutes{Handler: GeneralHandler, router: router}
}

func (r *generalsRoutes) Routes() {
	// List of routes of "General" category
	r.router.GET("/health-check", r.Handler.HealthCheck)
}
