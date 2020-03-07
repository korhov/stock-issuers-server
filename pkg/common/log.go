package common

import (
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"

	"context"
	"fmt"
)

func LogMiddleware(log *zap.SugaredLogger) endpoint.Middleware {
	return func(ep endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = ep(ctx, request)

			if err != nil {
				log.Errorw("error ocuried", "err", err)
				fmt.Printf("%+v \n", err)
			}

			return
		}
	}
}
