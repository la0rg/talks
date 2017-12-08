package opentracing

import "context"

func GlobalTracer() Tracer {}

func InitGlobalTracer(tracer Tracer) {}

func StartSpan(operationName string, opts ...StartSpanOption) Span {}

func ContextWithSpan(ctx context.Context, span Span) context.Context {}

func SpanFromContext(ctx context.Context) Span {}

func StartSpanFromContext(
	ctx context.Context, operationName string, opts ...StartSpanOption
) (Span, context.Context) {}
