package handler

import (
	"net/http"

	"github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	e, l, err := envelope.GetEnvelopeAndLogger(r)
	if err != nil {
		e.SetError(err).Finish(w, r, l)

		return
	}

	e.SetError(apierror.ErrRouteNotFound).Finish(w, r, l)
}
