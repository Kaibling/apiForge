package middleware

import (
	"context"
	"net/http"

	"github.com/kaibling/apiforge/ctxkeys"
)

func AddContext(key ctxkeys.String, value any) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), key, value)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
