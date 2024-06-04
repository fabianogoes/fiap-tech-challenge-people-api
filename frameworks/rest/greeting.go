package rest

import (
	"github.com/fabianogoes/fiap-people/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the People API",
	})
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func Environment(c *gin.Context) {
	config, err := entities.NewConfig()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, config)
}
