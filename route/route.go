package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ReadURLParam(key string, r *http.Request) string {
	return chi.URLParam(r, key)
}

// takes the http request an  decode the data into the given struct.
func ReadPostData(r *http.Request, a any) error {
	return json.NewDecoder(r.Body).Decode(&a)
}

func Render(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	if err := render.Render(w, r, v); err != nil {
		fmt.Println(err.Error()) //nolint: forbidigo
	}
}
