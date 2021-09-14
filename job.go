package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/saskaradit/echo-test/util"
)

// type Job struct {
// 	Id       string
// 	AppName  string // job name and the .jmx file
// 	SlaveNum string // the number of slaves
// 	CSVCount string // the number of csv files
// 	CSVFile  string // the name of the csv splitted by ';' -> for example (sku;user)
// }

func RunTest(c echo.Context) error {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config")
	}

	out, err := exec.Command("/bin/bash", "jenkinsjob.sh", config.JenkinsUser, config.JenkinsToken).Output()
	if err != nil {
		log.Fatalln(err)
	}

	return c.String(http.StatusOK, string(out))
}
