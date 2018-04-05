package controllers

import (
	"encoding/json"
	"net/http"
	"technology/day15/officer/data"
)

func GetOfficer(w http.ResponseWriter, r *http.Request) {
	context := Context{}
	d := DBInitial(context.DB)
	defer d.Close()

	repo := &data.OfficerRepository{d}
	officer := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(officer)
	w.Write(mdl)
}
