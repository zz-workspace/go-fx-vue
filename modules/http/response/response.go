package response

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	StatusCode int `json:"status_code"`
	Timestamp time.Time `json:"timestamp"`
	Error string `json:"error"`
}

type ValidationErrorResponse struct {
	StatusCode int `json:"status_code"`
	Timestamp time.Time `json:"timestamp"`
	Errors []ApiError `json:"errors"`
}

type SuccessResponse struct {
	StatusCode int `json:"status_code"`
	Timestamp time.Time `json:"timestamp"`
	Data interface{} `json:"data"`
}


func MsgForTag(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fieldError.Error() // default error
}


func getErrors(err error) []ApiError {
	errors := make([]ApiError, len(err.(validator.ValidationErrors)))
	for i, err := range err.(validator.ValidationErrors) {
		errors[i] = ApiError{err.StructField(), MsgForTag(err)}
	}
	return errors
}


func JSON(ctx *gin.Context, statusCode int, data interface{}) {
	result := SuccessResponse{}
	result.StatusCode = statusCode
	result.Timestamp = time.Now()
	result.Data = data
	ctx.JSON(statusCode, result)
}

func JSONWithPaging(ctx *gin.Context, statusCode int, data interface{}) {
	result := make(map[string]interface{})
	result["status_code"] = statusCode
	result["timestamp"] = time.Now()
	result["data"] = data
	// result["paging"] = paginator
	ctx.JSON(statusCode, result)
}

func Error(ctx *gin.Context, statusCode int, err error) {
	result := ErrorResponse{}
	result.StatusCode = statusCode
	result.Timestamp = time.Now()
	result.Error = err.Error()
	ctx.JSON(statusCode, result)
}

func ValidationError(ctx *gin.Context, statusCode int, err error) {
	result := ValidationErrorResponse{}
	result.StatusCode = statusCode
	result.Timestamp = time.Now()
	result.Errors = getErrors(err)
	ctx.JSON(statusCode, result)
}
