package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/params"
)

func ParseQueryParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO hardcoded default values
		qp := params.Pagination{
			Limit: 20,
			Order: "asc",
		}

		if val, ok := r.URL.Query()["filter"]; ok {
			if len(val) > 0 {
				qp.Filter = val[0]
			}
		}
		if val, ok := r.URL.Query()["limit"]; ok {
			if len(val) > 0 {
				if l, err := strconv.Atoi(val[0]); err != nil {
					// TODO log error
					fmt.Println(err.Error())
				} else {
					qp.Limit = l
				}
			}
		}

		if val, ok := r.URL.Query()["order"]; ok {
			if len(val) > 0 {
				qp.Order = val[0]
			}
		}
		if val, ok := r.URL.Query()["before"]; ok {
			if len(val) > 0 {
				qp.Before = &val[0]
			}
		}
		if val, ok := r.URL.Query()["after"]; ok {
			if len(val) > 0 {
				qp.After = &val[0]
			}
		}

		ctx := context.WithValue(r.Context(), ctxkeys.QueryParamsKey, qp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
