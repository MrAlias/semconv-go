// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated from semantic convention specification. DO NOT EDIT.

package semconv // import "github.com/MrAlias/semconv-go/semconv/v1.32.0/cpython"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// GCGenerationAttr is an attribute conforming to the cpython.gc.generation
// semantic conventions. It represents the value of the garbage collector
// collection generation.
type GCGenerationAttr int64

var (
	// GCGenerationGeneration0 is the generation 0.
	GCGenerationGeneration0 GCGenerationAttr = 0
	// GCGenerationGeneration1 is the generation 1.
	GCGenerationGeneration1 GCGenerationAttr = 1
	// GCGenerationGeneration2 is the generation 2.
	GCGenerationGeneration2 GCGenerationAttr = 2
)

// GCCollectedObjects is an instrument used to record metric values conforming to
// the "cpython.gc.collected_objects" semantic conventions. It represents the
// total number of objects collected inside a generation since interpreter start.
type GCCollectedObjects struct {
	inst metric.Int64Counter
}

// NewGCCollectedObjects returns a new GCCollectedObjects instrument.
func NewGCCollectedObjects(m metric.Meter) (GCCollectedObjects, error) {
	i, err := m.Int64Counter(
	    "cpython.gc.collected_objects",
	    metric.WithDescription("The total number of objects collected inside a generation since interpreter start."),
	    metric.WithUnit("{object}"),
	)
	if err != nil {
	    return GCCollectedObjects{}, err
	}
	return GCCollectedObjects{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (GCCollectedObjects) Name() string {
	return "cpython.gc.collected_objects"
}

// Unit returns the semantic convention unit of the instrument
func (GCCollectedObjects) Unit() string {
	return "{object}"
}

// Description returns the semantic convention description of the instrument
func (GCCollectedObjects) Description() string {
	return "The total number of objects collected inside a generation since interpreter start."
}

// Add adds incr to the existing count.
//
// The gcGeneration is the value of the garbage collector collection generation.
func (m GCCollectedObjects) Add(
	ctx context.Context,
	incr int64,
	gcGeneration GCGenerationAttr,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			append(
				attrs,
				attribute.int64("cpython.gc.generation", int64(gcGeneration)),
			)...,
		),
	)
}

// GCCollections is an instrument used to record metric values conforming to the
// "cpython.gc.collections" semantic conventions. It represents the number of
// times a generation was collected since interpreter start.
type GCCollections struct {
	inst metric.Int64Counter
}

// NewGCCollections returns a new GCCollections instrument.
func NewGCCollections(m metric.Meter) (GCCollections, error) {
	i, err := m.Int64Counter(
	    "cpython.gc.collections",
	    metric.WithDescription("The number of times a generation was collected since interpreter start."),
	    metric.WithUnit("{collection}"),
	)
	if err != nil {
	    return GCCollections{}, err
	}
	return GCCollections{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (GCCollections) Name() string {
	return "cpython.gc.collections"
}

// Unit returns the semantic convention unit of the instrument
func (GCCollections) Unit() string {
	return "{collection}"
}

// Description returns the semantic convention description of the instrument
func (GCCollections) Description() string {
	return "The number of times a generation was collected since interpreter start."
}

// Add adds incr to the existing count.
//
// The gcGeneration is the value of the garbage collector collection generation.
func (m GCCollections) Add(
	ctx context.Context,
	incr int64,
	gcGeneration GCGenerationAttr,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			append(
				attrs,
				attribute.int64("cpython.gc.generation", int64(gcGeneration)),
			)...,
		),
	)
}

// GCUncollectableObjects is an instrument used to record metric values
// conforming to the "cpython.gc.uncollectable_objects" semantic conventions. It
// represents the total number of objects which were found to be uncollectable
// inside a generation since interpreter start.
type GCUncollectableObjects struct {
	inst metric.Int64Counter
}

// NewGCUncollectableObjects returns a new GCUncollectableObjects instrument.
func NewGCUncollectableObjects(m metric.Meter) (GCUncollectableObjects, error) {
	i, err := m.Int64Counter(
	    "cpython.gc.uncollectable_objects",
	    metric.WithDescription("The total number of objects which were found to be uncollectable inside a generation since interpreter start."),
	    metric.WithUnit("{object}"),
	)
	if err != nil {
	    return GCUncollectableObjects{}, err
	}
	return GCUncollectableObjects{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (GCUncollectableObjects) Name() string {
	return "cpython.gc.uncollectable_objects"
}

// Unit returns the semantic convention unit of the instrument
func (GCUncollectableObjects) Unit() string {
	return "{object}"
}

// Description returns the semantic convention description of the instrument
func (GCUncollectableObjects) Description() string {
	return "The total number of objects which were found to be uncollectable inside a generation since interpreter start."
}

// Add adds incr to the existing count.
//
// The gcGeneration is the value of the garbage collector collection generation.
func (m GCUncollectableObjects) Add(
	ctx context.Context,
	incr int64,
	gcGeneration GCGenerationAttr,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			append(
				attrs,
				attribute.int64("cpython.gc.generation", int64(gcGeneration)),
			)...,
		),
	)
}