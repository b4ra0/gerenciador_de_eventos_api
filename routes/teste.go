package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func teste(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"mensagem": "Hello World"})
}
