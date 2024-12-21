package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct{
	db *sqlx.DB
}



func New(db *sqlx.DB) *Server{
	s := &Server{}

	s.db=db

	return s
}

func (server *Server)SetupRoutes(r *gin.Engine)*gin.Engine{
	
	publicRoutes:=r.Group("/public")
	publicRoutes.GET("/signin",server.Signin)
	publicRoutes.POST("/signup",server.Signup)

	privateRoutes:=r.Group("/private")
	privateRoutes.Use(server.AuthenticateToken)

	privateRoutes.GET("/home",server.Home)

	return r
}


func (s *Server) Start()error{

	r := gin.Default()

	r=s.SetupRoutes(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server error", "err", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}

	slog.Info("Server exiting")
	return nil
}