package main

import (
	"gerenciador_de_eventos/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//db.Conectar()

	routes.RegistrarRotas(server)

	server.Run(":8080") // localhost:8080
}
