package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/log"
)

func InitEnvelope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqID string

		if val, ok := r.Header["X-Request-Id"]; ok {
			reqID = val[0]
		} else {
			reqID = utils.NewULID().String()
		}

		l, ok := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(log.Writer)
		if !ok {
			// TODO error
			fmt.Println("logger is missing in context") //nolint: forbidigo
		}

		l.With(log.NewField("request_id", reqID)).Debug("request_id added")

		env := envelope.New()
		env.RequestID = reqID
		ctx := context.WithValue(r.Context(), ctxkeys.EnvelopeKey, env)
		ctx = context.WithValue(ctx, ctxkeys.RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
