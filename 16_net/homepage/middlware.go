package homepage

import "net/http"

// This type is just to make code more readable
type Middleware func(http.HandlerFunc) http.HandlerFunc

func with(handler http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	if len(middleware) < 1 {
		return handler
	}

	wrapped := handler

	// loop in reverse to preserve middleware order
	for i := len(middleware) - 1; i >= 0; i-- {
		wrapped = middleware[i](wrapped)
	}

	return wrapped
}
