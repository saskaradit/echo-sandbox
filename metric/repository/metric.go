package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/saskaradit/echo-test/domain"
)

type metricRepository struct {
	db *sql.DB
}

func NewMetricRepository(db *sql.DB) domain.MetricRepository {
	return &metricRepository{
		db: db,
	}
}

func (met *metricRepository) Fetch(ctx context.Context) ([]domain.Metric, error) {
	mets := []domain.Metric{}
	rows, err := met.db.Query(`SELECT * FROM Metric`)
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {

		var met domain.Metric

		err = rows.Scan(&met.Id, &met.Name, &met.Percentile99th, &met.Percentile95th, &met.Percentile90th, &met.Throughput)
		if err != nil {
			log.Fatalln(err)
		}

		mets = append(mets, met)
	}

	return mets, nil
}

func (met *metricRepository) Create(ctx context.Context, metric *domain.Metric) error {
	insertStmt := `insert into metric("name", "percentile99th", "percentile95th", "percentile90th", "throughput") values($1, $2, $3, $4, $5)`

	_, err := met.db.Exec(insertStmt, metric.Name, metric.Percentile99th, metric.Percentile95th, metric.Percentile90th, metric.Throughput)

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
