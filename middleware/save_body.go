package middleware

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/kaibling/apiforge/ctxkeys"
)

func SaveBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(b))
		ctx := context.WithValue(r.Context(), ctxkeys.ByteBodyKey, b)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
