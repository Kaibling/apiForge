package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaibling/apiforge/status"
)

func AddReadyChecks() chi.Router { //nolint: ireturn
	r := chi.NewRouter()
	r.Get("/live", fetchLiveStatus)
	r.Get("/ready", fetchReadyStatus)

	return r
}

func fetchLiveStatus(w http.ResponseWriter, r *http.Request) { //nolint: revive
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{\"status\":\"OK\"}")
}

func fetchReadyStatus(w http.ResponseWriter, r *http.Request) { //nolint: revive
	if !status.IsReady.Load().(bool) { //nolint: forcetypeassert
		// If not ready, respond with 503 Service Unavailable
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "{\"status\":\"not ready\"}")

		return
	}
	// If ready, respond with 200 OK
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{\"status\":\"ready\"}")
}
