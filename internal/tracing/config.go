package tracing

// Config is config for jaeger tracer
type Config struct {
	Reporter    string
	Enabled     bool
	ServiceName string
}
