package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartContactRoutes() {
	r := mux.NewRouter().PathPrefix("/api/contacts/v1").Subrouter()

	r.HandleFunc("", save).Methods(http.MethodPost)
	r.HandleFunc("", findAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", findById).Methods(http.MethodGet)
	r.HandleFunc("/{id}", updateById).Methods(http.MethodPut)
	r.HandleFunc("/{id}", deleteById).Methods(http.MethodDelete)

	http.Handle("/", r)
}
