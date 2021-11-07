package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/db"
	"github.com/Tutor2Tutee/T2T-GO/middlewares"
	"github.com/Tutor2Tutee/T2T-GO/repository"
	"github.com/gin-gonic/gin"
)

func GetRouter(resource *db.Resource) *gin.Engine {
	// Engine init
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	r := router.Group("/api")
	repository.SetDatabase(resource)

	// Available Routes
	userRouterInit(r)
	classesRouterInit(r)
	quizzesRouterInit(r)
	refreshTokenRouterInit(r)

	return router
}
