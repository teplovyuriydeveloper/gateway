package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teplovyuriydeveloper/gateway/internal/handlers"
)

func main() {
	r := gin.Default()

	r.Any("/users/*path", handlers.Proxy("http://localhost:8081"))
	r.Any("/orders/*path", handlers.Proxy("http://localhost:8082"))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	r.Run(":8080")
}
