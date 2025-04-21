package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kaibling/apiforge/api"
	"github.com/kaibling/apiforge/log"
)

type ServerConfig struct {
	BindingIP         string
	BindingPort       string
	LogLevel          string
	ReadTimeout       int
	ReadHeaderTimeout int
	WriteTimeout      int
	IdleTimeout       int
}

func setDefaultConfig(cfg ServerConfig) ServerConfig {
	if cfg.BindingIP == "" {
		cfg.BindingIP = "0.0.0.0"
	}
	if cfg.BindingPort == "" {
		cfg.BindingPort = "8080"
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = 5
	}
	if cfg.ReadHeaderTimeout == 0 {
		cfg.ReadHeaderTimeout = 2
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = 10
	}
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = 15
	}
	return cfg
}

type Server struct {
	ctx context.Context
	cfg ServerConfig
	l   log.Writer
}

func New(cxt context.Context, cfg ServerConfig) *Server {
	return &Server{ctx: cxt, cfg: setDefaultConfig(cfg)}
}

func (s *Server) Start(r chi.Router) error {
	r.Mount("/", api.AddReadyChecks())

	// if s.l == nil {
	// 	s.l = zap.New(s.cfg.LogLevel,s.cfg.)
	// }

	listeningStr := fmt.Sprintf("%s:%s", s.cfg.BindingIP, s.cfg.BindingPort)
	server := http.Server{
		Addr:              listeningStr,
		Handler:           r,
		ReadTimeout:       time.Duration(s.cfg.ReadTimeout) * time.Second,       // Max duration for reading entire request
		ReadHeaderTimeout: time.Duration(s.cfg.ReadHeaderTimeout) * time.Second, // Max duration for reading headers
		WriteTimeout:      time.Duration(s.cfg.WriteTimeout) * time.Second,      // Max duration before timing out writes
		IdleTimeout:       time.Duration(s.cfg.IdleTimeout) * time.Second,       // Max time to keep idle connections open
	}

	go func() {
		<-s.ctx.Done()

		err := server.Shutdown(s.ctx)
		if err != nil {
			s.l.Error("server shutdown", err)
		}

		s.l.Info("shutting down api server")
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.l.Error("http server serving", err)
		}
	}()
	s.l.Info("listening on " + listeningStr)

	return nil
}
