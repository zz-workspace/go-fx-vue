package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fast-api.io/internal/repository"
	"fast-api.io/models"
	"fast-api.io/modules/database"
	"fast-api.io/modules/http/response"
	querybuilder "fast-api.io/modules/query-builder"
	mvpgetby "fast-api.io/modules/query-builder/mvp-get-by"
	mvprawsql "fast-api.io/modules/query-builder/mvp-raw-sql"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gorm.io/datatypes"
)

type EndpointRequest struct {
	Name        string `json:"name" validate:"required"`
	Method      string `json:"method" validate:"required"`
	WorkspaceID uint64 `json:"workspace_id" validate:"required"`
	ApiID       uint64 `json:"api_id" validate:"required"`
}

type EndpointInputRequest struct {
	Input     []models.EndpointInput    `json:"input"`
	Functions []models.EndpointFunction `json:"functions"`
	Response  []models.EndpointResponse `json:"response"`
}

type TestEndpointRequest struct {
	InputData datatypes.JSON `json:"input_data"`
}

type EndpointController struct {
	endpointRepository  *repository.EndpointRepository
	workspaceRepository *repository.WorkspaceRepository
	tableRepository     *repository.TableRepository
}

func InitEndpointController(
	r *gin.RouterGroup,
	endpointRepository *repository.EndpointRepository,
	workspaceRepository *repository.WorkspaceRepository,
	tableRepository *repository.TableRepository,
) {
	controller := &EndpointController{
		endpointRepository:  endpointRepository,
		workspaceRepository: workspaceRepository,
		tableRepository:     tableRepository,
	}
	endpointRoutes := r.Group("endpoints")
	endpointRoutes.GET("", controller.GetListEndpoint)
	endpointRoutes.POST("", controller.CreateEndpoint)
	endpointRoutes.PUT(":id", controller.UpdateEndpoint)
	endpointRoutes.POST(":id/test", controller.TestEndpoint)
}

func (e EndpointController) CreateEndpoint(ctx *gin.Context) {
	var body EndpointRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err)
		return
	}
	endpoint := e.endpointRepository.CreateEndpoint(&models.Endpoint{
		Name:        body.Name,
		Method:      body.Method,
		WorkspaceID: body.WorkspaceID,
		ApiID:       body.ApiID,
	})
	response.JSON(ctx, http.StatusCreated, endpoint)
}

func (e EndpointController) DeleteEndpoint(ctx *gin.Context) {
	endpoint := e.endpointRepository.FindEndpoint()
	response.JSON(ctx, http.StatusOK, endpoint)
}

func (e EndpointController) GetListEndpoint(ctx *gin.Context) {
	endpoints := e.endpointRepository.FindEndpoint()
	response.JSON(ctx, http.StatusOK, endpoints)
}

func (e EndpointController) TestEndpoint(ctx *gin.Context) {
	var body querybuilder.TestEndpointRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err)
		return
	}
	endpointId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	endpoint := e.endpointRepository.FindBy(models.Endpoint{
		ID: endpointId,
	})
	workspace := e.workspaceRepository.FindByID(endpoint.WorkspaceID)

	endpointJSON, _ := json.Marshal(endpoint)
	endpointInput := gjson.Get(string(endpointJSON), "input")
	endpointFunctions := gjson.Get(string(endpointJSON), "functions")
	endpointResponse := gjson.Get(string(endpointJSON), "response")

	var functionResult = make(map[string]interface{})
	var responseOutput = make(map[string]interface{})
	for _, function := range endpointFunctions.Array() {
		functionName := function.Map()["name"]
		functionAs := function.Map()["as"]
		functionContext := function.Map()["context"]
		var query string
		switch functionName.String() {
		case string(models.MvpDboRawSql):
			query = mvprawsql.CreateQueryBuilder(endpointInput, function, body)
		case string(models.MvpDboGetby):
			dbo := functionContext.Map()["dbo"]
			tableId := dbo.Map()["id"]
			table, _ := e.tableRepository.FindTableById(uint64(tableId.Num))
			query = mvpgetby.CreateQueryBuilder(endpointInput, function, body, table.Name)
		}
		dsn := database.GetDNSByWorkspace(workspace)
		fmt.Println(query)
		result, err := database.PostgresQuery(dsn, query)
		if err != nil {
			panic(err)
		}

		switch functionName.String() {
		case string(models.MvpDboRawSql):
			if functionContext.Map()["dbo"].Map()["first"].Bool() && result != nil {
				functionResult[functionAs.Str] = result[0]
			} else {
				functionResult[functionAs.Str] = result
			}
		case string(models.MvpDboGetby):
			if result != nil {
				functionResult[functionAs.Str] = result[0]
			}
		}
	}

	for _, response := range endpointResponse.Array() {
		responseOutput[response.Map()["as"].Str] = functionResult[response.Map()["value"].Str]
	}
	fmt.Println(endpointResponse)
	res := make(map[string]interface{})
	res["output"] = responseOutput
	response.JSON(ctx, http.StatusOK, res)
}

func (e EndpointController) UpdateEndpoint(ctx *gin.Context) {
	var body EndpointInputRequest
	endpointId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err)
		return
	}

	var dataUpdate = models.Endpoint{
		ID: endpointId,
	}
	if body.Input != nil {
		input, _ := json.Marshal(body.Input)
		dataUpdate.Input = input
	}

	if body.Functions != nil {
		functions, _ := json.Marshal(body.Functions)
		dataUpdate.Functions = functions
	}

	if body.Response != nil {
		response, _ := json.Marshal(body.Response)
		dataUpdate.Response = response
	}

	endpoint := e.endpointRepository.UpdateEndpoint(&dataUpdate)

	response.JSON(ctx, http.StatusOK, endpoint)
}
