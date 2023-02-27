package middleware

import "net/http"

const (
	MethodParam  = "_method"
	MethodHeader = "X-HTTP-Method-Override"
)

type MethodGetter func(*http.Request) string

// The MethodOverride middleware checks for a HTTP method override in the request and
// uses it instead of the original method.
//
// For security reasons only `POST` can be overridden with `PUT`, `PATCH` or `DELETE`.
func MethodOverride(getters ...MethodGetter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				for _, getter := range getters {
					if m := getter(r); m != "" {
						if m == "PUT" || m == "PATCH" || m == "DELETE" {
							r.Method = m
						}
						break
					}

				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func MethodFromHeader(header string) MethodGetter {
	return func(r *http.Request) string {
		return r.Header.Get(header)
	}
}

func MethodFromForm(param string) MethodGetter {
	return func(r *http.Request) string {
		return r.FormValue(param)
	}
}
