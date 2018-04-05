package controllers

import (
	"encoding/json"
	"net/http"
	"technology/day15/selling/data"
)

func GetSelling(w http.ResponseWriter, r *http.Request) {
	context := Context{}
	d := DBInitial(context.DB)
	defer d.Close()

	repo := &data.SellingRepository{d}
	selling := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(selling)
	w.Write(mdl)
}
