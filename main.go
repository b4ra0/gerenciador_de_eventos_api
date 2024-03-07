package main

import (
	"fmt"
	"gerenciador_de_eventos/db"
	"gerenciador_de_eventos/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	server := gin.Default()

	db.Connectar()

	server.GET("/eventos", getEventos)
	server.GET("/eventos/:id", showEvento)
	server.POST("/eventos", createEvento)

	server.Run(":8080") // localhost:8080
}

func showEvento(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Você deve informar um número inteiro"})
		return
	}
	evento, err := models.GetEventoById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Não foi possível mostrar o evento"})
		return
	}
	context.JSON(http.StatusOK, evento)
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

	err = evento.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Ocorreu um erro ao salvar no banco de dados"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"mensagem": "Evento criado!", "evento": evento})
}

func getEventos(context *gin.Context) {
	eventos, err := models.GetAllEventos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mensagem": "Não foi possível mostrar os eventos"})
		return
	}
	context.JSON(http.StatusOK, eventos)
}
