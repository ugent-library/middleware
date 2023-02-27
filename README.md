[![Go Reference](https://pkg.go.dev/badge/github.com/ugent-library/middleware.svg)](https://pkg.go.dev/github.com/ugent-library/middleware)

# ugent-library/middleware

Package middleware contains some generic middlewares and helper functions to
make composing middleware more readable.

## Install

```sh
go get -u github.com/ugent-library/middleware
```
## Examples

```go
handler = middleware.Apply(handler,
	middleware.Recover(func(err any) {
		logger.Error(err)
	}),
	middleware.If(config.Production, middleware.SetRequestID(uuid.NewString),
	middleware.MethodOverride(
		middleware.MethodFromHeader(middlewware.MethodHeader),
		middleware.MethodFromForm(middlewware.MethodParam),
	)
)
```
