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
		fmt.Println("error occured1", token)
		utils.ResponseFormatter(g, http.StatusForbidden, false, nil, ErrUserForbidden)
		g.Abort()
		return
	}

	user, err := utils.ValidateToken(token)

	if err != nil {
		fmt.Println("error occured2", token)
		utils.ResponseFormatter(g, http.StatusForbidden, false, nil, ErrErrorOccurred)
		g.Abort()
		return
	}
	g.Set("user", user)
	g.Next()
}
