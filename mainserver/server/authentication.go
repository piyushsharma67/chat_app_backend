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

func (s *Server) Health(g *gin.Context) {
	utils.ResponseFormatter(g, http.StatusOK, true, string([]byte("I am healthy")), nil)
}

type AuthenticationResponse struct {
	User  schema.User `json:"user"`
	Token string      `json:"token"`
}

// controller for user signup

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

	if user.Email != "" {
		utils.ResponseFormatter(r, http.StatusOK, false, nil, utils.ErrorUserAlreadyExists)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		utils.ResponseFormatter(r, http.StatusInternalServerError, false, nil, utils.ErrErrorOccurred)
		return
	}

	if user.ID == 0 {
		user, err = s.queries.CreateUser(r, schema.CreateUserParams{Name: req.Username, Email: req.Email, Password: hashedPassword})
	}

	if err != nil {
		utils.ResponseFormatter(r, http.StatusBadRequest, true, nil, err)
		return
	}

	token, err := utils.GenerateToken(int(user.ID), user.Email)

	if err != nil {
		utils.ResponseFormatter(r, http.StatusBadRequest, true, nil, err)
		return
	}
	var response AuthenticationResponse
	user.Password = ""

	response.User = user
	response.Token = token

	utils.ResponseFormatter(r, http.StatusOK, true, response, nil)
}

func (s *Server) Signin(r *gin.Context) {
	var req models.SigninRequest

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

	isvalidPassword := utils.ComparePasswords(user.Password, req.Password)

	if !isvalidPassword {
		utils.ResponseFormatter(r, http.StatusBadRequest, false, nil, utils.ErrorInvalidpassword)
		return
	}

	token, err := utils.GenerateToken(int(user.ID), user.Email)

	if err != nil {
		utils.ResponseFormatter(r, http.StatusBadRequest, true, nil, err)
		return
	}
	var response AuthenticationResponse

	response.User = user
	response.Token = token

	utils.ResponseFormatter(r, http.StatusOK, true, response, nil)

}
