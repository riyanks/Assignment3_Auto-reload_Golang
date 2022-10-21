package main

import (
	"log"
	"net/http"
	"time"

	"github.com/husfuu/go-auto-reload/handler"
	"github.com/husfuu/go-auto-reload/helper"
	"github.com/husfuu/go-auto-reload/service"
)

func main() {
	go autoUpdate()
	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal()
	}
}

func autoUpdate() {
	disaster := helper.GetDisaster()
	for range time.Tick(15 * time.Second) {
		service.UpdateJSON(disaster)
	}
}
