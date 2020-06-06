package main

import (
	"log"
	"net/http"
)

var (
	// injected at build-time
	version string
)

func main() {
	ctrl, err := NewController(config(version))
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":80", ctrl.HTTP()); err != nil {
		log.Fatal(err)
	}
}
