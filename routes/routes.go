package routes

import (
	"github.com/gin-gonic/gin"
)

func RegistrarRotas(server *gin.Engine) {
	server.GET("/hello-world", teste)
	server.GET("/eventos", getEventos)
	server.GET("/eventos/:id", showEvento)
	server.POST("/eventos", createEvento)
	server.PUT("/eventos/:id", updateEvento)
	server.DELETE("/eventos/:id", deleteEvento)
	//server.POST("/signin")
}
