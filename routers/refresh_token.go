package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func refreshTokenRouterInit(r *gin.RouterGroup) {
	refresh := r.Group("/refresh_token")
	{
		refresh.GET("", controllers.RefreshToken)
	}
}
