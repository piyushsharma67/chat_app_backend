package server

import (
	"database/sql"
	"errors"
	"mainserver/models"
	"mainserver/schema"
	"mainserver/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationResponse struct {
	User  schema.User `json:"user"`
	Token string      `json:"token"`
}

func (s *Server) Signup(r *gin.Context) {
	var req models.SignupRequest

	if err := r.ShouldBindJSON(&req); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.queries.GetUser(r, req.Email)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			// Handle case where user does not exist
			utils.ResponseFormatter(r, http.StatusNotFound, true, nil, err)
		}
	}

	if user.ID == 0 {
		user, err = s.queries.CreateUser(r, schema.CreateUserParams{Name: req.Username, Email: req.Email, Password: req.Password})
	}

	if err != nil {
		utils.ResponseFormatter(r, http.StatusBadRequest, true, nil, err)
	}

	token, err := utils.GenerateToken(int(user.ID), user.Email)

	if err != nil {
		utils.ResponseFormatter(r, http.StatusBadRequest, true, nil, err)
	}
	var response AuthenticationResponse

	response.User = user
	response.Token = token

	utils.ResponseFormatter(r, http.StatusOK, true, response, nil)
}

func (s *Server) Health(g *gin.Context) {
	utils.ResponseFormatter(g, http.StatusOK, true, string([]byte("I am healthy")), nil)
}

func (s *Server) Signin(r *gin.Context) {

}
