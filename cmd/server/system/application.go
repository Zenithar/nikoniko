package system

import (
	"net/http"

	"zenithar.org/go/nikoniko/cmd/server/routes"
	"zenithar.org/go/nikoniko/cmd/server/shared"
	"zenithar.org/go/nikoniko/services"

	"github.com/aqtrans/ctx-csrf"
	"github.com/gorilla/sessions"
	"github.com/sebest/xff"
	urender "github.com/unrolled/render"
	goji "goji.io"
	"zenithar.org/go/common/cache"
	"zenithar.org/go/common/eventbus"
	mdlwr "zenithar.org/go/common/web/middleware"
	"zenithar.org/go/common/web/middleware/renderer"
	zsessions "zenithar.org/go/common/web/middleware/sessions"
)

// Application is the application instance
type Application interface {
	Router() http.Handler
	Start() error
	Stop() error
}

type baseApplication struct {
	users services.UserService
	votes services.VoteService

	config     *shared.Flags
	bus        eventbus.EventBus
	cacheStore cache.CacheStore
}

// -----------------------------------------------------------------------------

func (a *baseApplication) Router() http.Handler {
	// Create a new goji mux
	root := goji.NewMux()

	// Initialize middlewares
	xffhandler, _ := xff.Default()
	// Forwarded-For support
	root.Use(xffhandler.Handler)

	// Compression support
	root.Use(mdlwr.GzipHandler)

	// Session middleware
	store := sessions.NewCookieStore([]byte(a.config.CookieKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   a.config.CookieExpiration * 60,
		Secure:   a.config.CookieSecure,
		HttpOnly: true,
	}
	root.Use(zsessions.NewMiddleware("nikoniko", store))

	// Renderer
	opts := urender.Options{
		Directory:     "views",
		Layout:        "layout",
		Extensions:    []string{".tmpl", ".tpl"},
		IndentJSON:    a.config.DevMode,
		IndentXML:     a.config.DevMode,
		IsDevelopment: a.config.DevMode,
	}
	root.Use(renderer.NewMiddleware(urender.New(opts)))

	// Panic middleware
	root.Use(routes.PanicMiddleware)
	// NotFound middleware
	root.Use(routes.NotFoundMiddleware)

	// CSRF middleware
	root.Use(csrf.Protect([]byte(a.config.CookieKey),
		csrf.HttpOnly(true),
		csrf.Secure(a.config.CookieSecure),
		csrf.MaxAge(a.config.CookieExpiration*60),
		csrf.Path("/"),
		csrf.CookieName("_csrf"),
		csrf.FieldName("csrf_token"),
	))

	// Render data from context
	root.Use(routes.ContentRenderMiddleware)

	// Configure Router

	return root
}

// Start asynchronous tasks
func (a *baseApplication) Start() error {
	return nil
}

// Stop asynchronous tasks
func (a *baseApplication) Stop() error {
	return nil
}
