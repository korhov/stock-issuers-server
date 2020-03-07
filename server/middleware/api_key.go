package middleware

import (
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"

	"context"
	"fmt"
	"net/http"
)

const APIKeyHeader = "api-key"
const APIKeyQueryParam = "api_key"
const APIKeyContextParam = "api_key"

type HTTPCodableError struct {
	err  error
	code int
}

func NewError(msg string, code int) HTTPCodableError {
	return HTTPCodableError{
		err:  errors.New(msg),
		code: code,
	}
}

func (err HTTPCodableError) Error() string {
	return err.err.Error()
}

func (err HTTPCodableError) GetCode() int {
	return err.code
}

func APIKey(expectedAPIKey string) endpoint.Middleware {
	return func(ep endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			apiKey := ctx.Value(APIKeyContextParam)
			apiKeyStr := apiKey.(string)

			fmt.Println("from middleware", apiKeyStr)

			if len(apiKeyStr) == 0 {
				return nil, NewError("api key cannot be empty", http.StatusUnauthorized)
			}

			if apiKeyStr != expectedAPIKey {
				return nil, NewError("bad api key", http.StatusForbidden)
			}

			return ep(ctx, request)
		}
	}
}

func APIKeyServerOption() kithttp.ServerOption {
	return kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
		apiKey := request.Header.Get(APIKeyHeader)

		if len(apiKey) == 0 {
			apiKey = request.URL.Query().Get(APIKeyQueryParam)
		}

		return context.WithValue(ctx, APIKeyContextParam, apiKey) // nolint
	})
}
