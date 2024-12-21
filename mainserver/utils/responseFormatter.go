package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  error       `json:"error,omitempty"`
}

func ResponseFormatter(r *gin.Context, statusCode int, status bool, data interface{}, errMessage error) {
	response := Response{
		Status: status,
	}

	if status {
		response.Data = data
	} else {
		response.Error = errMessage
	}

	r.JSON(statusCode, response)
}
