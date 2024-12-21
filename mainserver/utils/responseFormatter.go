package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func ResponseFormatter(r *gin.Context, statusCode int, status bool, data interface{}, errMessage error) {
	response := Response{
		Status: status,
	}

	if status {
		response.Data = data
	} else {
		response.Error = errMessage.Error()
	}
	fmt.Println(response)
	r.JSON(statusCode, response)
}
