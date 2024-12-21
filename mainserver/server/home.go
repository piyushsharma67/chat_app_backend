package server

import "github.com/gin-gonic/gin"

func (s *Server) Home(g *gin.Context){
	g.JSON(200, gin.H{"message":"Home reached"})
}