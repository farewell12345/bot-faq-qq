package core

import (
	"bot-faq-qq/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Data struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	status, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    status,
		Message: message,
		Data:    data,
	})
}
