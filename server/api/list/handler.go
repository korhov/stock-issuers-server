package list

import (
	kithttp "github.com/go-kit/kit/transport/http"
	pkgCommon "github.com/korhov/stock-issuers-server/pkg/common"
	pkgSecurities "github.com/korhov/stock-issuers-server/pkg/securities"
	"github.com/korhov/stock-issuers-server/server/middleware"
	"go.uber.org/zap"

	"net/http"
)

func Handler(log *zap.SugaredLogger, svcSecurities pkgSecurities.Service) http.Handler {
	options := pkgCommon.DefaultServerOpts(log)
	options = append(options, middleware.APIKeyServerOption())
	ep := pkgCommon.LogMiddleware(log)(Endpoint(svcSecurities))
	return kithttp.NewServer(
		ep,
		RequestEncoder(svcSecurities),
		pkgCommon.JSONResponseEncoder(),
		options...,
	)
}
