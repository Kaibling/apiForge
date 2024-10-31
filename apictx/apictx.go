package apictx

import (
	"context"
)

type String string

// func GetContext(key String, r *http.Request) interface{} {
// 	parameter := r.Context().Value(key)
// 	// if parameter == nil {
// 	// 	panic(apierrors.NewClientError(errors.New("context parameter '" + key + "' missing")))
// 	// }
// 	return parameter
// }

func GetValue(ctx context.Context, key String) any {
	return ctx.Value(key)
}
