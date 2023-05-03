package controller

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"fast-api.io/internal/repository"
	"fast-api.io/models"
	"fast-api.io/modules/database"
	"fast-api.io/modules/http/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"gorm.io/datatypes"
)

var validate = validator.New()

type TableNewContentRequest struct {
	Data datatypes.JSON `json:"data"`
}

type TableNewColumnRequest struct {
	Name     string `json:"name" validate:"required"`
	DataType string `json:"data_type" validate:"required"`
}

type TableChangeColumnNameRequest struct {
	OldName string `json:"old_name" validate:"required"`
	NewName string `json:"new_name" validate:"required"`
}

type TableDropColumnRequest struct {
	Name string `json:"name" validate:"required"`
}

type TableRequest struct {
	Name string `json:"name" validate:"required"`
}

type TableUpdateContentRequest struct {
	Name  string      `json:"name" validate:"required"`
	Value interface{} `json:"value" validate:"required"`
}

type TableController struct {
	tableRepository     repository.TableRepository
	workspaceRepository repository.WorkspaceRepository
}

type TableColumn struct {
	Name     string      `json:"name"`
	DataType string      `json:"data_type"`
	Nullable bool        `json:"nullable"`
	Default  interface{} `json:"default"`
}

type UpdateViewRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type UpdateColumnRequest struct {
	OldColumn string `json:"old_column"`
	NewColumn string `json:"new_column"`
	DataType  string `json:"data_type"`
}

type InsertColumnRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

var dnsExample = database.GetDNSByWorkspace(&models.Workspace{})

func InitTableController(r *gin.RouterGroup, tableRepository *repository.TableRepository, workspaceRepository *repository.WorkspaceRepository) {
	controller := &TableController{
		tableRepository:     *tableRepository,
		workspaceRepository: *workspaceRepository,
	}

	routes := r.Group("projects/:workspaceId/tables")
	routes.GET("", controller.GetListTable)
	routes.POST("", controller.CreateTable)
	routes.GET(":tableName/views", controller.TableViews)
	routes.POST(":tableName/views", controller.InsertView)
	routes.PUT(":tableName/views/:viewId", controller.UpdateView)
	routes.DELETE(":tableName/views/:viewId", controller.DeleteView)
	routes.GET(":tableName/columns", controller.TableColumns)
	routes.POST(":tableName/columns", controller.InsertColumn)
	routes.PUT(":tableName/columns", controller.UpdateColumn)
	routes.DELETE(":tableName/columns/:columnName", controller.DeleteColumn)

}

func (e TableController) CreateTable(ctx *gin.Context) {
	var body TableRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err)
		return
	}

	var schema = fmt.Sprintf(`
		CREATE TABLE %s (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT NOW()
		);
	`, body.Name)
	result, dbErr := database.PostgresExec(dnsExample, schema)

	if dbErr != nil {
		response.Error(ctx, http.StatusBadRequest, dbErr)
		return
	}

	response.JSON(ctx, http.StatusCreated, result)
}

func (e TableController) GetListTable(ctx *gin.Context) {
	query := fmt.Sprintf(`SELECT table_name FROM information_schema.tables WHERE table_schema='public' ORDER BY table_name asc`)
	result, dbErr := database.PostgresQuery(dnsExample, query)
	var tables []map[string]interface{}

	if dbErr != nil {
		panic(dbErr)
	}

	for _, s := range result {
		tables = append(tables, map[string]interface{}{
			"name": string(s["table_name"].([]uint8)),
		})
	}
	response.JSON(ctx, http.StatusOK, tables)
}

func (e TableController) TableViews(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	fmt.Println(tableName)
	views, dbErr := database.PostgresQuery(dnsExample, fmt.Sprintf(`SELECT * FROM %s ORDER BY id`, tableName))
	if dbErr != nil {
		panic(dbErr)
	}
	response.JSON(ctx, http.StatusOK, views)
}

func (e TableController) InsertView(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	var body UpdateViewRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	columns := e.GetTableColumnInfo(tableName)

	var columnList []string = lo.Map(columns, func(column TableColumn, _ int) string {
		return fmt.Sprint(`"`, column.Name, `"`)
	})

	var columnValues []string = lo.Map(columns, func(_ TableColumn, _ int) string {
		return "default"
	})

	// INSERT INTO %s (id) VALUES (nextval('%s_id_seq'::regclass))
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s) RETURNING id`, tableName, strings.Join(columnList, ","), strings.Join(columnValues, ", "))
	fmt.Println(query)
	lastInsertId, dbErr := database.PostgresQueryRow(dnsExample, query)

	if dbErr != nil {
		response.JSON(ctx, http.StatusBadRequest, dbErr.Error())
		return
	}

	view := e.GetTableByID(tableName, *lastInsertId)

	response.JSON(ctx, http.StatusOK, view)
}

func (e TableController) UpdateView(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	viewId := ctx.Param("viewId")
	var body UpdateViewRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		response.JSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(body.Value)
	valueType := reflect.TypeOf(body.Value).String()
	var query string
	if valueType == "float64" {
		query = fmt.Sprintf(`UPDATE %s SET "%s" = %v WHERE id = %s`, tableName, body.Key, body.Value, viewId)
	} else {
		query = fmt.Sprintf(`UPDATE %s SET "%s" = '%v' WHERE id = %s`, tableName, body.Key, body.Value, viewId)

	}

	fmt.Println(query)

	result, dbErr := database.PostgresQuery(dnsExample, query)

	if dbErr != nil {
		response.JSON(ctx, http.StatusBadRequest, dbErr.Error())
		return
	}
	response.JSON(ctx, http.StatusOK, result)
}

func (e TableController) DeleteView(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	viewId := ctx.Param("viewId")
	var body UpdateViewRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = %v`, tableName, viewId)
	fmt.Println(query)
	result, dbErr := database.PostgresQuery(dnsExample, query)

	if dbErr != nil {
		panic(dbErr)
	}

	response.JSON(ctx, http.StatusOK, result)
}

func (e TableController) TableColumns(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	columns := e.GetTableColumnInfo(tableName)
	response.JSON(ctx, http.StatusOK, columns)
}

func (e TableController) UpdateColumn(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	var body UpdateColumnRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		panic(err)
	}

	if body.NewColumn != body.OldColumn {
		query := fmt.Sprintf(`ALTER TABLE %s RENAME COLUMN "%s" TO "%s"`, tableName, body.OldColumn, body.NewColumn)
		fmt.Println(query)
		_, dbErr := database.PostgresQuery(dnsExample, query)

		if dbErr != nil {
			response.JSON(ctx, http.StatusBadRequest, dbErr.Error())
			return
		}
	}
	response.JSON(ctx, http.StatusOK, tableName)
}

func (e TableController) InsertColumn(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	var body InsertColumnRequest
	ctx.ShouldBindJSON(&body)
	err := validate.Struct(body)
	if err != nil {
		panic(err)
	}

	query := fmt.Sprintf(`ALTER TABLE %s ADD "%s" %s`, tableName, body.Name, body.Type)
	fmt.Println(query)
	_, dbErr := database.PostgresQuery(dnsExample, query)

	if dbErr != nil {
		response.JSON(ctx, http.StatusBadRequest, dbErr.Error())
		return
	}
	response.JSON(ctx, http.StatusOK, tableName)
}

func (e TableController) DeleteColumn(ctx *gin.Context) {
	tableName := ctx.Param("tableName")
	columnName := ctx.Param("columnName")

	query := fmt.Sprintf(`ALTER TABLE %s DROP COLUMN "%s"`, tableName, columnName)
	fmt.Println(query)
	_, dbErr := database.PostgresQuery(dnsExample, query)

	if dbErr != nil {
		response.JSON(ctx, http.StatusBadRequest, dbErr.Error())
		return
	}
	response.JSON(ctx, http.StatusOK, tableName)
}

func (e TableController) GetTableColumnInfo(tableName string) []TableColumn {
	result, dbErr := database.PostgresQuery(dnsExample, fmt.Sprintf(`select *
	from INFORMATION_SCHEMA.COLUMNS
	where TABLE_NAME='%s' AND table_schema  != 'information_schema' ORDER BY ORDINAL_POSITION`, tableName))

	if dbErr != nil {
		panic(dbErr)
	}

	var columns []TableColumn

	for i, s := range result {
		fmt.Println(i)
		var nullable = false
		if s["is_nullable"].(string) == "YES" {
			nullable = true
		}
		columns = append(columns, TableColumn{
			Name:     string(s["column_name"].([]uint8)),
			DataType: s["data_type"].(string),
			Nullable: nullable,
			Default:  s["column_default"],
		})
	}

	return columns
}

func (e TableController) GetTableByID(tableName string, id int) map[string]interface{} {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = %v", tableName, id)
	result, dbErr := database.PostgresQuery(dnsExample, query)

	if dbErr != nil {
		panic(dbErr)
	}

	if result != nil {
		return result[0]
	}
	return nil
}
