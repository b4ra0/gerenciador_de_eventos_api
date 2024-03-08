package routes

import (
	"fmt"
	"gerenciador_de_eventos/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func updateEvento(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	idint := int(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Você deve informar um número inteiro"})
		return
	}
	_, err = models.GetEventoById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Este evento não existe"})
		return
	}

	var eventoAtualizado models.Evento
	err = context.ShouldBindJSON(&eventoAtualizado)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "dados incorretos"})
		return
	}
	eventoAtualizado.Id = idint
	err = eventoAtualizado.Update()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Ocorreu um erro ao atualizar no banco de dados"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"mensagem": "Evento atualizado com sucesso!"})

}

func deleteEvento(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Você deve informar um número inteiro"})
		return
	}
	evento, err := models.GetEventoById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Este evento não existe"})
		return
	}

	err = evento.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mensagem": "Não foi possível excluir o evento"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Evento deletado com sucesso!"})

}
