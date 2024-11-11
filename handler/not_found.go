package handler

import (
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	e, l := envelope.GetEnvelopeAndLogger(r)
	e.SetError(apierror.NotFound).Finish(w, r, l)
}
