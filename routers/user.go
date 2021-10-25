package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func userRouterInit(r *gin.RouterGroup) {
	//specific Route groups
	user := r.Group("/user")

	{
		user.GET("/", controllers.GetUser)
	}
}
