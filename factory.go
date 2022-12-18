package tailtracer

// import (
// 	"time"
// 	"context"

// 	"go.opentelemetry.io/collector/component"
// 	"go.opentelemetry.io/collector/config"
// 	"go.opentelemetry.io/collector/consumer"
// )

// const (
// 	typeStr = "tailtracer"
// 	defaultInterval = "1m"
// )

// func createDefaultConfig() component.Config {
// 	return &Config{
// 		ReceiverSettings:   config.NewReceiverSettings(component.NewID(typeStr)),
// 		TimeInterval: interval.Duration(1 * time.Second),
// 	}
// }

// func createTracesReceiver(_ context.Context, params component.ReceiverCreateSettings, baseCfg config.Receiver, consumer consumer.Traces) (component.TracesReceiver, error) {
// 	if consumer == nil {
// 		return nil, component.ErrNilNextConsumer
// 	}

// 	logger := params.Logger
// 	tailtracerCfg := baseCfg.(*Config)

// 	traceRcvr := &tailtracerReceiver{
// 		logger:       logger,
// 		nextConsumer: consumer,
// 		config:       tailtracerCfg,
// 	}

// 	return traceRcvr, nil

// }

// // NewFactory creates a factory for tailtracer receiver.
// func NewFactory() component.ReceiverFactory {
// 	return component.NewReceiverFactory(
// 		typeStr,
// 		createDefaultConfig,
// 		component.WithTracesReceiver(createTracesReceiver))
// }

import (
	"context"
    "strconv"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	typeStr = "tailtracer"
	defaultInterval = 1
)

func createDefaultConfig() component.Config {
	return &Config{
		ReceiverSettings:   config.NewReceiverSettings(component.NewID(typeStr)),	
		Interval:           strconv.Itoa(defaultInterval),
	}
}

func createTracesReceiver(_ context.Context, params receiver.CreateSettings , cfg component.Config, nextConsumer consumer.Traces) (receiver.Traces, error) {
  return nil,nil
}

// NewFactory creates a factory for tailtracer receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, component.StabilityLevelStable))
}