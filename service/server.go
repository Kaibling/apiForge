package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaibling/apiforge/api"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/logging/zap"
)

type ServerConfig struct {
	BindingIP   string
	BindingPort string
}

type Server struct {
	ctx context.Context
	cfg ServerConfig
	l   *logging.Logger
}

func New(cxt context.Context, cfg ServerConfig) *Server {
	return &Server{ctx: cxt, cfg: cfg}
}

func (s *Server) AddCustomLogger(lw logging.LogWriter) {
	s.l = logging.New(lw)
}

func (s *Server) Start(r chi.Router) error {
	r.Mount("/", api.AddReadyChecks())

	if s.l == nil {
		s.l = logging.New(zap.New())
	}

	listeningStr := fmt.Sprintf("%s:%s", s.cfg.BindingIP, s.cfg.BindingPort)
	server := http.Server{Addr: listeningStr, Handler: r}

	go func() {
		<-s.ctx.Done()
		err := server.Shutdown(s.ctx)
		if err != nil {
			s.l.LogLine(err.Error())
		}
		s.l.LogLine("shutting down api server")
	}()

	go func() {
		if err := http.ListenAndServe(listeningStr, r); err != nil {
			slog.Error(err.Error())
			return
		}
	}()
	s.l.LogLine(fmt.Sprintf("listening on %s", listeningStr))
	return nil
}
