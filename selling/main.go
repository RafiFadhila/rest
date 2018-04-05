package main

import (
	"log"
	"net/http"
	"technology/day15/selling/routers"
)

func main() {
	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":8886", router))
}
