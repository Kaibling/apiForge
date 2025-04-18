package envelope

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/log"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/apiforge/route"
)

type Envelope struct {
	Success        bool               `json:"success"`
	RequestID      string             `json:"request_id"`
	Time           string             `json:"time"`
	Data           any                `json:"data"`
	Error          string             `json:"error,omitempty"`
	Errors         []string           `json:"errors,omitempty"`
	Message        string             `json:"message,omitempty"`
	HTTPStatusCode int                `json:"-"`
	Pagination     *params.Pagination `json:"pagination,omitempty"`
}

func New() *Envelope {
	return &Envelope{HTTPStatusCode: http.StatusOK}
}

func (e *Envelope) Render(w http.ResponseWriter, r *http.Request) error { //nolint: revive
	return nil
}

func (e *Envelope) SetResponse(resp any) *Envelope {
	e.Success = true
	e.Data = resp

	return e
}

func (e *Envelope) SetSuccess() *Envelope {
	e.Success = true

	return e
}

func (e *Envelope) SetPagination(p params.Pagination) *Envelope {
	e.Pagination = &p

	return e
}

func (e *Envelope) SetError(err apierror.HTTPError) *Envelope {
	e.Success = false
	e.HTTPStatusCode = err.HTTPStatus()
	e.Error = err.Error()
	e.Errors = err.Errors()

	return e
}

// func (e *Envelope) SetErrors(err apierror.HTTPError) *Envelope {
// 	e.Success = false
// 	e.HTTPStatusCode = err.HTTPStatus()
// 	e.Error = err.Error()
// 	e.Errors = err.Errors()

// 	return e
// }

func (e *Envelope) Finish(w http.ResponseWriter, r *http.Request, logger log.Writer) {
	e.Time = time.Now().Format(time.RFC3339)

	if !e.Success {
		logger.Warn(e.Error)
	}

	render.Status(r, e.HTTPStatusCode)
	route.Render(w, r, e)
}

func ReadEnvelope(r *http.Request) *Envelope {
	return ctxkeys.GetValue(r.Context(), ctxkeys.EnvelopeKey).(*Envelope) //nolint:forcetypeassert
}

func GetEnvelopeAndLogger(r *http.Request) (*Envelope, log.Writer, apierror.HTTPError) { //nolint: ireturn
	errs := []string{}

	l, ok := ctxkeys.GetValue(r.Context(), ctxkeys.LoggerKey).(log.Writer)
	if !ok {
		errs = append(errs, apierror.ErrContextMissingLogger.Error())
	}

	e, ok := ctxkeys.GetValue(r.Context(), ctxkeys.EnvelopeKey).(*Envelope)
	if !ok {
		errs = append(errs, apierror.ErrContextMissingEnvelope.Error())
	}

	if len(errs) > 0 {
		return nil, nil, apierror.NewMulti(apierror.ErrContextMissing, errs)
	}

	return e, l, nil
}
