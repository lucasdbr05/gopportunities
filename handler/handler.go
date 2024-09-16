package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdbr05/gopportunities/config"
	"net/http"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetDatabase()
}

func CreateOpeningHandler(ctx *gin.Context){
	var request CreateOpeningRequest

	ctx.BindJSON(request)

	logger.Infof("%v", request)
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