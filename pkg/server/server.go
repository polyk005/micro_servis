package server

import (
	"context"
	"net/http"
	"time"
	
	"github.com/spf13/viper"
	"github.com/polyk005/micro_servis/pkg/queue"
	"github.com/polyk005/micro_servis/pkg/handler"
	"github.com/polyk005/micro_servis/pkg/repository"
)

type Server struct {
	httpServer *http.Server
	repo      repository.Repository
}

func NewServer(repo repository.Repository) *Server {
	return &Server{
		repo: repo,
	}
}

func (s *Server) Run(port string) error {
	q := queue.NewRedisQueue(
		viper.GetString("redis.addr"),
		viper.GetString("redis.stream"),
	)

	router := http.NewServeMux()
	taskHandler := handler.NewTaskHandler(q)
	router.HandleFunc("/tasks", taskHandler.CreateTask)
	
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
} 