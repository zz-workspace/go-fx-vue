package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"fast-api.io/internal/repository"
	"fast-api.io/models"
	"fast-api.io/modules/http/response"
	"github.com/gin-gonic/gin"
)

type ApiRequest struct {
	Name        string `json:"name" validate:"required"`
	WorkspaceID uint64 `json:"workspace_id" validate:"required"`
}

type ApiController struct {
	apiRepository *repository.ApiRepository
}

func InitApiController(r *gin.RouterGroup, apiRepository *repository.ApiRepository) {
	controller := &ApiController{
		apiRepository: apiRepository,
	}

	apiRoutes := r.Group("api")
	apiRoutes.GET("", controller.GetListAPI)
	apiRoutes.POST("", controller.CreateApi)
}

func (e ApiController) CreateApi(ctx *gin.Context) {
	var body ApiRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err)
		return
	}
	api := e.apiRepository.CreateApi(&models.API{
		Name:        body.Name,
		WorkspaceID: body.WorkspaceID,
	})

	response.JSON(ctx, http.StatusOK, api)
}

func (e ApiController) GetListAPI(ctx *gin.Context) {
	workspaceID, err := strconv.ParseUint(ctx.Query("workspace_id"), 10, 64)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, fmt.Errorf("Missing workspace_id"))
		return
	}
	apis := e.apiRepository.ListAPI(workspaceID)
	response.JSON(ctx, http.StatusOK, apis)
}
