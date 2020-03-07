package routes

import (
	"github.com/go-chi/chi"
	pkgSearch "github.com/korhov/stock-issuers-server/pkg/search"
	pkgSecurities "github.com/korhov/stock-issuers-server/pkg/securities"
	apiList "github.com/korhov/stock-issuers-server/server/api/list"
	apiSearch "github.com/korhov/stock-issuers-server/server/api/search"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func apiRoutes(r chi.Router, logger *zap.SugaredLogger, tracer opentracing.Tracer) {
	svcSecurities := pkgSecurities.NewService(logger, tracer)
	svcResources := pkgSearch.NewService(logger, tracer, svcSecurities)

	r.Get("/search", apiSearch.Handler(logger, svcResources).ServeHTTP)
	r.Get("/list", apiList.Handler(logger, svcSecurities).ServeHTTP)
}
