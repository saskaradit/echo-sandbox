package usecase

import (
	"context"

	"github.com/saskaradit/echo-test/domain"
)

type metricUseCase struct {
	metricRepo domain.MetricRepository
}

func NewMetricUseCase(metrepo domain.MetricRepository) domain.MetricUseCase {
	return &metricUseCase{
		metricRepo: metrepo,
	}
}

func (c *metricUseCase) Fetch(ctx context.Context) ([]domain.Metric, error) {
	metrics, err := c.metricRepo.Fetch(ctx)
	return metrics, err
}

func (c *metricUseCase) Create(ctx context.Context, met *domain.Metric) error {
	err := c.metricRepo.Create(ctx, met)
	return err
}
