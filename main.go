package main

import (
	"log"

	userApp "github.com/alejandroimen/LongYShortPolling.git/src/Persons/application"
	userController "github.com/alejandroimen/LongYShortPolling.git/src/Persons/infrastructure/controllers"
	userRepo "github.com/alejandroimen/LongYShortPolling.git/src/Persons/infrastructure/repository"
	userRoutes "github.com/alejandroimen/LongYShortPolling.git/src/Persons/infrastructure/routes"
	"github.com/alejandroimen/LongYShortPolling.git/src/core"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conexión a MySQL
	db, err := core.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Repositorios
	personRepository := userRepo.NewCreateUserRepoMySQL(db)

	// Casos de uso para usuarios
	createPerson := userApp.NewCreatePerson(personRepository)
	getUsers := userApp.NewGetPersons(personRepository)
	countGender := userApp.NewCountGender(personRepository)

	// Controladores para usuarios
	CreatePersonController := userController.NewCreateUserController(createPerson)
	getUserController := userController.NewUsersController(getUsers)
	CountGenderController := userController.NewDeleteUserController(countGender)

	// Configuración del enrutador de Gin
	r := gin.Default()

	// Configurar rutas de usuarios
	userRoutes.SetupUserRoutes(r, CreatePersonController, getUserController, CountGenderController)

	// Iniciar servidor
	log.Println("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
