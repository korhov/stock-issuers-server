package common

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/korhov/stock-issuers-server/pkg/utils/text"

	"context"
	"net/http"
)

const (
	RequestIDKey    string = "request_id"
	LengthRequestID int    = 20
)

func RequestIDServerOption() kithttp.ServerOption {
	return kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
		requestID := request.Header.Get("request-id")
		if len(requestID) == 0 {
			requestID, _ = ctx.Value(RequestIDKey).(string)
		}

		if len(requestID) == 0 {
			requestID = text.GetRandomString(LengthRequestID)
		}

		return context.WithValue(ctx, RequestIDKey, requestID) // nolint
	})
}
