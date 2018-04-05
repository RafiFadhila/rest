package controllers

import (
	"encoding/json"
	"net/http"
	"technology/day15/item/data"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	context := Context{}
	d := DBInitial(context.DB)
	defer d.Close()

	repo := &data.ItemRepository{d}
	item := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}
