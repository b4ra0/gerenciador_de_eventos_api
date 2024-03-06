package main

import (
	"fmt"
	"gerenciador_de_eventos/db"
	"gerenciador_de_eventos/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	db.Connectar()

	server.GET("/eventos", getEventos)
	server.POST("/eventos", createEvento)

	server.Run(":8080") // localhost:8080
}

func createEvento(context *gin.Context) {
	var evento models.Evento
	err := context.ShouldBindJSON(&evento)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "dados incorretos"})
		return
	}

	evento.Id = 1
	evento.UserID = 1

	evento.Save()

	context.JSON(http.StatusCreated, gin.H{"mensagem": "Evento criado!", "evento": evento})
}

func getEventos(context *gin.Context) {
	eventos := models.GetAllEventos()
	context.JSON(http.StatusOK, eventos)
}
