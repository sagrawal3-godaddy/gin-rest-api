package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTeacherController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello wordl teacher controller",
	})
}
