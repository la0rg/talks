// Tracer is a simple, thin interface for Span creation and SpanContext
// propagation.
type Tracer interface {
	StartSpan(operationName string, opts ...StartSpanOption) Span

	Inject(sm SpanContext, format interface{}, carrier interface{}) error

	Extract(format interface{}, carrier interface{}) (SpanContext, error)
}
