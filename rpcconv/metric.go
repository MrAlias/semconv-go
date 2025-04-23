// Code generated from semantic convention specification. DO NOT EDIT.

package rpcconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// ClientDuration is an instrument used to record metric values conforming to the
// "rpc.client.duration" semantic conventions. It represents the measures the
// duration of outbound RPC.
type ClientDuration struct {
	inst metric.Int64Histogram
}

// NewClientDuration returns a new ClientDuration instrument.
func NewClientDuration(m metric.Meter) (ClientDuration, error) {
	i, err := m.Int64Histogram(
	    "rpc.client.duration",
	    metric.WithDescription("Measures the duration of outbound RPC."),
	    metric.WithUnit("ms"),
	)
	if err != nil {
	    return ClientDuration{inst: noop.Int64Histogram{}}, err
	}
	return ClientDuration{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ClientDuration) Name() string {
	return "rpc.client.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ClientDuration) Unit() string {
	return "ms"
}

// Description returns the semantic convention description of the instrument
func (ClientDuration) Description() string {
	return "Measures the duration of outbound RPC."
}

// Record records val to the current distribution.
//
// While streaming RPCs may record this metric as start-of-batch
// to end-of-batch, it's hard to interpret in practice.
//
// **Streaming**: N/A.
func (m ClientDuration) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ClientRequestSize is an instrument used to record metric values conforming to
// the "rpc.client.request.size" semantic conventions. It represents the measures
// the size of RPC request messages (uncompressed).
type ClientRequestSize struct {
	inst metric.Int64Histogram
}

// NewClientRequestSize returns a new ClientRequestSize instrument.
func NewClientRequestSize(m metric.Meter) (ClientRequestSize, error) {
	i, err := m.Int64Histogram(
	    "rpc.client.request.size",
	    metric.WithDescription("Measures the size of RPC request messages (uncompressed)."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return ClientRequestSize{inst: noop.Int64Histogram{}}, err
	}
	return ClientRequestSize{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ClientRequestSize) Name() string {
	return "rpc.client.request.size"
}

// Unit returns the semantic convention unit of the instrument
func (ClientRequestSize) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (ClientRequestSize) Description() string {
	return "Measures the size of RPC request messages (uncompressed)."
}

// Record records val to the current distribution.
//
// **Streaming**: Recorded per message in a streaming batch
func (m ClientRequestSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ClientRequestsPerRPC is an instrument used to record metric values conforming
// to the "rpc.client.requests_per_rpc" semantic conventions. It represents the
// measures the number of messages received per RPC.
type ClientRequestsPerRPC struct {
	inst metric.Int64Histogram
}

// NewClientRequestsPerRPC returns a new ClientRequestsPerRPC instrument.
func NewClientRequestsPerRPC(m metric.Meter) (ClientRequestsPerRPC, error) {
	i, err := m.Int64Histogram(
	    "rpc.client.requests_per_rpc",
	    metric.WithDescription("Measures the number of messages received per RPC."),
	    metric.WithUnit("{count}"),
	)
	if err != nil {
	    return ClientRequestsPerRPC{inst: noop.Int64Histogram{}}, err
	}
	return ClientRequestsPerRPC{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ClientRequestsPerRPC) Name() string {
	return "rpc.client.requests_per_rpc"
}

// Unit returns the semantic convention unit of the instrument
func (ClientRequestsPerRPC) Unit() string {
	return "{count}"
}

// Description returns the semantic convention description of the instrument
func (ClientRequestsPerRPC) Description() string {
	return "Measures the number of messages received per RPC."
}

// Record records val to the current distribution.
//
// Should be 1 for all non-streaming RPCs.
//
// **Streaming**: This metric is required for server and client streaming RPCs
func (m ClientRequestsPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ClientResponseSize is an instrument used to record metric values conforming to
// the "rpc.client.response.size" semantic conventions. It represents the
// measures the size of RPC response messages (uncompressed).
type ClientResponseSize struct {
	inst metric.Int64Histogram
}

// NewClientResponseSize returns a new ClientResponseSize instrument.
func NewClientResponseSize(m metric.Meter) (ClientResponseSize, error) {
	i, err := m.Int64Histogram(
	    "rpc.client.response.size",
	    metric.WithDescription("Measures the size of RPC response messages (uncompressed)."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return ClientResponseSize{inst: noop.Int64Histogram{}}, err
	}
	return ClientResponseSize{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ClientResponseSize) Name() string {
	return "rpc.client.response.size"
}

// Unit returns the semantic convention unit of the instrument
func (ClientResponseSize) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (ClientResponseSize) Description() string {
	return "Measures the size of RPC response messages (uncompressed)."
}

// Record records val to the current distribution.
//
// **Streaming**: Recorded per response in a streaming batch
func (m ClientResponseSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ClientResponsesPerRPC is an instrument used to record metric values conforming
// to the "rpc.client.responses_per_rpc" semantic conventions. It represents the
// measures the number of messages sent per RPC.
type ClientResponsesPerRPC struct {
	inst metric.Int64Histogram
}

// NewClientResponsesPerRPC returns a new ClientResponsesPerRPC instrument.
func NewClientResponsesPerRPC(m metric.Meter) (ClientResponsesPerRPC, error) {
	i, err := m.Int64Histogram(
	    "rpc.client.responses_per_rpc",
	    metric.WithDescription("Measures the number of messages sent per RPC."),
	    metric.WithUnit("{count}"),
	)
	if err != nil {
	    return ClientResponsesPerRPC{inst: noop.Int64Histogram{}}, err
	}
	return ClientResponsesPerRPC{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ClientResponsesPerRPC) Name() string {
	return "rpc.client.responses_per_rpc"
}

// Unit returns the semantic convention unit of the instrument
func (ClientResponsesPerRPC) Unit() string {
	return "{count}"
}

// Description returns the semantic convention description of the instrument
func (ClientResponsesPerRPC) Description() string {
	return "Measures the number of messages sent per RPC."
}

// Record records val to the current distribution.
//
// Should be 1 for all non-streaming RPCs.
//
// **Streaming**: This metric is required for server and client streaming RPCs
func (m ClientResponsesPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ServerDuration is an instrument used to record metric values conforming to the
// "rpc.server.duration" semantic conventions. It represents the measures the
// duration of inbound RPC.
type ServerDuration struct {
	inst metric.Int64Histogram
}

// NewServerDuration returns a new ServerDuration instrument.
func NewServerDuration(m metric.Meter) (ServerDuration, error) {
	i, err := m.Int64Histogram(
	    "rpc.server.duration",
	    metric.WithDescription("Measures the duration of inbound RPC."),
	    metric.WithUnit("ms"),
	)
	if err != nil {
	    return ServerDuration{inst: noop.Int64Histogram{}}, err
	}
	return ServerDuration{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ServerDuration) Name() string {
	return "rpc.server.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ServerDuration) Unit() string {
	return "ms"
}

// Description returns the semantic convention description of the instrument
func (ServerDuration) Description() string {
	return "Measures the duration of inbound RPC."
}

// Record records val to the current distribution.
//
// While streaming RPCs may record this metric as start-of-batch
// to end-of-batch, it's hard to interpret in practice.
//
// **Streaming**: N/A.
func (m ServerDuration) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ServerRequestSize is an instrument used to record metric values conforming to
// the "rpc.server.request.size" semantic conventions. It represents the measures
// the size of RPC request messages (uncompressed).
type ServerRequestSize struct {
	inst metric.Int64Histogram
}

// NewServerRequestSize returns a new ServerRequestSize instrument.
func NewServerRequestSize(m metric.Meter) (ServerRequestSize, error) {
	i, err := m.Int64Histogram(
	    "rpc.server.request.size",
	    metric.WithDescription("Measures the size of RPC request messages (uncompressed)."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return ServerRequestSize{inst: noop.Int64Histogram{}}, err
	}
	return ServerRequestSize{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ServerRequestSize) Name() string {
	return "rpc.server.request.size"
}

// Unit returns the semantic convention unit of the instrument
func (ServerRequestSize) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (ServerRequestSize) Description() string {
	return "Measures the size of RPC request messages (uncompressed)."
}

// Record records val to the current distribution.
//
// **Streaming**: Recorded per message in a streaming batch
func (m ServerRequestSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ServerRequestsPerRPC is an instrument used to record metric values conforming
// to the "rpc.server.requests_per_rpc" semantic conventions. It represents the
// measures the number of messages received per RPC.
type ServerRequestsPerRPC struct {
	inst metric.Int64Histogram
}

// NewServerRequestsPerRPC returns a new ServerRequestsPerRPC instrument.
func NewServerRequestsPerRPC(m metric.Meter) (ServerRequestsPerRPC, error) {
	i, err := m.Int64Histogram(
	    "rpc.server.requests_per_rpc",
	    metric.WithDescription("Measures the number of messages received per RPC."),
	    metric.WithUnit("{count}"),
	)
	if err != nil {
	    return ServerRequestsPerRPC{inst: noop.Int64Histogram{}}, err
	}
	return ServerRequestsPerRPC{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ServerRequestsPerRPC) Name() string {
	return "rpc.server.requests_per_rpc"
}

// Unit returns the semantic convention unit of the instrument
func (ServerRequestsPerRPC) Unit() string {
	return "{count}"
}

// Description returns the semantic convention description of the instrument
func (ServerRequestsPerRPC) Description() string {
	return "Measures the number of messages received per RPC."
}

// Record records val to the current distribution.
//
// Should be 1 for all non-streaming RPCs.
//
// **Streaming** : This metric is required for server and client streaming RPCs
func (m ServerRequestsPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ServerResponseSize is an instrument used to record metric values conforming to
// the "rpc.server.response.size" semantic conventions. It represents the
// measures the size of RPC response messages (uncompressed).
type ServerResponseSize struct {
	inst metric.Int64Histogram
}

// NewServerResponseSize returns a new ServerResponseSize instrument.
func NewServerResponseSize(m metric.Meter) (ServerResponseSize, error) {
	i, err := m.Int64Histogram(
	    "rpc.server.response.size",
	    metric.WithDescription("Measures the size of RPC response messages (uncompressed)."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return ServerResponseSize{inst: noop.Int64Histogram{}}, err
	}
	return ServerResponseSize{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ServerResponseSize) Name() string {
	return "rpc.server.response.size"
}

// Unit returns the semantic convention unit of the instrument
func (ServerResponseSize) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (ServerResponseSize) Description() string {
	return "Measures the size of RPC response messages (uncompressed)."
}

// Record records val to the current distribution.
//
// **Streaming**: Recorded per response in a streaming batch
func (m ServerResponseSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ServerResponsesPerRPC is an instrument used to record metric values conforming
// to the "rpc.server.responses_per_rpc" semantic conventions. It represents the
// measures the number of messages sent per RPC.
type ServerResponsesPerRPC struct {
	inst metric.Int64Histogram
}

// NewServerResponsesPerRPC returns a new ServerResponsesPerRPC instrument.
func NewServerResponsesPerRPC(m metric.Meter) (ServerResponsesPerRPC, error) {
	i, err := m.Int64Histogram(
	    "rpc.server.responses_per_rpc",
	    metric.WithDescription("Measures the number of messages sent per RPC."),
	    metric.WithUnit("{count}"),
	)
	if err != nil {
	    return ServerResponsesPerRPC{inst: noop.Int64Histogram{}}, err
	}
	return ServerResponsesPerRPC{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ServerResponsesPerRPC) Name() string {
	return "rpc.server.responses_per_rpc"
}

// Unit returns the semantic convention unit of the instrument
func (ServerResponsesPerRPC) Unit() string {
	return "{count}"
}

// Description returns the semantic convention description of the instrument
func (ServerResponsesPerRPC) Description() string {
	return "Measures the number of messages sent per RPC."
}

// Record records val to the current distribution.
//
// Should be 1 for all non-streaming RPCs.
//
// **Streaming**: This metric is required for server and client streaming RPCs
func (m ServerResponsesPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}