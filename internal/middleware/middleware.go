package middleware

import (
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/internal/logger"
)

var (
	log = logger.CreateNew()
)

type Middleware func(http.Handler) http.Handler

func CreateMiddlewareStack(ms ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}

		return next
	}
}

func StripSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		log.Printf("Requested URL: %s\n", url)
		if url != "" && len(url) > 1 && url[len(url)-1] == '/' {
			url = url[:len(url)-1]
			r.URL.Path = url
			log.Printf("URL after trimming: %s\n", r.URL.Path)
		}
		next.ServeHTTP(w, r)
	})
}
