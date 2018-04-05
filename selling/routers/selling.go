package routers

import (
	"technology/day15/selling/controllers"

	"github.com/gorilla/mux"
)

func SetSellingRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/selling", controllers.GetSelling).Methods("GET")
	return router
}
