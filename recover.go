package middleware

import (
	"net/http"
)

// Recover is a middleware that catches any panics in the call stack, calls
// recoverHandler and then responds with an internal server error.
func Recover(recoverHandler func(any)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// don't recover from http.ErrAbortHandler (see http.ErrAbortHandler docs)
					if err == http.ErrAbortHandler {
						panic(err)
					}

					recoverHandler(err)

					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
