package routes

import (
	"github.com/gin-gonic/gin"
)

func RegistrarRotas(server *gin.Engine) {
	server.GET("/eventos", getEventos)
	server.GET("/eventos/:id", showEvento)
	server.POST("/eventos", createEvento)
}
