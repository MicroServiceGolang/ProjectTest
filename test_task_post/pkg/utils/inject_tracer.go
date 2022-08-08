package utils 

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

func InjectTracer(ctx context.Context, event, comment string) (context.Context, error) {
	span, spanContext := opentracing.StartSpanFromContext(ctx, event)
	defer span.Finish()

	span.LogFields(
		log.String("event", event),
		log.String("comment", comment),
	)

	carrier := opentracing.TextMapCarrier{}

	ext.SpanKindRPCClient.Set(span)
	err := span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		carrier,
	)
	if err != nil {
		return nil, err
	}
	return spanContext, nil
}
