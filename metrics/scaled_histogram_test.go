package metrics_test

import (
	"testing"

	"github.com/peterbourgon/gokit/metrics"
	"github.com/peterbourgon/gokit/metrics/expvar"
)

func TestScaledHistogram(t *testing.T) {
	quantiles := []int{50, 90, 99}
	scale := int64(10)
	metricName := "test_scaled_histogram"

	var h metrics.Histogram
	h = expvar.NewHistogram(metricName, 0, 1000, 3, quantiles...)
	h = metrics.NewScaledHistogram(h, scale)

	const seed, mean, stdev = 333, 500, 100          // input values
	populateNormalHistogram(t, h, seed, mean, stdev) // will be scaled down
	assertExpvarNormalHistogram(t, metricName, mean/scale, stdev/scale, quantiles)
}
