package main

import (
	"log"
	"net/http"

	"github.com/adrianosela/harvest/api/config"
	"github.com/adrianosela/harvest/api/ctrl"
)

var (
	// injected at build-time
	version string
)

func main() {
	ctrl, err := ctrl.NewController(config.FromEnv(version))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":80", ctrl.Handler()); err != nil {
		log.Fatal(err)
	}
}
