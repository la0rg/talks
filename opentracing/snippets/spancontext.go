type SpanContext interface {
	ForeachBaggageItem(handler func(k, v string) bool)
}
