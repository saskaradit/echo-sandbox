package domain

import "context"

// Metric
type Metric struct {
	Id             string
	Name           string
	Percentile99th string
	Percentile95th string
	Percentile90th string
	Throughput     string
}

type MetricUseCase interface {
	Fetch(ctx context.Context) ([]Metric, error)
	Create(ctx context.Context, Metric *Metric) error
}

type MetricRepository interface {
	Fetch(ctx context.Context) ([]Metric, error)
	Create(ctx context.Context, Metric *Metric) error
}
