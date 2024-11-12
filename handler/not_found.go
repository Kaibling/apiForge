package handler

import (
	"fmt"
	"net/http"

	"github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	e, l, err := envelope.GetEnvelopeAndLogger(r)
	if err != nil {
		fmt.Println(err.Error()) //nolint: forbidigo
	}

	e.SetError(apierror.ErrNotFound).Finish(w, r, l)
}
