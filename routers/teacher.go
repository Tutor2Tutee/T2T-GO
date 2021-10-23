package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func teacherRouterInit(r *gin.Engine) {
	//specific Route groups
	teacher := r.Group("/teacher")

	{
		teacher.GET("/", controllers.GetTeacher)
	}
}
