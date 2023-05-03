package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"fast-api.io/helpers/converter"
	"fast-api.io/internal/repository"
	"fast-api.io/models"
	"fast-api.io/modules/http/response"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type FunctionRequest struct {
	Name   models.FunctionName `json:"name" validate:"required"`
	Input  datatypes.JSON      `json:"input"`
	Run    datatypes.JSON      `json:"run"`
	Output datatypes.JSON      `json:"output"`
}

type FunctionController struct {
	functionRepository *repository.FunctionRepository
}

func InitFunctionController(r *gin.RouterGroup, functionRepository *repository.FunctionRepository) {
	controller := &FunctionController{
		functionRepository: functionRepository,
	}

	functionRoutes := r.Group("functions")
	functionRoutes.GET("", controller.ListFunctions)
	functionRoutes.GET(":id", controller.FindById)
	functionRoutes.POST("", controller.CreateFunction)
	functionRoutes.PATCH(":id", controller.UpdateFunction)
}

func (e FunctionController) CreateFunction(ctx *gin.Context) {
	var body FunctionRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err)
		return
	}

	e.functionRepository.Create(&models.Function{
		Name:   body.Name,
		Run:    body.Run,
		Output: body.Output,
	})

	response.JSON(ctx, http.StatusCreated, body)
}

func (f FunctionController) FindById(ctx *gin.Context) {
	functionId, err := converter.StringToInt(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}
	function := f.functionRepository.FindById(functionId)
	response.JSON(ctx, http.StatusOK, function)
}

func (f FunctionController) ListFunctions(ctx *gin.Context) {
	functions, _ := f.functionRepository.Find()
	response.JSON(ctx, http.StatusOK, functions)
}

func (e FunctionController) UpdateFunction(ctx *gin.Context) {
	functionId, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(functionId)

	function := e.functionRepository.FindById(functionId)
	var body FunctionRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err)
		return
	}

	function.Name = body.Name
	function.Input = body.Input
	function.Output = body.Output
	function.Run = body.Run
	e.functionRepository.Update(&function)
	response.JSON(ctx, http.StatusCreated, function)
}
