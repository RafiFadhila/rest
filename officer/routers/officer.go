package routers

import (
	"technology/day15/officer/controllers"

	"github.com/gorilla/mux"
)

func SetOfficerRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/officer", controllers.GetOfficer).Methods("GET")
	return router
}
