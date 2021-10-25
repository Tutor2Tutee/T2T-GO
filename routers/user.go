package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func userRouterInit(r *gin.RouterGroup) {
	//specific Route groups
	user := r.Group("/user")

	{
		user.POST("/register", controllers.RegisterUser)
		// user.POST("/login", controllers.LoginUser)
		// user.POST("/:userId", controllers.GetUserByID)
		// user.PATCH("/:userId", controllers.UpdateUserByID)
		// user.DELETE("/:userId", controllers.DeleteUserByID)
	}
}
