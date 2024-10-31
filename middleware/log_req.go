package middleware

import (
	"net/http"
	"time"

	"github.com/kaibling/apiforge/apictx"
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
		body := string(apictx.GetValue(r.Context(), "byte_body").([]uint8))
		method := r.Method
		var username string
		if u, ok := apictx.GetValue(r.Context(), "user_name").(string); ok {
			username = u
		} else {
			username = "unauthenticated"
		}
		requestID := apictx.GetValue(r.Context(), "request_id").(string)

		e := apictx.GetValue(r.Context(), "envelope").(*envelope.Envelope)
		//fmt.Printf("duration: %s\n", duration)
		logger := apictx.GetValue(r.Context(), "logger").(*logging.Logger)

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
