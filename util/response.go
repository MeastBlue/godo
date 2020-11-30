package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/model"
)

func SendJsonOK(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusOK, &data)
	sendJsonWithStatus(c, &formatData)
}

func SendJsonCreated(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusCreated, &data)
	sendJsonWithStatus(c, &formatData)
}

func SendJsonNoCotent(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusNoContent, &data)
	sendJsonWithStatus(c, &formatData)
}

func SendJsonUnauthorized(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusUnauthorized, &data)
	sendJsonWithStatus(c, &formatData)
}

func SendJsonForbidden(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusForbidden, &data)
	sendJsonWithStatus(c, &formatData)
}

func SendJsonNotFound(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusNotFound, &data)
	sendJsonWithStatus(c, &formatData)
}

func SendJsonUnprocessableEntity(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusUnprocessableEntity, &data)
	sendJsonWithStatus(c, &formatData)
}

func SendJsonError(c *gin.Context, data interface{}) {
	formatData := formatJsonResponse(http.StatusInternalServerError, &data)
	sendJsonWithStatus(c, &formatData)
}

func formatJsonResponse(code int, data interface{}) model.Response {
	return model.Response{
		Status: code,
		Data:   data,
	}
}

func sendJsonWithStatus(c *gin.Context, res *model.Response) {
	c.IndentedJSON(res.Status, gin.H{
		"code": res.Status,
		"data": res.Data,
	})
}
