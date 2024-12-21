package server

import (
	"fmt"
	"mainserver/models"
	"mainserver/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Signin(r *gin.Context) {
	r.JSON(http.StatusOK, string([]byte("I am Healthy")))
}

func (s *Server) Signup(r *gin.Context) {
	var req models.SignupRequest

	if err := r.ShouldBindJSON(&req); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.queries.GetUser(r, req.Email)

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(user)

}

func (s *Server) AuthenticateToken(r *gin.Context) {
	token := r.GetHeader("Authorization")

	if token == "" {
		r.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	validated := utils.ValidateToken(token)

	if !validated {
		r.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	r.Next()
}
