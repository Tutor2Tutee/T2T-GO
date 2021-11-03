package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/Tutor2Tutee/T2T-GO/db"
	"github.com/Tutor2Tutee/T2T-GO/middlewares"
	"github.com/gin-gonic/gin"
)

func GetRouter(database *db.Resource) *gin.Engine {
	// Engine init
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	r := router.Group("/api")
	controllers.Start(database)

	// Available Routes
	userRouterInit(r)
	classesRouterInit(r)
	quizzesRouterInit(r)

	return router
}
