package middleware

import (
	"errors"
	"fmt"
	"mainserver/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrUserForbidden = errors.New("login required")
	ErrErrorOccurred = errors.New("error occurred during authentication")
)

func Authenticateuser(g *gin.Context) {
	token := g.GetHeader("Authorization")

	if token == "" {
		utils.ResponseFormatter(g, http.StatusForbidden, false, nil, ErrUserForbidden)
	}

	user, err := utils.ValidateToken(token)

	if err != nil {
		fmt.Println("error occured")
		utils.ResponseFormatter(g, http.StatusForbidden, false, nil, ErrErrorOccurred)
		g.Abort()
	}
	g.Set("user", user)
	g.Next()
}
