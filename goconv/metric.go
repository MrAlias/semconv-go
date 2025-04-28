// Code generated from semantic convention specification. DO NOT EDIT.

package goconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// MemoryTypeAttr is an attribute conforming to the go.memory.type semantic
// conventions. It represents the type of memory.
type MemoryTypeAttr string

var (
	// MemoryTypeStack is the memory allocated from the heap that is reserved for
	// stack space, whether or not it is currently in-use.
	MemoryTypeStack MemoryTypeAttr = "stack"
	// MemoryTypeOther is the memory used by the Go runtime, excluding other
	// categories of memory usage described in this enumeration.
	MemoryTypeOther MemoryTypeAttr = "other"
)

// ConfigGogc is an instrument used to record metric values conforming to the
// "go.config.gogc" semantic conventions. It represents the heap size target
// percentage configured by the user, otherwise 100.
type ConfigGogc struct {
	metric.Int64ObservableUpDownCounter
}

// NewConfigGogc returns a new ConfigGogc instrument.
func NewConfigGogc(m metric.Meter) (ConfigGogc, error) {
	i, err := m.Int64ObservableUpDownCounter(
	    "go.config.gogc",
	    metric.WithDescription("Heap size target percentage configured by the user, otherwise 100."),
	    metric.WithUnit("%"),
	)
	if err != nil {
	    return ConfigGogc{noop.Int64ObservableUpDownCounter{}}, err
	}
	return ConfigGogc{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ConfigGogc) Inst() metric.Int64ObservableUpDownCounter {
	return m.Int64ObservableUpDownCounter
}

// Name returns the semantic convention name of the instrument.
func (ConfigGogc) Name() string {
	return "go.config.gogc"
}

// Unit returns the semantic convention unit of the instrument
func (ConfigGogc) Unit() string {
	return "%"
}

// Description returns the semantic convention description of the instrument
func (ConfigGogc) Description() string {
	return "Heap size target percentage configured by the user, otherwise 100."
}

// GoroutineCount is an instrument used to record metric values conforming to the
// "go.goroutine.count" semantic conventions. It represents the count of live
// goroutines.
type GoroutineCount struct {
	metric.Int64ObservableUpDownCounter
}

// NewGoroutineCount returns a new GoroutineCount instrument.
func NewGoroutineCount(m metric.Meter) (GoroutineCount, error) {
	i, err := m.Int64ObservableUpDownCounter(
	    "go.goroutine.count",
	    metric.WithDescription("Count of live goroutines."),
	    metric.WithUnit("{goroutine}"),
	)
	if err != nil {
	    return GoroutineCount{noop.Int64ObservableUpDownCounter{}}, err
	}
	return GoroutineCount{i}, nil
}

// Inst returns the underlying metric instrument.
func (m GoroutineCount) Inst() metric.Int64ObservableUpDownCounter {
	return m.Int64ObservableUpDownCounter
}

// Name returns the semantic convention name of the instrument.
func (GoroutineCount) Name() string {
	return "go.goroutine.count"
}

// Unit returns the semantic convention unit of the instrument
func (GoroutineCount) Unit() string {
	return "{goroutine}"
}

// Description returns the semantic convention description of the instrument
func (GoroutineCount) Description() string {
	return "Count of live goroutines."
}

// MemoryAllocated is an instrument used to record metric values conforming to
// the "go.memory.allocated" semantic conventions. It represents the memory
// allocated to the heap by the application.
type MemoryAllocated struct {
	metric.Int64ObservableCounter
}

// NewMemoryAllocated returns a new MemoryAllocated instrument.
func NewMemoryAllocated(m metric.Meter) (MemoryAllocated, error) {
	i, err := m.Int64ObservableCounter(
	    "go.memory.allocated",
	    metric.WithDescription("Memory allocated to the heap by the application."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return MemoryAllocated{noop.Int64ObservableCounter{}}, err
	}
	return MemoryAllocated{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MemoryAllocated) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (MemoryAllocated) Name() string {
	return "go.memory.allocated"
}

// Unit returns the semantic convention unit of the instrument
func (MemoryAllocated) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (MemoryAllocated) Description() string {
	return "Memory allocated to the heap by the application."
}

// MemoryAllocations is an instrument used to record metric values conforming to
// the "go.memory.allocations" semantic conventions. It represents the count of
// allocations to the heap by the application.
type MemoryAllocations struct {
	metric.Int64ObservableCounter
}

// NewMemoryAllocations returns a new MemoryAllocations instrument.
func NewMemoryAllocations(m metric.Meter) (MemoryAllocations, error) {
	i, err := m.Int64ObservableCounter(
	    "go.memory.allocations",
	    metric.WithDescription("Count of allocations to the heap by the application."),
	    metric.WithUnit("{allocation}"),
	)
	if err != nil {
	    return MemoryAllocations{noop.Int64ObservableCounter{}}, err
	}
	return MemoryAllocations{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MemoryAllocations) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (MemoryAllocations) Name() string {
	return "go.memory.allocations"
}

// Unit returns the semantic convention unit of the instrument
func (MemoryAllocations) Unit() string {
	return "{allocation}"
}

// Description returns the semantic convention description of the instrument
func (MemoryAllocations) Description() string {
	return "Count of allocations to the heap by the application."
}

// MemoryGCGoal is an instrument used to record metric values conforming to the
// "go.memory.gc.goal" semantic conventions. It represents the heap size target
// for the end of the GC cycle.
type MemoryGCGoal struct {
	metric.Int64ObservableUpDownCounter
}

// NewMemoryGCGoal returns a new MemoryGCGoal instrument.
func NewMemoryGCGoal(m metric.Meter) (MemoryGCGoal, error) {
	i, err := m.Int64ObservableUpDownCounter(
	    "go.memory.gc.goal",
	    metric.WithDescription("Heap size target for the end of the GC cycle."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return MemoryGCGoal{noop.Int64ObservableUpDownCounter{}}, err
	}
	return MemoryGCGoal{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MemoryGCGoal) Inst() metric.Int64ObservableUpDownCounter {
	return m.Int64ObservableUpDownCounter
}

// Name returns the semantic convention name of the instrument.
func (MemoryGCGoal) Name() string {
	return "go.memory.gc.goal"
}

// Unit returns the semantic convention unit of the instrument
func (MemoryGCGoal) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (MemoryGCGoal) Description() string {
	return "Heap size target for the end of the GC cycle."
}

// MemoryLimit is an instrument used to record metric values conforming to the
// "go.memory.limit" semantic conventions. It represents the go runtime memory
// limit configured by the user, if a limit exists.
type MemoryLimit struct {
	metric.Int64ObservableUpDownCounter
}

// NewMemoryLimit returns a new MemoryLimit instrument.
func NewMemoryLimit(m metric.Meter) (MemoryLimit, error) {
	i, err := m.Int64ObservableUpDownCounter(
	    "go.memory.limit",
	    metric.WithDescription("Go runtime memory limit configured by the user, if a limit exists."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return MemoryLimit{noop.Int64ObservableUpDownCounter{}}, err
	}
	return MemoryLimit{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MemoryLimit) Inst() metric.Int64ObservableUpDownCounter {
	return m.Int64ObservableUpDownCounter
}

// Name returns the semantic convention name of the instrument.
func (MemoryLimit) Name() string {
	return "go.memory.limit"
}

// Unit returns the semantic convention unit of the instrument
func (MemoryLimit) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (MemoryLimit) Description() string {
	return "Go runtime memory limit configured by the user, if a limit exists."
}

// MemoryUsed is an instrument used to record metric values conforming to the
// "go.memory.used" semantic conventions. It represents the memory used by the Go
// runtime.
type MemoryUsed struct {
	metric.Int64ObservableCounter
}

// NewMemoryUsed returns a new MemoryUsed instrument.
func NewMemoryUsed(m metric.Meter) (MemoryUsed, error) {
	i, err := m.Int64ObservableCounter(
	    "go.memory.used",
	    metric.WithDescription("Memory used by the Go runtime."),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return MemoryUsed{noop.Int64ObservableCounter{}}, err
	}
	return MemoryUsed{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MemoryUsed) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (MemoryUsed) Name() string {
	return "go.memory.used"
}

// Unit returns the semantic convention unit of the instrument
func (MemoryUsed) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (MemoryUsed) Description() string {
	return "Memory used by the Go runtime."
}

// AttrMemoryType returns an optional attribute for the "go.memory.type" semantic
// convention. It represents the type of memory.
func (MemoryUsed) AttrMemoryType(val MemoryTypeAttr) attribute.KeyValue {
	return attribute.String("go.memory.type", string(val))
}

// ProcessorLimit is an instrument used to record metric values conforming to the
// "go.processor.limit" semantic conventions. It represents the number of OS
// threads that can execute user-level Go code simultaneously.
type ProcessorLimit struct {
	metric.Int64ObservableUpDownCounter
}

// NewProcessorLimit returns a new ProcessorLimit instrument.
func NewProcessorLimit(m metric.Meter) (ProcessorLimit, error) {
	i, err := m.Int64ObservableUpDownCounter(
	    "go.processor.limit",
	    metric.WithDescription("The number of OS threads that can execute user-level Go code simultaneously."),
	    metric.WithUnit("{thread}"),
	)
	if err != nil {
	    return ProcessorLimit{noop.Int64ObservableUpDownCounter{}}, err
	}
	return ProcessorLimit{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ProcessorLimit) Inst() metric.Int64ObservableUpDownCounter {
	return m.Int64ObservableUpDownCounter
}

// Name returns the semantic convention name of the instrument.
func (ProcessorLimit) Name() string {
	return "go.processor.limit"
}

// Unit returns the semantic convention unit of the instrument
func (ProcessorLimit) Unit() string {
	return "{thread}"
}

// Description returns the semantic convention description of the instrument
func (ProcessorLimit) Description() string {
	return "The number of OS threads that can execute user-level Go code simultaneously."
}

// ScheduleDuration is an instrument used to record metric values conforming to
// the "go.schedule.duration" semantic conventions. It represents the time
// goroutines have spent in the scheduler in a runnable state before actually
// running.
type ScheduleDuration struct {
	metric.Float64Histogram
}

// NewScheduleDuration returns a new ScheduleDuration instrument.
func NewScheduleDuration(m metric.Meter) (ScheduleDuration, error) {
	i, err := m.Float64Histogram(
	    "go.schedule.duration",
	    metric.WithDescription("The time goroutines have spent in the scheduler in a runnable state before actually running."),
	    metric.WithUnit("s"),
	)
	if err != nil {
	    return ScheduleDuration{noop.Float64Histogram{}}, err
	}
	return ScheduleDuration{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ScheduleDuration) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (ScheduleDuration) Name() string {
	return "go.schedule.duration"
}

// Unit returns the semantic convention unit of the instrument
func (ScheduleDuration) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (ScheduleDuration) Description() string {
	return "The time goroutines have spent in the scheduler in a runnable state before actually running."
}

// Record records val to the current distribution.
//
// Computed from `/sched/latencies:seconds`. Bucket boundaries are provided by
// the runtime, and are subject to change.
func (m ScheduleDuration) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val)
	} else {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}