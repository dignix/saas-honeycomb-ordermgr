package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
