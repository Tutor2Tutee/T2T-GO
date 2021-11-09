package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/Tutor2Tutee/T2T-GO/middlewares"
	"github.com/gin-gonic/gin"
)

func classesRouterInit(r *gin.RouterGroup) {
	//classController := controllers.Class()
	classes := r.Group("/classes")
	{
		classes.GET("", middlewares.JWTAuthenticationMiddleware, controllers.ClassController{}.GetAll)
		classes.POST("", middlewares.JWTAuthenticationMiddleware, controllers.ClassController{}.Create)
		classes.GET("/:cid", middlewares.JWTAuthenticationMiddleware, controllers.ClassController{}.GetOne)
	}
}
