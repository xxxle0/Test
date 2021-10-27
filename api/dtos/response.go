package dtos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	ErrorMsg   string      `json:"error_message"`
	StatusCode int         `json:"status_code"`
}

func HttpResponse(c *gin.Context, res Response) {
	c.JSON(res.StatusCode, gin.H{
		"status_code": res.StatusCode,
		"data":        res.Data,
		"message":     res.Message,
	})
}

func ErrorResponse(c *gin.Context, res Response) {
	c.JSON(res.StatusCode, gin.H{
		"status_code": res.StatusCode,
		"data":        res.Data,
		"error":       res.ErrorMsg,
	})
}

func BadRequest(c *gin.Context, res Response) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status_code": http.StatusBadRequest,
		"error":       res.ErrorMsg,
	})
}
