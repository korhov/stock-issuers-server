package search

import (
	pkgSearch "github.com/korhov/stock-issuers-server/pkg/search"

	"context"
	"net/http"
)

func RequestEncoder(svcResources pkgSearch.Service) func(ctx context.Context, req *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		return pkgSearch.Request{
			Ticket: r.URL.Query().Get("ticket"),
		}, nil
	}
}
