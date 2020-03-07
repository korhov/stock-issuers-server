package tracing

import (
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"

	"context"
	"net/http"
	"sync"
	"time"
)

var tracerError error
var localTracer opentracing.Tracer
var syncTracer sync.Once

// Init creates jaeger tracer
func Init(lg *zap.SugaredLogger, cfg Config) (opentracing.Tracer, error) {
	lg.Infow("start jaeger", "service", cfg.ServiceName, "reporter", cfg.Reporter, "enabled", cfg.Enabled)
	syncTracer.Do(func() {
		jcfg := config.Configuration{
			Disabled: !cfg.Enabled,
			Sampler: &config.SamplerConfig{
				Type:  "const",
				Param: 1, // nolint
			},
			Reporter: &config.ReporterConfig{
				LogSpans:            false,
				BufferFlushInterval: 1 * time.Second, // nolint
				LocalAgentHostPort:  cfg.Reporter,
			},
			ServiceName: cfg.ServiceName,
		}

		tracer, _, err := jcfg.NewTracer()
		if err != nil {
			lg.Errorw(
				"error of starting jaeger",
				"service", cfg.ServiceName,
				"reporter", cfg.Reporter,
				"enabled", cfg.Enabled,
				"error", err,
			) // nolint
			tracerError = err
		}

		localTracer = tracer
	})

	return localTracer, tracerError
}

// StartSpan starts jaeger span
func StartSpan(ctx context.Context, t opentracing.Tracer, name string) (opentracing.Span, context.Context) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span = t.StartSpan(name, opentracing.ChildOf(span.Context()))
	} else {
		span = t.StartSpan(name)
	}

	return span, opentracing.ContextWithSpan(ctx, span)
}

// ErrorSpan starts error jaeger span
func ErrorSpan(span opentracing.Span, err error) {
	ext.Error.Set(span, true)
	span.LogKV("error", err.Error())
}

// HTTPToContext returns function that move http headers to context
// copied from go-kit
// create operationName from request
func HTTPToContext(tracer opentracing.Tracer) kithttp.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		// Try to join to a trace propagated in `req`.
		var span opentracing.Span
		wireContext, err := tracer.Extract(
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(req.Header),
		)

		if err != nil {
			span = tracer.StartSpan("HTTP " + req.Method + " " + req.URL.String())
		} else {
			span = tracer.StartSpan("HTTP "+req.Method+" "+req.URL.String(), ext.RPCServerOption(wireContext))
		}
		defer span.Finish()

		ext.HTTPMethod.Set(span, req.Method)
		ext.HTTPUrl.Set(span, req.URL.String())
		return opentracing.ContextWithSpan(ctx, span)
	}
}

// TraceServer is gokit middleware for tracing
// copied from go-kit
// dont rewrite operation name
func TraceServer(tracer opentracing.Tracer, operationName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			serverSpan := opentracing.SpanFromContext(ctx)
			if serverSpan == nil {
				// All we can do is create a new root span.
				serverSpan = tracer.StartSpan(operationName)
				defer serverSpan.Finish()
			}
			ext.SpanKindRPCServer.Set(serverSpan)
			ctx = opentracing.ContextWithSpan(ctx, serverSpan)
			res, err := next(ctx, request)

			if err != nil {
				serverSpan.SetTag("error", true)
				serverSpan.SetTag("error_msg", err.Error())
			}

			return res, err
		}
	}
}
