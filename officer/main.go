package main

import (
	"log"
	"net/http"
	"technology/day15/officer/routers"
)

func main() {
	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":8887", router))
}
