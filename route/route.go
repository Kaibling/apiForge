package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// type Route struct {
// 	r *chi.Mux
// }

// func New(baseUrl string) *Route {
// 	return &Route{r: chi.NewRouter()}
// }

// func (r *Route) DeleteFunc(url string, f func(w http.ResponseWriter, r *http.Request)) {
// 	r.r.Delete(url, f)
// }

// func (r *Route) PostFunc(url string, f func(w http.ResponseWriter, r *http.Request)) {
// 	r.r.Post(url, f)
// }
// func (r *Route) GetFunc(url string, f func(w http.ResponseWriter, r *http.Request)) {
// 	r.r.Get(url, f)
// }

// func (r *Route) AddRoute(url string, subRoute *Route) {
// 	r.r.Mount(url, subRoute.r)
// }

// func (r *Route) R() *chi.Mux {
// 	return r.r
// }

// func (r *Route) AddMiddleware(middlewares func(http.Handler) http.Handler) {
// 	r.r.Use(middlewares)
// }

// func (r *Route) AddContext(key apictx.String, value any) {
// 	r.R().Use(
// 		func(next http.Handler) http.Handler {
// 			fn := func(w http.ResponseWriter, r *http.Request) {
// 				ctx := context.WithValue(r.Context(), key, value)
// 				fmt.Println(key)
// 				next.ServeHTTP(w, r.WithContext(ctx))
// 			}
// 			return http.HandlerFunc(fn)
// 		})
// }

func ReadUrlParam(key string, r *http.Request) string {
	return chi.URLParam(r, key)
}

// takes the http request an  decode the data into the given struct
func ReadPostData(r *http.Request, a any) error {
	return json.NewDecoder(r.Body).Decode(&a)
}

func Render(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	if err := render.Render(w, r, v); err != nil {
		fmt.Println(err.Error())
		//log.Error(r.Context(), err)
	}
}
