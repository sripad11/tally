// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tally

import (
	ubertally "github.com/uber-go/tally/v4"
)

// Scope is a namespace wrapper around a stats reporter, ensuring that
// all emitted values have a given prefix or set of tags.
//
// IMPORTANT: When using Prometheus reporters, users must take care to
//
//	not create metrics from both parent scopes and subscopes
//	that have the same metric name but different tag keys,
//	as metric allocation will panic.
type Scope = ubertally.Scope

// Counter is the interface for emitting counter type metrics.
type Counter = ubertally.Counter

// Gauge is the interface for emitting gauge metrics.
type Gauge = ubertally.Gauge

// Timer is the interface for emitting timer metrics.
type Timer = ubertally.Timer

// Histogram is the interface for emitting histogram metrics
type Histogram = ubertally.Histogram

// Stopwatch is a helper for simpler tracking of elapsed time, use the
// Stop() method to report time elapsed since its created back to the
// timer or histogram.
type Stopwatch = ubertally.Stopwatch

// NewStopwatch creates a new immutable stopwatch for recording the start
// time to a stopwatch reporter.

var NewStopwatch = ubertally.NewStopwatch

// func NewStopwatch(start time.Time, r StopwatchRecorder) Stopwatch {
// 	return ubertally.NewStopwatch{start, r}
// }

// StopwatchRecorder is a recorder that is called when a stopwatch is
// stopped with Stop().
type StopwatchRecorder = ubertally.StopwatchRecorder

// Buckets is an interface that can represent a set of buckets
// either as float64s or as durations.
type Buckets = ubertally.Buckets

// BucketPair describes the lower and upper bounds
// for a derived bucket from a buckets set.
type BucketPair = ubertally.BucketPair

// Capabilities is a description of metrics reporting capabilities.
type Capabilities = ubertally.Capabilities
