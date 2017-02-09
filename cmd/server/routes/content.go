package routes

import (
	"context"
	"net/http"
	"time"

	"zenithar.org/go/nikoniko/cmd/server/shared"

	"github.com/aqtrans/ctx-csrf"
	"zenithar.org/go/common/web/middleware/renderer"
	"zenithar.org/go/common/web/middleware/sessions"
)

// -----------------------------------------------------------------------------

const (
	DataNoRender int = iota
	DataHTML
	DataBinary
	DataText
	DataJSON
	DataJSONP
	DataXML
)

const (
	FlashErr  = "_flash_err"
	FlashWarn = "_flash_warn"
	FlashInfo = "_flash_info"
)

// -----------------------------------------------------------------------------

type Content struct {
	Type     int
	Status   int
	Template string // for DataHTML
	Callback string // for DataJSONP
	Data     interface{}
}

type contextKey string

func (c contextKey) String() string {
	return "zenithar.org/go/nikoniko/" + string(c)
}

var (
	contextKeyData = contextKey("data")
)

// FromContext gets the content from the context.
func FromContext(ctx context.Context) (*Content, bool) {
	content, ok := ctx.Value(contextKeyData).(*Content)
	return content, ok
}

// -----------------------------------------------------------------------------

func contentBeforeware(rw http.ResponseWriter, r *http.Request) context.Context {
	return context.WithValue(r.Context(), contextKeyData, &Content{
		Type:   DataHTML,
		Status: 200,
	})
}

func contentAfterware(rw http.ResponseWriter, r *http.Request) {
	ret, _ := FromContext(r.Context())
	renderer, _ := renderer.FromContext(r.Context())
	session, _ := sessions.FromContext(r.Context())

	// Set the CSRF token
	rw.Header().Set("X-CSRF-Token", csrf.Token(r.Context(), r))
	switch ret.Type {

	case DataNoRender:
		break

	case DataHTML:
		if ret.Template == "" {
			// guess we're not rendering anything
			break
		}

		if m, ok := ret.Data.(map[string]interface{}); ok {
			// Set the copyright on all pages
			m["copyright"] = time.Now().Year()

			// Add xsrf tokens
			m["xsrf_token"] = csrf.Token(r.Context(), r)
			m["xsrf_data"] = csrf.TemplateField(r.Context(), r)

			// Authentication
			m["authenticated"] = false

			// Environment
			if shared.Config.DevMode {
				m["environment"] = "dev"
			} else {
				m["environment"] = "prod"
			}

			// Add session flash stuff
			if f, has := session.Values[FlashErr]; has {
				m["flash_err"] = f
				delete(session.Values, FlashErr)
			}
			if f, has := session.Values[FlashWarn]; has {
				m["flash_warn"] = f
				delete(session.Values, FlashWarn)
			}
			if f, has := session.Values[FlashInfo]; has {
				m["flash_info"] = f
				delete(session.Values, FlashInfo)
			}
		}
		session.Save(r, rw)
		renderer.HTML(rw, ret.Status, ret.Template, ret.Data)
	case DataJSON:
		renderer.JSON(rw, ret.Status, ret.Data)
	case DataBinary:
		renderer.Data(rw, ret.Status, ret.Data.([]byte))
	case DataText:
		renderer.Text(rw, ret.Status, ret.Data.(string))
	case DataJSONP:
		renderer.JSONP(rw, ret.Status, ret.Callback, ret.Data)
	case DataXML:
		renderer.XML(rw, ret.Status, ret.Data)
	default:
		panic("no such data type")
	}
}

func ContentRenderMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		newCtx := contentBeforeware(rw, r)
		inner.ServeHTTP(rw, r.WithContext(newCtx))
		contentAfterware(rw, r.WithContext(newCtx))
	})
}
