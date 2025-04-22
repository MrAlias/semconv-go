// Code generated from semantic convention specification. DO NOT EDIT.

package semconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ErrorTypeAttr is an attribute conforming to the error.type semantic
// conventions. It represents the describes a class of error the operation ended
// with.
type ErrorTypeAttr string

var (
	// ErrorTypeOther is a fallback error value to be used when the instrumentation
	// doesn't define a custom value.
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)

// ComponentTypeAttr is an attribute conforming to the otel.component.type
// semantic conventions. It represents a name identifying the type of the
// OpenTelemetry component.
type ComponentTypeAttr string

var (
	// ComponentTypeBatchingSpanProcessor is the builtin SDK Batching Span
	// Processor.
	ComponentTypeBatchingSpanProcessor ComponentTypeAttr = "batching_span_processor"
	// ComponentTypeSimpleSpanProcessor is the builtin SDK Simple Span Processor.
	ComponentTypeSimpleSpanProcessor ComponentTypeAttr = "simple_span_processor"
	// ComponentTypeBatchingLogProcessor is the builtin SDK Batching LogRecord
	// Processor.
	ComponentTypeBatchingLogProcessor ComponentTypeAttr = "batching_log_processor"
	// ComponentTypeSimpleLogProcessor is the builtin SDK Simple LogRecord
	// Processor.
	ComponentTypeSimpleLogProcessor ComponentTypeAttr = "simple_log_processor"
	// ComponentTypeOtlpGRPCSpanExporter is the OTLP span exporter over gRPC with
	// protobuf serialization.
	ComponentTypeOtlpGRPCSpanExporter ComponentTypeAttr = "otlp_grpc_span_exporter"
	// ComponentTypeOtlpHTTPSpanExporter is the OTLP span exporter over HTTP with
	// protobuf serialization.
	ComponentTypeOtlpHTTPSpanExporter ComponentTypeAttr = "otlp_http_span_exporter"
	// ComponentTypeOtlpHTTPJSONSpanExporter is the OTLP span exporter over HTTP
	// with JSON serialization.
	ComponentTypeOtlpHTTPJSONSpanExporter ComponentTypeAttr = "otlp_http_json_span_exporter"
	// ComponentTypeOtlpGRPCLogExporter is the OTLP LogRecord exporter over gRPC
	// with protobuf serialization.
	ComponentTypeOtlpGRPCLogExporter ComponentTypeAttr = "otlp_grpc_log_exporter"
	// ComponentTypeOtlpHTTPLogExporter is the OTLP LogRecord exporter over HTTP
	// with protobuf serialization.
	ComponentTypeOtlpHTTPLogExporter ComponentTypeAttr = "otlp_http_log_exporter"
	// ComponentTypeOtlpHTTPJSONLogExporter is the OTLP LogRecord exporter over HTTP
	// with JSON serialization.
	ComponentTypeOtlpHTTPJSONLogExporter ComponentTypeAttr = "otlp_http_json_log_exporter"
)

// SpanSamplingResultAttr is an attribute conforming to the
// otel.span.sampling_result semantic conventions. It represents the result value
// of the sampler for this span.
type SpanSamplingResultAttr string

var (
	// SpanSamplingResultDrop is the span is not sampled and not recording.
	SpanSamplingResultDrop SpanSamplingResultAttr = "DROP"
	// SpanSamplingResultRecordOnly is the span is not sampled, but recording.
	SpanSamplingResultRecordOnly SpanSamplingResultAttr = "RECORD_ONLY"
	// SpanSamplingResultRecordAndSample is the span is sampled and recording.
	SpanSamplingResultRecordAndSample SpanSamplingResultAttr = "RECORD_AND_SAMPLE"
)

// SDKExporterLogExported is an instrument used to record metric values
// conforming to the "otel.sdk.exporter.log.exported" semantic conventions. It
// represents the number of log records for which the export has finished, either
// successful or failed.
type SDKExporterLogExported struct {
	inst metric.Int64Counter
}

// NewSDKExporterLogExported returns a new SDKExporterLogExported instrument.
func NewSDKExporterLogExported(m metric.Meter) (SDKExporterLogExported, error) {
	i, err := m.Int64Counter(
	    "otel.sdk.exporter.log.exported",
	    metric.WithDescription("The number of log records for which the export has finished, either successful or failed"),
	    metric.WithUnit("{log_record}"),
	)
	if err != nil {
	    return SDKExporterLogExported{}, err
	}
	return SDKExporterLogExported{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKExporterLogExported) Name() string {
	return "otel.sdk.exporter.log.exported"
}

// Unit returns the semantic convention unit of the instrument
func (SDKExporterLogExported) Unit() string {
	return "{log_record}"
}

// Description returns the semantic convention description of the instrument
func (SDKExporterLogExported) Description() string {
	return "The number of log records for which the export has finished, either successful or failed"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKExporterLogExported) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (SDKExporterLogExported) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKExporterLogExported) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKExporterLogExported) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the server domain name if available without
// reverse DNS lookup; otherwise, IP address or Unix domain socket name.
func (SDKExporterLogExported) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the server port number.
func (SDKExporterLogExported) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// SDKExporterLogInflight is an instrument used to record metric values
// conforming to the "otel.sdk.exporter.log.inflight" semantic conventions. It
// represents the number of log records which were passed to the exporter, but
// that have not been exported yet (neither successful, nor failed).
type SDKExporterLogInflight struct {
	inst metric.Int64UpDownCounter
}

// NewSDKExporterLogInflight returns a new SDKExporterLogInflight instrument.
func NewSDKExporterLogInflight(m metric.Meter) (SDKExporterLogInflight, error) {
	i, err := m.Int64UpDownCounter(
	    "otel.sdk.exporter.log.inflight",
	    metric.WithDescription("The number of log records which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)"),
	    metric.WithUnit("{log_record}"),
	)
	if err != nil {
	    return SDKExporterLogInflight{}, err
	}
	return SDKExporterLogInflight{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKExporterLogInflight) Name() string {
	return "otel.sdk.exporter.log.inflight"
}

// Unit returns the semantic convention unit of the instrument
func (SDKExporterLogInflight) Unit() string {
	return "{log_record}"
}

// Description returns the semantic convention description of the instrument
func (SDKExporterLogInflight) Description() string {
	return "The number of log records which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKExporterLogInflight) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKExporterLogInflight) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKExporterLogInflight) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the server domain name if available without
// reverse DNS lookup; otherwise, IP address or Unix domain socket name.
func (SDKExporterLogInflight) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the server port number.
func (SDKExporterLogInflight) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// SDKExporterSpanExportedCount is an instrument used to record metric values
// conforming to the "otel.sdk.exporter.span.exported.count" semantic
// conventions. It represents the number of spans for which the export has
// finished, either successful or failed.
type SDKExporterSpanExportedCount struct {
	inst metric.Int64Counter
}

// NewSDKExporterSpanExportedCount returns a new SDKExporterSpanExportedCount
// instrument.
func NewSDKExporterSpanExportedCount(m metric.Meter) (SDKExporterSpanExportedCount, error) {
	i, err := m.Int64Counter(
	    "otel.sdk.exporter.span.exported.count",
	    metric.WithDescription("The number of spans for which the export has finished, either successful or failed"),
	    metric.WithUnit("{span}"),
	)
	if err != nil {
	    return SDKExporterSpanExportedCount{}, err
	}
	return SDKExporterSpanExportedCount{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKExporterSpanExportedCount) Name() string {
	return "otel.sdk.exporter.span.exported.count"
}

// Unit returns the semantic convention unit of the instrument
func (SDKExporterSpanExportedCount) Unit() string {
	return "{span}"
}

// Description returns the semantic convention description of the instrument
func (SDKExporterSpanExportedCount) Description() string {
	return "The number of spans for which the export has finished, either successful or failed"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKExporterSpanExportedCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents the describes a class of error the operation ended
// with.
func (SDKExporterSpanExportedCount) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKExporterSpanExportedCount) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKExporterSpanExportedCount) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the server domain name if available without
// reverse DNS lookup; otherwise, IP address or Unix domain socket name.
func (SDKExporterSpanExportedCount) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the server port number.
func (SDKExporterSpanExportedCount) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// SDKExporterSpanInflightCount is an instrument used to record metric values
// conforming to the "otel.sdk.exporter.span.inflight.count" semantic
// conventions. It represents the number of spans which were passed to the
// exporter, but that have not been exported yet (neither successful, nor
// failed).
type SDKExporterSpanInflightCount struct {
	inst metric.Int64UpDownCounter
}

// NewSDKExporterSpanInflightCount returns a new SDKExporterSpanInflightCount
// instrument.
func NewSDKExporterSpanInflightCount(m metric.Meter) (SDKExporterSpanInflightCount, error) {
	i, err := m.Int64UpDownCounter(
	    "otel.sdk.exporter.span.inflight.count",
	    metric.WithDescription("The number of spans which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)"),
	    metric.WithUnit("{span}"),
	)
	if err != nil {
	    return SDKExporterSpanInflightCount{}, err
	}
	return SDKExporterSpanInflightCount{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKExporterSpanInflightCount) Name() string {
	return "otel.sdk.exporter.span.inflight.count"
}

// Unit returns the semantic convention unit of the instrument
func (SDKExporterSpanInflightCount) Unit() string {
	return "{span}"
}

// Description returns the semantic convention description of the instrument
func (SDKExporterSpanInflightCount) Description() string {
	return "The number of spans which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKExporterSpanInflightCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKExporterSpanInflightCount) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKExporterSpanInflightCount) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// AttrServerAddress returns an optional attribute for the "server.address"
// semantic convention. It represents the server domain name if available without
// reverse DNS lookup; otherwise, IP address or Unix domain socket name.
func (SDKExporterSpanInflightCount) AttrServerAddress(val string) attribute.KeyValue {
	return attribute.String("server.address", val)
}

// AttrServerPort returns an optional attribute for the "server.port" semantic
// convention. It represents the server port number.
func (SDKExporterSpanInflightCount) AttrServerPort(val int) attribute.KeyValue {
	return attribute.Int("server.port", val)
}

// SDKLogCreated is an instrument used to record metric values conforming to the
// "otel.sdk.log.created" semantic conventions. It represents the number of logs
// submitted to enabled SDK Loggers.
type SDKLogCreated struct {
	inst metric.Int64Counter
}

// NewSDKLogCreated returns a new SDKLogCreated instrument.
func NewSDKLogCreated(m metric.Meter) (SDKLogCreated, error) {
	i, err := m.Int64Counter(
	    "otel.sdk.log.created",
	    metric.WithDescription("The number of logs submitted to enabled SDK Loggers"),
	    metric.WithUnit("{log_record}"),
	)
	if err != nil {
	    return SDKLogCreated{}, err
	}
	return SDKLogCreated{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKLogCreated) Name() string {
	return "otel.sdk.log.created"
}

// Unit returns the semantic convention unit of the instrument
func (SDKLogCreated) Unit() string {
	return "{log_record}"
}

// Description returns the semantic convention description of the instrument
func (SDKLogCreated) Description() string {
	return "The number of logs submitted to enabled SDK Loggers"
}

func (m SDKLogCreated) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// SDKProcessorLogProcessed is an instrument used to record metric values
// conforming to the "otel.sdk.processor.log.processed" semantic conventions. It
// represents the number of log records for which the processing has finished,
// either successful or failed.
type SDKProcessorLogProcessed struct {
	inst metric.Int64Counter
}

// NewSDKProcessorLogProcessed returns a new SDKProcessorLogProcessed instrument.
func NewSDKProcessorLogProcessed(m metric.Meter) (SDKProcessorLogProcessed, error) {
	i, err := m.Int64Counter(
	    "otel.sdk.processor.log.processed",
	    metric.WithDescription("The number of log records for which the processing has finished, either successful or failed"),
	    metric.WithUnit("{log_record}"),
	)
	if err != nil {
	    return SDKProcessorLogProcessed{}, err
	}
	return SDKProcessorLogProcessed{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKProcessorLogProcessed) Name() string {
	return "otel.sdk.processor.log.processed"
}

// Unit returns the semantic convention unit of the instrument
func (SDKProcessorLogProcessed) Unit() string {
	return "{log_record}"
}

// Description returns the semantic convention description of the instrument
func (SDKProcessorLogProcessed) Description() string {
	return "The number of log records for which the processing has finished, either successful or failed"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKProcessorLogProcessed) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents a low-cardinality description of the failure reason.
// SDK Batching Log Record Processors MUST use `queue_full` for log records
// dropped due to a full queue.
func (SDKProcessorLogProcessed) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKProcessorLogProcessed) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKProcessorLogProcessed) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// SDKProcessorLogQueueCapacity is an instrument used to record metric values
// conforming to the "otel.sdk.processor.log.queue.capacity" semantic
// conventions. It represents the maximum number of log records the queue of a
// given instance of an SDK Log Record processor can hold.
type SDKProcessorLogQueueCapacity struct {
	inst metric.Int64UpDownCounter
}

// NewSDKProcessorLogQueueCapacity returns a new SDKProcessorLogQueueCapacity
// instrument.
func NewSDKProcessorLogQueueCapacity(m metric.Meter) (SDKProcessorLogQueueCapacity, error) {
	i, err := m.Int64UpDownCounter(
	    "otel.sdk.processor.log.queue.capacity",
	    metric.WithDescription("The maximum number of log records the queue of a given instance of an SDK Log Record processor can hold"),
	    metric.WithUnit("{log_record}"),
	)
	if err != nil {
	    return SDKProcessorLogQueueCapacity{}, err
	}
	return SDKProcessorLogQueueCapacity{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKProcessorLogQueueCapacity) Name() string {
	return "otel.sdk.processor.log.queue.capacity"
}

// Unit returns the semantic convention unit of the instrument
func (SDKProcessorLogQueueCapacity) Unit() string {
	return "{log_record}"
}

// Description returns the semantic convention description of the instrument
func (SDKProcessorLogQueueCapacity) Description() string {
	return "The maximum number of log records the queue of a given instance of an SDK Log Record processor can hold"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKProcessorLogQueueCapacity) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKProcessorLogQueueCapacity) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKProcessorLogQueueCapacity) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// SDKProcessorLogQueueSize is an instrument used to record metric values
// conforming to the "otel.sdk.processor.log.queue.size" semantic conventions. It
// represents the number of log records in the queue of a given instance of an
// SDK log processor.
type SDKProcessorLogQueueSize struct {
	inst metric.Int64UpDownCounter
}

// NewSDKProcessorLogQueueSize returns a new SDKProcessorLogQueueSize instrument.
func NewSDKProcessorLogQueueSize(m metric.Meter) (SDKProcessorLogQueueSize, error) {
	i, err := m.Int64UpDownCounter(
	    "otel.sdk.processor.log.queue.size",
	    metric.WithDescription("The number of log records in the queue of a given instance of an SDK log processor"),
	    metric.WithUnit("{log_record}"),
	)
	if err != nil {
	    return SDKProcessorLogQueueSize{}, err
	}
	return SDKProcessorLogQueueSize{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKProcessorLogQueueSize) Name() string {
	return "otel.sdk.processor.log.queue.size"
}

// Unit returns the semantic convention unit of the instrument
func (SDKProcessorLogQueueSize) Unit() string {
	return "{log_record}"
}

// Description returns the semantic convention description of the instrument
func (SDKProcessorLogQueueSize) Description() string {
	return "The number of log records in the queue of a given instance of an SDK log processor"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKProcessorLogQueueSize) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKProcessorLogQueueSize) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKProcessorLogQueueSize) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// SDKProcessorSpanProcessedCount is an instrument used to record metric values
// conforming to the "otel.sdk.processor.span.processed.count" semantic
// conventions. It represents the number of spans for which the processing has
// finished, either successful or failed.
type SDKProcessorSpanProcessedCount struct {
	inst metric.Int64Counter
}

// NewSDKProcessorSpanProcessedCount returns a new SDKProcessorSpanProcessedCount
// instrument.
func NewSDKProcessorSpanProcessedCount(m metric.Meter) (SDKProcessorSpanProcessedCount, error) {
	i, err := m.Int64Counter(
	    "otel.sdk.processor.span.processed.count",
	    metric.WithDescription("The number of spans for which the processing has finished, either successful or failed"),
	    metric.WithUnit("{span}"),
	)
	if err != nil {
	    return SDKProcessorSpanProcessedCount{}, err
	}
	return SDKProcessorSpanProcessedCount{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKProcessorSpanProcessedCount) Name() string {
	return "otel.sdk.processor.span.processed.count"
}

// Unit returns the semantic convention unit of the instrument
func (SDKProcessorSpanProcessedCount) Unit() string {
	return "{span}"
}

// Description returns the semantic convention description of the instrument
func (SDKProcessorSpanProcessedCount) Description() string {
	return "The number of spans for which the processing has finished, either successful or failed"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKProcessorSpanProcessedCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrErrorType returns an optional attribute for the "error.type" semantic
// convention. It represents a low-cardinality description of the failure reason.
// SDK Batching Span Processors MUST use `queue_full` for spans dropped due to a
// full queue.
func (SDKProcessorSpanProcessedCount) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	return attribute.String("error.type", string(val))
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKProcessorSpanProcessedCount) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKProcessorSpanProcessedCount) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// SDKProcessorSpanQueueCapacity is an instrument used to record metric values
// conforming to the "otel.sdk.processor.span.queue.capacity" semantic
// conventions. It represents the maximum number of spans the queue of a given
// instance of an SDK span processor can hold.
type SDKProcessorSpanQueueCapacity struct {
	inst metric.Int64UpDownCounter
}

// NewSDKProcessorSpanQueueCapacity returns a new SDKProcessorSpanQueueCapacity
// instrument.
func NewSDKProcessorSpanQueueCapacity(m metric.Meter) (SDKProcessorSpanQueueCapacity, error) {
	i, err := m.Int64UpDownCounter(
	    "otel.sdk.processor.span.queue.capacity",
	    metric.WithDescription("The maximum number of spans the queue of a given instance of an SDK span processor can hold"),
	    metric.WithUnit("{span}"),
	)
	if err != nil {
	    return SDKProcessorSpanQueueCapacity{}, err
	}
	return SDKProcessorSpanQueueCapacity{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKProcessorSpanQueueCapacity) Name() string {
	return "otel.sdk.processor.span.queue.capacity"
}

// Unit returns the semantic convention unit of the instrument
func (SDKProcessorSpanQueueCapacity) Unit() string {
	return "{span}"
}

// Description returns the semantic convention description of the instrument
func (SDKProcessorSpanQueueCapacity) Description() string {
	return "The maximum number of spans the queue of a given instance of an SDK span processor can hold"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKProcessorSpanQueueCapacity) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKProcessorSpanQueueCapacity) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKProcessorSpanQueueCapacity) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// SDKProcessorSpanQueueSize is an instrument used to record metric values
// conforming to the "otel.sdk.processor.span.queue.size" semantic conventions.
// It represents the number of spans in the queue of a given instance of an SDK
// span processor.
type SDKProcessorSpanQueueSize struct {
	inst metric.Int64UpDownCounter
}

// NewSDKProcessorSpanQueueSize returns a new SDKProcessorSpanQueueSize
// instrument.
func NewSDKProcessorSpanQueueSize(m metric.Meter) (SDKProcessorSpanQueueSize, error) {
	i, err := m.Int64UpDownCounter(
	    "otel.sdk.processor.span.queue.size",
	    metric.WithDescription("The number of spans in the queue of a given instance of an SDK span processor"),
	    metric.WithUnit("{span}"),
	)
	if err != nil {
	    return SDKProcessorSpanQueueSize{}, err
	}
	return SDKProcessorSpanQueueSize{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKProcessorSpanQueueSize) Name() string {
	return "otel.sdk.processor.span.queue.size"
}

// Unit returns the semantic convention unit of the instrument
func (SDKProcessorSpanQueueSize) Unit() string {
	return "{span}"
}

// Description returns the semantic convention description of the instrument
func (SDKProcessorSpanQueueSize) Description() string {
	return "The number of spans in the queue of a given instance of an SDK span processor"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKProcessorSpanQueueSize) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrComponentName returns an optional attribute for the "otel.component.name"
// semantic convention. It represents a name uniquely identifying the instance of
// the OpenTelemetry component within its containing SDK instance.
func (SDKProcessorSpanQueueSize) AttrComponentName(val string) attribute.KeyValue {
	return attribute.String("otel.component.name", val)
}

// AttrComponentType returns an optional attribute for the "otel.component.type"
// semantic convention. It represents a name identifying the type of the
// OpenTelemetry component.
func (SDKProcessorSpanQueueSize) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	return attribute.String("otel.component.type", string(val))
}

// SDKSpanEndedCount is an instrument used to record metric values conforming to
// the "otel.sdk.span.ended.count" semantic conventions. It represents the number
// of created spans for which the end operation was called.
type SDKSpanEndedCount struct {
	inst metric.Int64Counter
}

// NewSDKSpanEndedCount returns a new SDKSpanEndedCount instrument.
func NewSDKSpanEndedCount(m metric.Meter) (SDKSpanEndedCount, error) {
	i, err := m.Int64Counter(
	    "otel.sdk.span.ended.count",
	    metric.WithDescription("The number of created spans for which the end operation was called"),
	    metric.WithUnit("{span}"),
	)
	if err != nil {
	    return SDKSpanEndedCount{}, err
	}
	return SDKSpanEndedCount{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKSpanEndedCount) Name() string {
	return "otel.sdk.span.ended.count"
}

// Unit returns the semantic convention unit of the instrument
func (SDKSpanEndedCount) Unit() string {
	return "{span}"
}

// Description returns the semantic convention description of the instrument
func (SDKSpanEndedCount) Description() string {
	return "The number of created spans for which the end operation was called"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKSpanEndedCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrSpanSamplingResult returns an optional attribute for the
// "otel.span.sampling_result" semantic convention. It represents the result
// value of the sampler for this span.
func (SDKSpanEndedCount) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	return attribute.String("otel.span.sampling_result", string(val))
}

// SDKSpanLiveCount is an instrument used to record metric values conforming to
// the "otel.sdk.span.live.count" semantic conventions. It represents the number
// of created spans for which the end operation has not been called yet.
type SDKSpanLiveCount struct {
	inst metric.Int64UpDownCounter
}

// NewSDKSpanLiveCount returns a new SDKSpanLiveCount instrument.
func NewSDKSpanLiveCount(m metric.Meter) (SDKSpanLiveCount, error) {
	i, err := m.Int64UpDownCounter(
	    "otel.sdk.span.live.count",
	    metric.WithDescription("The number of created spans for which the end operation has not been called yet"),
	    metric.WithUnit("{span}"),
	)
	if err != nil {
	    return SDKSpanLiveCount{}, err
	}
	return SDKSpanLiveCount{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (SDKSpanLiveCount) Name() string {
	return "otel.sdk.span.live.count"
}

// Unit returns the semantic convention unit of the instrument
func (SDKSpanLiveCount) Unit() string {
	return "{span}"
}

// Description returns the semantic convention description of the instrument
func (SDKSpanLiveCount) Description() string {
	return "The number of created spans for which the end operation has not been called yet"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m SDKSpanLiveCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrSpanSamplingResult returns an optional attribute for the
// "otel.span.sampling_result" semantic convention. It represents the result
// value of the sampler for this span.
func (SDKSpanLiveCount) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	return attribute.String("otel.span.sampling_result", string(val))
}