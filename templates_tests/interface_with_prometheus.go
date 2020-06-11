package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/gabstv/gowrap tool
// using ../templates/prometheus template

//go:generate gowrap gen -p github.com/gabstv/gowrap/templates_tests -i TestInterface -t ../templates/prometheus -o interface_with_prometheus.go

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TestInterfaceWithPrometheus implements TestInterface interface with all methods wrapped
// with Prometheus metrics
type TestInterfaceWithPrometheus struct {
	base         TestInterface
	instanceName string
}

var testinterfaceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "testinterface_duration_seconds",
		Help:       "testinterface runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewTestInterfaceWithPrometheus returns an instance of the TestInterface decorated with prometheus summary metric
func NewTestInterfaceWithPrometheus(base TestInterface, instanceName string) TestInterfaceWithPrometheus {
	return TestInterfaceWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// F implements TestInterface
func (_d TestInterfaceWithPrometheus) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		testinterfaceDurationSummaryVec.WithLabelValues(_d.instanceName, "F", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.F(ctx, a1, a2...)
}

// NoError implements TestInterface
func (_d TestInterfaceWithPrometheus) NoError(s1 string) (s2 string) {
	_since := time.Now()
	defer func() {
		result := "ok"
		testinterfaceDurationSummaryVec.WithLabelValues(_d.instanceName, "NoError", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.NoError(s1)
}

// NoParamsOrResults implements TestInterface
func (_d TestInterfaceWithPrometheus) NoParamsOrResults() {
	_since := time.Now()
	defer func() {
		result := "ok"
		testinterfaceDurationSummaryVec.WithLabelValues(_d.instanceName, "NoParamsOrResults", result).Observe(time.Since(_since).Seconds())
	}()
	_d.base.NoParamsOrResults()
	return
}
