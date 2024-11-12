package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/params"
)

func ParsePagination(next http.Handler) http.Handler { //nolint: gocognit,cyclop
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO hardcoded default values
		defaultLimit := 20

		qp := params.Pagination{
			Limit: defaultLimit,
			Order: "ASC",
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
					fmt.Println(err.Error()) //nolint: forbidigo
				} else {
					qp.Limit = l
				}
			}
		}

		if val, ok := r.URL.Query()["order"]; ok {
			if len(val) > 0 {
				p := strings.ToUpper(val[0])
				if p == "DESC" || p == "ASC" {
					qp.Order = p
				}
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

		ctx := context.WithValue(r.Context(), ctxkeys.PaginationKey, qp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
