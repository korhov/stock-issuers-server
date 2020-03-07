package list

import (
	"github.com/go-kit/kit/endpoint"
	pkgSecurities "github.com/korhov/stock-issuers-server/pkg/securities"

	"context"
)

func Endpoint(svcSecurities pkgSecurities.Service) endpoint.Endpoint {
	return func(ctx context.Context, raw interface{}) (resp interface{}, err error) {
		return svcSecurities.GetList(ctx, raw.(pkgSecurities.RequestList))
	}
}
