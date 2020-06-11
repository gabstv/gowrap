package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/gabstv/gowrap tool
// using ../templates/timeout template

//go:generate gowrap gen -p github.com/gabstv/gowrap/templates_tests -i TestInterface -t ../templates/timeout -o interface_with_timeout.go

import (
	"context"
	"time"
)

// TestInterfaceWithTimeout implements TestInterface interface instrumented with timeouts
type TestInterfaceWithTimeout struct {
	TestInterface
	config TestInterfaceWithTimeoutConfig
}

type TestInterfaceWithTimeoutConfig struct {
	FTimeout time.Duration
}

// NewTestInterfaceWithTimeout returns TestInterfaceWithTimeout
func NewTestInterfaceWithTimeout(base TestInterface, config TestInterfaceWithTimeoutConfig) TestInterfaceWithTimeout {
	return TestInterfaceWithTimeout{
		TestInterface: base,
		config:        config,
	}
}

// F implements TestInterface
func (_d TestInterfaceWithTimeout) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	var cancelFunc func()
	if _d.config.FTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.FTimeout)
		defer cancelFunc()
	}
	return _d.TestInterface.F(ctx, a1, a2...)
}
