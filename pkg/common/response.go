package common

import (
	kitHttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"context"
	"encoding/json"
	"net/http"
)

type Response struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Errors    []Error     `json:"errors"`
	RequestID string      `json:"request_id"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func JSONResponseEncoder() func(ctx context.Context, w http.ResponseWriter, data interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, data interface{}) error {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		requestID, _ := ctx.Value(RequestIDKey).(string)
		resp := Response{
			RequestID: requestID,
			Success:   true,
			Data:      data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			return err
		}

		return nil
	}
}

func GetDefaultErrorEncoder(lg *zap.SugaredLogger) func(ctx context.Context, err error, w http.ResponseWriter) {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		httpStatus := http.StatusInternalServerError
		errorCode := "undefined"
		requestID, _ := ctx.Value(RequestIDKey).(string)

		/*if ctxErr, ok := err.(pkgErrors.ContextError); ok {
			httpStatus = ctxErr.ValueInt(pkgErrors.HttpStatus, http.StatusOK)
			errorCode = ctxErr.ValueString(pkgErrors.Code, "undefined")
		}*/

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(httpStatus)

		res := Response{
			Success:   false,
			RequestID: requestID,
			Errors: []Error{{
				Message: err.Error(),
				Code:    errorCode,
			}},
		}
		_ = json.NewEncoder(w).Encode(res) // @todo: ...
	}
}

func DefaultServerOpts(log *zap.SugaredLogger) (opts []kitHttp.ServerOption) {
	return []kitHttp.ServerOption{
		RequestIDServerOption(),
		kitHttp.ServerErrorEncoder(GetDefaultErrorEncoder(log)),
	}
}
