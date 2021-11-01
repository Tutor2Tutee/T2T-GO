package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/middlewares"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	// Engine init
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	r := router.Group("/api")

	// Available Routes
	userRouterInit(r)
	classesRouterInit(r)
	quizzesRouterInit(r)

	return router
}
