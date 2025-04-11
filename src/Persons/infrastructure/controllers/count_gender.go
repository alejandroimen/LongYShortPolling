package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/LongYShortPolling.git/src/Persons/application"
	"github.com/gin-gonic/gin"
)

type CountGenderController struct {
	countGender *application.CountGender
}

func NewCountGenderController(countGender *application.CountGender) *CountGenderController {
	return &CountGenderController{countGender: countGender}
}

func (c *CountGenderController) LongPoll(ctx *gin.Context) {
	log.Println("Iniciando long polling para conteo de género")

	// Tiempo máximo de espera
	timeout := time.Now().Add(30 * time.Second)

	// Últimos valores conocidos
	var lastCount []int

	for {
		// Llamar al caso de uso para obtener los conteos actuales
		currentCount, err := c.countGender.Run()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Si es la primera iteración o los valores han cambiado, respondemos
		if lastCount == nil || lastCount[0] != currentCount[0] || lastCount[1] != currentCount[1] {
			lastCount = currentCount
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Datos actualizados",
				"countMan": currentCount[0],
				"countWomen": currentCount[1],
			})
			return
		}

		// Si el tiempo de espera ha terminado, devolvemos la respuesta sin cambios
		if time.Now().After(timeout) {
			ctx.JSON(http.StatusOK, gin.H{"message": "No hay cambios en los datos"})
			return
		}

		// Dormimos un poco antes de volver a consultar
		time.Sleep(2 * time.Second)
	}
}
