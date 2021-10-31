package routers

import (
	"github.com/Tutor2Tutee/T2T-GO/controllers"
	"github.com/gin-gonic/gin"
)

func quizzesRouterInit(r *gin.RouterGroup) {
	quizzes := r.Group("/quizzes")
	{
		quizzes.POST("/", controllers.CreateQuiz)
		quizzes.GET("/", controllers.GetAllQuiz)
		quizzes.GET("/:quizID", controllers.GetQuizByID)
	}
}
