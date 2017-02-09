package routes

import (
	"net/http"

	"goji.io/middleware"
)

// NotFoundMiddleware is used to handle error when a route is not handled
func NotFoundMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		routeFound := middleware.Pattern(r.Context())
		if routeFound != nil {
			inner.ServeHTTP(rw, r)
			return
		}
		panic(ErrNotFound)
	})
}
