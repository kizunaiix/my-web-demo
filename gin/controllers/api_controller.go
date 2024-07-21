package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.String(http.StatusOK, "welcome to api")
}

func Login(ctx *gin.Context) {
	var request struct {
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Content == "hello api" {
		ctx.JSON(http.StatusOK, gin.H{"content": "hello my client!"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"content": "who r u?"})
	}
}
