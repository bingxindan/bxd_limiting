package main

import (
	"sync"
)

type DefaultMetricCollector struct {
	mutex *sync.RWMutex

	numRequests *Number
	errors *Number

	successes *Number
	failures *Number
	rejects *Number
	shortCircuits *Number
	timeouts *Number
	contextCanceled *Number
	contextDeadlineExceeded *Number
	fallbackSuccesses *Number
	fallbackFailures *Number
	totalDuration *Timing
	runDuration *Timing
}

func (d *DefaultMetricCollector) Update(r MetricResult) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	d.numRequests.Increment(r.Attempts)
	d.errors.Increment(r.Errors)
	d.successes.Increment(r.Successes)
	d.failures.Increment(r.Failures)
	d.rejects.Increment(r.Rejects)
	d.shortCircuits.Increment(r.ShortCircuits)
	d.timeouts.Increment(r.Timeouts)
	d.fallbackSuccesses.Increment(r.FallbackSuccesses)
	d.fallbackFailures.Increment(r.FallbackFailures)
	d.contextCanceled.Increment(r.ContextCanceled)
	d.contextDeadlineExceeded.Increment(r.ContextDeadlineExceeded)

	d.totalDuration.Add(r.TotalDuration)
	d.runDuration.Add(r.RunDuration)
}

// Reset resets all metrics in this collector to 0.
func (d *DefaultMetricCollector) Reset() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.numRequests = NewNumber()
	d.errors = NewNumber()
	d.successes = NewNumber()
	d.rejects = NewNumber()
	d.shortCircuits = NewNumber()
	d.failures = NewNumber()
	d.timeouts = NewNumber()
	d.fallbackSuccesses = NewNumber()
	d.fallbackFailures = NewNumber()
	d.contextCanceled = NewNumber()
	d.contextDeadlineExceeded = NewNumber()
	d.totalDuration = NewTiming()
	d.runDuration = NewTiming()
}

// NumRequests returns the rolling number of requests
func (d *DefaultMetricCollector) NumRequests() *Number {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.numRequests
}

// Errors returns the rolling number of errors
func (d *DefaultMetricCollector) Errors() *Number {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.errors
}
