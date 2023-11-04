package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtonynet/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreatOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error)
		SendError(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	SendSucces(ctx, "create-opening", opening)
}
