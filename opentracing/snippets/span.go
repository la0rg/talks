type Span interface {
	Finish()

	Context() SpanContext

	SetOperationName(operationName string) Span

	SetTag(key string, value interface{}) Span

	LogFields(fields ...log.Field)

	SetBaggageItem(restrictedKey, value string) Span

	BaggageItem(restrictedKey string) string

	Tracer() Tracer
}
