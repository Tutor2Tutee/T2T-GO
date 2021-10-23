package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func studentRouterInit(r *gin.Engine) {
	//specific Route groups
	student := r.Group("/student")

	{
		student.GET("/", controllers.GetStudent)
	}
}
