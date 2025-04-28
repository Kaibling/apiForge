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

const requestIDHeader = "x-request-id"

func InitEnvelope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(requestIDHeader)
		if reqID == "" {
			reqID = utils.NewULID().String()
		}

		l, ok := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(log.Writer)
		if !ok {
			// TODO error
			fmt.Println("logger is missing in context") //nolint: forbidigo
		}
		reqLogger := l.New("request", log.NewField("request_id", reqID))

		env := envelope.New()
		env.RequestID = reqID

		ctx := context.WithValue(r.Context(), ctxkeys.EnvelopeKey, env)
		ctx = context.WithValue(ctx, ctxkeys.RequestIDKey, reqID)
		ctx = context.WithValue(ctx, ctxkeys.LoggerKey, reqLogger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
