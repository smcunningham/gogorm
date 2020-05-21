package main

import (
	"gogorm/api"
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe(":8080", api.Route()))

}
