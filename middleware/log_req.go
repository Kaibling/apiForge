package middleware

import (
	"net/http"
	"time"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/logging"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)

		finished := time.Now()
		duration := finished.Sub(start)
		url := r.URL.String()
		body := string(ctxkeys.GetValue(r.Context(), ctxkeys.ByteBodyKey).([]uint8))
		method := r.Method
		var username string
		if u, ok := ctxkeys.GetValue(r.Context(), ctxkeys.UserNameKey).(string); ok {
			username = u
		} else {
			username = "unauthenticated"
		}
		requestID := ctxkeys.GetValue(r.Context(), ctxkeys.RequestIDKey).(string)

		e := ctxkeys.GetValue(r.Context(), ctxkeys.EnvelopeKey).(*envelope.Envelope)
		//fmt.Printf("duration: %s\n", duration)
		logger := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(*logging.Logger)

		ld := logging.LogData{
			ReqId:          requestID,
			URL:            url,
			HttpStatusCode: e.HTTPStatusCode,
			Duration:       int(duration.Milliseconds()),
			UserName:       username,
			Method:         method,
			ResponseBody:   e,
		}
		// TODO only censor password and token
		if url != "/auth/login" {
			ld.RequestBody = body
		}
		logger.LogRequest(ld)
	})
}
