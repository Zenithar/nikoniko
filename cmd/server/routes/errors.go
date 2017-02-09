package routes

import (
	"net/http"

	"zenithar.org/go/common/web/middleware/renderer"

	errors "gopkg.in/errors.v0"
)

var (
	ErrNotFound         = errors.New("nikoniko: path not found")
	ErrMethodNotAllowed = errors.New("nikoniko: method not allowed")
	ErrBadRequest       = errors.New("nikoniko: bad request")
	ErrForbidden        = errors.New("nikoniko: forbidden")
	ErrPermission       = errors.New("nikoniko: permission error")
)

// PanicMiddleware is the panic handler
func PanicMiddleware(inner http.Handler) http.Handler {
	panicHandler := &PanicHandler{}
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				panicHandler.ServeHTTP(err, rw, r)
			}
		}()
		inner.ServeHTTP(rw, r)
	})
}

// PanicHandler is a middleware used to intercept error
type PanicHandler struct{}

func (p *PanicHandler) ServeHTTP(ex interface{}, rw http.ResponseWriter, r *http.Request) {
	switch ex {
	case ErrNotFound:
		p.err404(rw, r)
	case ErrMethodNotAllowed:
		p.err405(rw, r)
	case ErrBadRequest:
		p.err400(rw, r)
	case ErrForbidden:
		p.err403(rw, r)
	case ErrPermission:
		p.err550(rw, r)
	default:
		p.err500(ex, rw, r)
	}
}

func (p *PanicHandler) err404(rw http.ResponseWriter, r *http.Request) {
	renderer, _ := renderer.FromContext(r.Context())
	renderer.HTML(rw, 404, "errors/404", nil)
}

func (p *PanicHandler) err400(rw http.ResponseWriter, r *http.Request) {
	renderer, _ := renderer.FromContext(r.Context())
	renderer.HTML(rw, 400, "errors/400", nil)
}

func (p *PanicHandler) err403(rw http.ResponseWriter, r *http.Request) {
	renderer, _ := renderer.FromContext(r.Context())
	renderer.HTML(rw, 403, "errors/403", nil)
}

func (p *PanicHandler) err405(rw http.ResponseWriter, r *http.Request) {
	renderer, _ := renderer.FromContext(r.Context())
	renderer.HTML(rw, 405, "errors/405", nil)
}

func (p *PanicHandler) err500(ex interface{}, rw http.ResponseWriter, r *http.Request) {
	renderer, _ := renderer.FromContext(r.Context())
	renderer.HTML(rw, 500, "errors/500", nil)
}

func (p *PanicHandler) err550(rw http.ResponseWriter, r *http.Request) {
	renderer, _ := renderer.FromContext(r.Context())
	renderer.HTML(rw, 550, "errors/550", nil)
}
