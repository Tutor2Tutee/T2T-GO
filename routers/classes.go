package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func classesRouterInit(r *gin.RouterGroup) {
	classes := r.Group("/classes")
	{
		classes.GET("", controllers.GetAll)
		classes.POST("", controllers.Create)
	}
}
