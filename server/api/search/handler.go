package search

import (
	kithttp "github.com/go-kit/kit/transport/http"
	pkgCommon "github.com/korhov/stock-issuers-server/pkg/common"
	pkgSearch "github.com/korhov/stock-issuers-server/pkg/search"
	"go.uber.org/zap"

	"net/http"
)

func Handler(log *zap.SugaredLogger, svcResources pkgSearch.Service) http.Handler {
	options := pkgCommon.DefaultServerOpts(log)
	ep := pkgCommon.LogMiddleware(log)(Endpoint(svcResources))
	return kithttp.NewServer(
		ep,
		RequestEncoder(svcResources),
		pkgCommon.JSONResponseEncoder(),
		options...,
	)
}
