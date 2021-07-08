package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = "8091"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

