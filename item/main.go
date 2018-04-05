package main

import (
	"log"
	"net/http"
	"technology/day15/item/routers"
)

func main() {
	router := routers.InitRouters()

	log.Fatal(http.ListenAndServe(":8888", router))
}
