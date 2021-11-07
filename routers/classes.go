package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func classesRouterInit(r *gin.RouterGroup) {
	//classController := controllers.Class()
	classes := r.Group("/classes")
	{
		classes.GET("", controllers.ClassController{}.GetAll)
		classes.POST("", controllers.ClassController{}.Create)
		classes.GET("/:cid", controllers.ClassController{}.GetOne)
	}
}
