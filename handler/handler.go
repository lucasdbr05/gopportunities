package handler

import (
	"fmt"
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
		sendError(ctx, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query").Error())
		return
	}

	opening := schemas.Opening{} 

	q := db.First(&opening, id)
	if(q.Error != nil){
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("opening with id %s is not found", id))
		return
	}

	sendSuccess(ctx, "find-opening", opening)
}

func UpdateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	logger.Infof("%+v", request)

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query").Error())
		return
	}
	
	opening := schemas.Opening{}
	query:= db.First(&opening, id)

	if(query.Error != nil){
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("opening with id %s is not found", id))
		return
	}
	
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	update := db.Save(&opening)
	if update.Error != nil {
		logger.Errorf("error updating opening: %v", update.Error.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "update-opening", opening)

}

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query").Error())
		return
	}

	opening := schemas.Opening{} 

	q := db.First(&opening, id)

	if(q.Error != nil){
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("opening with id %s is not found", id))
		return
	}
	
	del := db.Delete(&opening)
	
	if(del.Error != nil) {
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("error deleting opening with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	query := db.Find(&openings)
	if(query.Error != nil){
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
	}

	sendSuccess(ctx, "list-openings", openings)

}
