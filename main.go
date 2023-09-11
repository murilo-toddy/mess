package main

import (
	"log"
	"net/http"
    "github.com/murilo-toddy/api/router"
)

const port = ":8080"

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}
