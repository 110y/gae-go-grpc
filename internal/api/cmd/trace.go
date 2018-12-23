package cmd

import (
	"context"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

// Start registers stackdriver exporter to opencensus/trace.
func setupStackdriverTrace(ctx context.Context, project string) (func(), error) {
	opts := stackdriver.Options{
		ProjectID: project,
		Context:   ctx,
	}

	exporter, err := stackdriver.NewExporter(opts)
	if err != nil {
		return nil, err
	}

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	return func() {
		exporter.Flush()
	}, nil
}
