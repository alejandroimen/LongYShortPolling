package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPersonController struct {
	getPerson *application.getPerson
}

func NewpersonsController(getPerson *application.getPerson) *GetPersonController {
	return &GetPersonController{getPerson: getPerson}
}

func (gu *GetPersonController) Handle(ctx *gin.Context) {
	log.Println("Petición de listar todos los usuarios, recibido")

	person, err := gu.getPerson.Run()
	if err != nil {
		log.Printf("Error buscando usuarios")
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retornando %d usuarios", len(person))
	ctx.JSON(200, person)
}
func (c *GetPersonController) ShortPoll(ctx *gin.Context) {
	// Obtener los productos (esto simula si hay cambios o no)
	products, err := c.getPerson.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		// No hay usuarios (o cambios)
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
		return
	}

	// Devolver productos (o cambios detectados)
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Datos actualizados",
		"products": products,
	})
}
