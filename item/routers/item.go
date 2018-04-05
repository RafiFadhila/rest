package routers

import (
	"technology/day15/item/controllers"

	"github.com/gorilla/mux"
)

func SetItemRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/item", controllers.GetItem).Methods("GET")
	return router
}
