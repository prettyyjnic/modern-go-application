package jaeger

import (
	"github.com/goph/emperror"
	"github.com/pkg/errors"
	"go.opencensus.io/exporter/jaeger"
)

// NewExporter creates a new, configured Jaeger exporter.
func NewExporter(config Config, serviceName string, errorHandler emperror.Handler) (*jaeger.Exporter, error) {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		Endpoint:      config.Endpoint,
		AgentEndpoint: config.AgentEndpoint,
		Username:      config.Username,
		Password:      config.Password,
		OnError:       errorHandler.Handle,
		Process: jaeger.Process{
			ServiceName: serviceName,
		},
	})

	return exporter, errors.Wrap(err, "failed to create jaeger exporter")
}
