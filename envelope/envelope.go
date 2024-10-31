package envelope

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/kaibling/apiforge/apictx"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"
)

type Envelope struct {
	Success        bool   `json:"success"`
	RequestID      string `json:"request_id"`
	Time           string `json:"time"`
	Response       any    `json:"response"`
	Error          string `json:"error,omitempty"`
	Message        string `json:"message,omitempty"`
	HTTPStatusCode int    `json:"-"`
}

func New() *Envelope {
	return &Envelope{HTTPStatusCode: http.StatusOK}
}
func (e *Envelope) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (e *Envelope) SetResponse(resp any) *Envelope {
	e.Success = true
	//e.Time = time.Now().Format(time.RFC822)
	e.Response = resp
	return e
}

func (e *Envelope) SetSuccess() *Envelope {
	e.Success = true
	//e.Time = time.Now().Format(time.RFC822)
	return e
}

func (e *Envelope) SetError(err apierror.HTTPError) *Envelope {
	e.Success = false
	e.HTTPStatusCode = err.HTTPStatus()
	e.Response = err.Error()
	return e
}

func (e *Envelope) Finish(w http.ResponseWriter, r *http.Request) {
	e.Time = time.Now().Format(time.RFC3339)
	//e.Time = time.Now().Format(time.RFC822)
	render.Status(r, e.HTTPStatusCode)
	route.Render(w, r, e)
}

func ReadEnvelope(r *http.Request) *Envelope {
	return apictx.GetValue(r.Context(), "envelope").(*Envelope)
}
