package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kaibling/apiforge/config"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/log"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		finished := time.Now()
		duration := finished.Sub(start)
		url := r.URL.String()
		method := r.Method

		body, ok := ctxkeys.GetValue(r.Context(), ctxkeys.ByteBodyKey).([]uint8)
		if !ok {
			fmt.Println("bytebody is missing in context") //nolint: forbidigo
		}

		var username string

		if u, ok := ctxkeys.GetValue(r.Context(), ctxkeys.UserNameKey).(string); ok {
			username = u
		} else {
			username = "unauthenticated"
		}

		requestID, ok := ctxkeys.GetValue(r.Context(), ctxkeys.RequestIDKey).(string)
		if !ok {
			fmt.Println("request_id is missing in context") //nolint: forbidigo
		}

		e, ok := ctxkeys.GetValue(r.Context(), ctxkeys.EnvelopeKey).(*envelope.Envelope)
		if !ok {
			fmt.Println("envelope is missing in context") //nolint: forbidigo
		}

		logger, ok := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(log.Writer)
		if !ok {
			fmt.Println("logger is missing in context") //nolint: forbidigo
		}

		ld := log.LogData{
			RequestID:      requestID,
			URL:            url,
			HTTPStatusCode: e.HTTPStatusCode,
			Duration:       int(duration.Milliseconds()),
			UserName:       username,
			Method:         method,
		}

		if config.LogResponseBody {
			ld.ResponseBody = e
		}

		// TODO only censor password and token
		if url != "/auth/login" {
			if config.LogRequestBody {
				ld.RequestBody = body
			}
		}

		logger.LogRequest(ld)
	})
}
