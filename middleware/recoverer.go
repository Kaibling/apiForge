package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/logging"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				errMessage := fmt.Sprintf("Panic: %v\n%s", err, debug.Stack())
				logger := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(logging.Writer)
				logger.ErrorMsg(errMessage)

				e, ok := ctxkeys.GetValue(r.Context(), ctxkeys.EnvelopeKey).(*envelope.Envelope)
				if ok {
					e.SetError(apierror.ServerError).Finish(w, r, logger)
					return
				}
				// if no envelope available, send 500
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
