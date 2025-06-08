package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
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