package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
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

func (s *Server) StartBlocking(r chi.Router) {
	if s.l == nil {
		s.l = logging.New(zap.New())
	}
	cl := make(chan os.Signal, 1)
	signal.Notify(cl, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	listeningStr := fmt.Sprintf("%s:%s", s.cfg.BindingIP, s.cfg.BindingPort)

	server := http.Server{Addr: listeningStr, Handler: r}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	done := make(chan bool, 1)

	go func() {
		<-done
		s.l.LogLine("shutown api server")
		// Shutdown signal with grace period of 5 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 5*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				slog.Warn("graceful shutdown timed out.. forcing exit.")
			}
		}()
		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			slog.Error(err.Error())
		}
		slog.Info("server shutting down")
		serverStopCtx()
		cancel()
	}()

	go func() {
		if err := http.ListenAndServe(listeningStr, r); err != nil {
			//if err := http.ListenAndServe(listeningStr, cr); err != nil {
			slog.Error(err.Error())
		}
	}()
	s.l.LogLine(fmt.Sprintf("listening on %s", listeningStr))
	<-cl

	done <- true
	// todo loop through cancel contexes
	//scheduleCancel()
}
