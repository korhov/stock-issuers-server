package list

import (
	pkgSecurities "github.com/korhov/stock-issuers-server/pkg/securities"

	"context"
	"net/http"
)

func RequestEncoder(svcSecurities pkgSecurities.Service) func(ctx context.Context, req *http.Request) (interface{}, error) { // nolint
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		req := pkgSecurities.RequestList{}
		req.InstrumentType = r.URL.Query().Get("InstrumentType")
		return req, nil
	}
}
