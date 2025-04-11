package routes

import (
	"github.com/alejandroimen/LongYShortPolling.git/src/Persons/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, CreatePersonController *controllers.CreatePersonController, GetPersonController *controllers.GetPersonController, CountGender *controllers.CountGenderController) {

	r.POST("/addPerson", CreatePersonController.Handle)
	r.GET("/newPersonIsAdded", GetPersonController.ShortPoll)
	r.GET("/countGender", CountGender.LongPoll)
}
