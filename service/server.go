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
	ReadTimeout       int
	ReadHeaderTimeout int
	WriteTimeout      int
	IdleTimeout       int
	EnableTLS         bool
	TLSCertPath       string
	TLSCertKeyPath    string
}

func setDefaultConfig(cfg ServerConfig) ServerConfig {
	if cfg.BindingIP == "" {
		cfg.BindingIP = "0.0.0.0"
	}

	if cfg.BindingPort == "" {
		cfg.BindingPort = "8080"
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
}

func New(cxt context.Context, cfg ServerConfig) *Server {
	return &Server{ctx: cxt, cfg: setDefaultConfig(cfg)}
}

func (s *Server) Start(r chi.Router, l log.Writer) error {
	r.Mount("/health/", api.AddReadyChecks())

	listeningStr := fmt.Sprintf("%s:%s", s.cfg.BindingIP, s.cfg.BindingPort)
	server := http.Server{
		Addr:              listeningStr,
		Handler:           r,
		ReadTimeout:       time.Duration(s.cfg.ReadTimeout) * time.Second,       // Max duration for reading entire request
		ReadHeaderTimeout: time.Duration(s.cfg.ReadHeaderTimeout) * time.Second, // Max duration for reading headers
		WriteTimeout:      time.Duration(s.cfg.WriteTimeout) * time.Second,      // Max duration before timing out writes
		IdleTimeout:       time.Duration(s.cfg.IdleTimeout) * time.Second,       // Max time to keep idle connections open
	}
	serverlogger := l.Named("api_server")

	go func() {
		<-s.ctx.Done()

		err := server.Shutdown(s.ctx)
		if err != nil {
			serverlogger.Error("server shutdown", err)
		}

		serverlogger.Info("shutting down api server")
	}()

	go func() {
		var err error
		if s.cfg.EnableTLS {
			err = server.ListenAndServeTLS(s.cfg.TLSCertPath, s.cfg.TLSCertKeyPath)
		} else {
			err = server.ListenAndServe()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverlogger.Error("http server serving", err)
		}

	}()
	serverlogger.Info("listening on " + listeningStr)

	return nil
}
