package middleware

import (
	"context"
	"net/http"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/lib/utils"
)

func InitEnvelope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqID string
		if val, ok := r.Header["X-Request-Id"]; ok {
			reqID = val[0]
		} else {
			reqID = utils.NewULID().String()
		}

		env := envelope.New()
		env.RequestID = reqID
		ctx := context.WithValue(r.Context(), ctxkeys.EnvelopeKey, env)
		ctx = context.WithValue(ctx, ctxkeys.RequestIDKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
