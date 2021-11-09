package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/Tutor2Tutee/T2T-GO/middlewares"
	"github.com/gin-gonic/gin"
)

func userRouterInit(r *gin.RouterGroup) {
	//specific Route groups
	user := r.Group("/users")

	{
		user.POST("/register", controllers.RegisterUser)
		user.POST("/login", controllers.LoginUser)
		user.GET("/:userId", middlewares.JWTAuthenticationMiddleware, controllers.GetUserByID)
		user.PATCH("/:userId", middlewares.JWTAuthenticationMiddleware, controllers.UpdateUserByID)
		user.DELETE("/:userId", middlewares.JWTAuthenticationMiddleware, controllers.DeleteUserByID)
	}
}
