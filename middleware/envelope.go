package middleware

import (
	"context"
	"net/http"

	"github.com/kaibling/apiforge/apictx"
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

		// db := apictx.GetContext("db", r).(*gorm.DB)
		// username, ok := apictx.GetContext("username", r).(string)
		// if !ok {
		// 	username = "unauthenticated"
		// }
		// lr := gormrepo.NewLogRepo(db, username)
		// ls := services.NewLogService(lr)
		env := envelope.New()
		env.RequestID = reqID
		ctx := context.WithValue(r.Context(), apictx.String("envelope"), env)
		ctx = context.WithValue(ctx, apictx.String("request_id"), reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
