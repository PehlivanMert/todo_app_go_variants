package utils

import (
	"net/http"
	"todo-app/dto"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func CreatedResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, dto.APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func PaginatedSuccessResponse(c *gin.Context, data interface{}, total int64, limit, offset int, message string) {
	c.JSON(http.StatusOK, dto.PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta: dto.MetaData{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		},
	})
}

func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message, "Bad Request")
}

func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, "Not Found")
}

func InternalServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message, "Internal Server Error")
}
