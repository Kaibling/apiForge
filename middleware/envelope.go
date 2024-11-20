package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/logging"
)

func InitEnvelope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqID string

		if val, ok := r.Header["X-Request-Id"]; ok {
			reqID = val[0]
		} else {
			reqID = utils.NewULID().String()
		}

		logger, ok := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(logging.Writer)
		if !ok {
			// TODO error
			fmt.Println("logger is missing in context") //nolint: forbidigo
		}

		logger.AddStringField("request_id", reqID)
		logger.Debug("request_id added")

		env := envelope.New()
		env.RequestID = reqID
		ctx := context.WithValue(r.Context(), ctxkeys.EnvelopeKey, env)
		ctx = context.WithValue(ctx, ctxkeys.RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
