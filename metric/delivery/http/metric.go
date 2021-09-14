package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/saskaradit/echo-test/domain"
)

type MetricHandler struct {
	MUsecase domain.MetricUseCase
}

func NewMetricHandler(e *echo.Echo, ms domain.MetricUseCase) {
	handler := &MetricHandler{
		MUsecase: ms,
	}

	g := e.Group("/v1/api")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	g.GET("/metrics", handler.FetchMetrics)
	g.POST("/metrics", handler.Create)
}

func (m *MetricHandler) FetchMetrics(c echo.Context) error {
	ctx := c.Request().Context()

	metrics, err := m.MUsecase.Fetch(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "There is something wrong with fetch")
	}
	return c.JSON(http.StatusOK, metrics)
}

func (m *MetricHandler) Create(c echo.Context) error {
	metric := domain.Metric{}

	defer c.Request().Body.Close()

	err := c.Bind(&metric)
	if err != nil {
		log.Fatalln("Failed to read ", err)
		return c.String(http.StatusInternalServerError, "")
	}

	ctx := c.Request().Context()
	err = m.MUsecase.Create(ctx, &metric)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error while saving")
	}

	return c.JSON(http.StatusCreated, metric)
}
