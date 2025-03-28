package routes

import (
	"github.com/alejandroimen/LongYShortPolling.git.git/src/Persons/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, CreatePersonController *controllers.CreatePersonController, GetPersonController *controllers.getPersonController) {

	r.POST("/addPerson", CreatePersonController.Handle)
	r.POST("/newPersonIsAdded", GetPersonController.ShortPoll)

}
