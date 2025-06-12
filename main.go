package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador do Gin
	router := gin.Default()

	// Define o endpoint protegido
	// Este endpoint só deve ser alcançado após o forward-auth do Traefik ser bem-sucedido.
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Acesso concedido ao recurso protegido!",
			"status":  "success",
		})
	})

	// Inicia o servidor na porta 8080
	// É uma boa prática usar uma porta diferente do seu serviço de autenticação.
	router.Run(":8080")
}
