package controllers

import (
	"log"

	"github.com/alejandroimen/LongYShortPolling/src/Persons/application"
	"github.com/gin-gonic/gin"
)

type CreatePersonController struct {
	CreatePerson *application.CreatePerson
}

func NewCreatePersonController(CreatePerson *application.CreatePerson) *CreatePersonController {
	return &CreatePersonController{CreatePerson: CreatePerson}
}

func (c *CreatePersonController) Handle(ctx *gin.Context) {
	log.Println("Petición de crear una persona, recibido")

	var request struct {
		Name   string `json:"name"`
		Age    string `json:"age"`
		Gender string `json:"gender"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando la petición del body: %v", err)
		ctx.JSON(400, gin.H{"error": "petición del body inválida"})
		return
	}
	log.Printf("Creando usuario: Name=%s, Age=%s, Genero=%s", request.Name, request.Age, request.Gender)

	if err := c.CreatePerson.Run(request.Age, request.Name, request.Gender); err != nil {
		log.Printf("Error creando el usuario: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Usuario creado exitosamente")
	ctx.JSON(201, gin.H{"message": "persona creado exitosamente"})
}
