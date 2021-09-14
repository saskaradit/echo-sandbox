package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/saskaradit/echo-test/db"
	"github.com/saskaradit/echo-test/util"

	_metricHttp "github.com/saskaradit/echo-test/metric/delivery/http"
	_metricRepo "github.com/saskaradit/echo-test/metric/repository"
	_metricUseCase "github.com/saskaradit/echo-test/metric/usecase"
)

var config util.Config

func init() {
	var err error
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config")
	}
}

func main() {

	defer db.DB.Close()
	db.ConnectDB()

	e := echo.New()
	metricRepo := _metricRepo.NewMetricRepository(db.DB)

	mu := _metricUseCase.NewMetricUseCase(metricRepo)

	_metricHttp.NewMetricHandler(e, mu)
	e.Start(config.ServerAddress)
	fmt.Println("Welcome to the server!")
}
