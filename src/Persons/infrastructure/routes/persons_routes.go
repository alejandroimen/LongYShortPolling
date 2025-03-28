package routes

import (
	"github.com/alejandroimen/LongYShortPolling.git/src/Persons/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, CreatePersonController *controllers.CreatePersonController) {

	r.POST("/addPerson", CreatePersonController.Handle)

}
