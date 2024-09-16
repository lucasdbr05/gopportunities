package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasdbr05/gopportunities/config"
	"github.com/lucasdbr05/gopportunities/schemas"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetDatabase()
}

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	logger.Infof("%+v", request)

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Salary:   request.Salary,
		Remote:   *request.Remote,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Database creation error: %v", err.Error())
		sendError(ctx,http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "HANDLER",
	})
}
