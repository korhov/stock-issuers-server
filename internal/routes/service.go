package routes

import (
	"github.com/go-chi/chi"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"net/http"
)

func Routes(logger *zap.SugaredLogger, tracer opentracing.Tracer) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	r.Route("/api/v1", func(r chi.Router) {
		apiRoutes(r, logger, tracer)
	})

	return r
}
