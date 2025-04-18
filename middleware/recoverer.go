package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/log"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo check if, context is needed for defer
		defer func() { //nolint: contextcheck
			if err := recover(); err != nil {
				errMessage := fmt.Sprintf("Panic: %v\n%s", err, debug.Stack())
				nerr := errors.New(errMessage)

				logger, ok := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(log.Writer)
				if !ok {
					fmt.Println("logger is missing in context") //nolint: forbidigo
				}

				logger.Error(errMessage, nerr)

				e, ok := ctxkeys.GetValue(r.Context(), ctxkeys.EnvelopeKey).(*envelope.Envelope)
				if ok {
					e.SetError(apierror.ErrServerError).Finish(w, r, logger)

					return
				}
				// if no envelope available, send 500
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
