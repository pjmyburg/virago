package main

import (
	"flag"
	"github.com/pjmyburg/virago/config"
	"github.com/pjmyburg/virago/server"
	"github.com/pjmyburg/virago/sqs"
	"github.com/pjmyburg/virago/sqs/api"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "enable debug logging")
	flag.Parse()

	log.SetOutput(os.Stdout)

	if debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Starting up in DEBUG mode")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	conf, err := config.LoadYaml("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	sqsInstance := sqs.New(conf)
	sqsAPI := api.NewAPI(sqsInstance)

	s := server.New(sqsAPI)
	if err := http.ListenAndServe("0.0.0.0:1234", s); err != nil {
		log.Fatal(err)
	}
}
