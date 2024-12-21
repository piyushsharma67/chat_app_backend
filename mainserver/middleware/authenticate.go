package middleware

import (
	"mainserver/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticateuser(g *gin.Context) {
	token := g.GetHeader("Authorization")

	if token == "" {
		utils.ResponseFormatter(g, http.StatusForbidden, false, nil, utils.ErrorUserForbidden)
		g.Abort()
		return
	}

	user, err := utils.ValidateToken(token)

	if err != nil {
		utils.ResponseFormatter(g, http.StatusForbidden, false, nil, utils.ErrorOccurredAuthentication)
		g.Abort()
		return
	}
	g.Set("user", user)
	g.Next()
}
