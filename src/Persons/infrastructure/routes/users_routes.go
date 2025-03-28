package routes

import (
	"github.com/alejandroimen/API_HEXAGONAL/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, createUserController *controllers.CreateUserController, getUserController *controllers.GetUsersController, deleteUserController *controllers.DeleteUserController, updateUserController *controllers.UpdateUserController) {

	r.POST("/users", createUserController.Handle)

}
