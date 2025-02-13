package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const Port = "80"

func main() {
	app := Config{}
	log.Println("starting broker service")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: app.router(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
