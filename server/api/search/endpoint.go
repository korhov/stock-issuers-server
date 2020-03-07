package search

import (
	"github.com/go-kit/kit/endpoint"
	pkgSearch "github.com/korhov/stock-issuers-server/pkg/search"

	"context"
)

func Endpoint(svcResources pkgSearch.Service) endpoint.Endpoint {
	return func(ctx context.Context, raw interface{}) (resp interface{}, err error) {
		return svcResources.SearchRequest(ctx, raw.(pkgSearch.Request))
	}
}
