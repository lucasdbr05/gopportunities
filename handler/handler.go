package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}

func ShowOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}

func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}

func DeleteOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}

func ListOpeningsHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}