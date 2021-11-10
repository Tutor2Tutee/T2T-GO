package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/Tutor2Tutee/T2T-GO/middlewares"
	"github.com/gin-gonic/gin"
)

func quizzesRouterInit(r *gin.RouterGroup) {
	quizzes := r.Group("/quizzes")
	{
		quizzes.POST("", middlewares.JWTAuthenticationMiddleware, controllers.CreateQuiz)
		quizzes.GET("", middlewares.JWTAuthenticationMiddleware, controllers.GetAllQuiz)
		quizzes.GET("/:quizID", middlewares.JWTAuthenticationMiddleware, controllers.GetQuizByID)
		quizzes.GET("/creator/:creatorID", middlewares.JWTAuthenticationMiddleware, controllers.GetQuizByCreatorID)
		quizzes.DELETE("/:quizID", middlewares.JWTAuthenticationMiddleware, controllers.DeleteQuizByID)
	}
}
