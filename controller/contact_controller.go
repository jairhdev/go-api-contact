package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jairhdev/go-api-contact/model/entities/entity_contact"
)

// ****** FUNÇÕES, MÉTODOS UTILS
// ******

func readBody(w http.ResponseWriter, r *http.Request) []byte {
	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	return result
}

func readParamId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}

// ****** SAVE
// ******
func save(w http.ResponseWriter, r *http.Request) {
	contact, err := entity_contact.NewContact(readBody(w, r))
	if err != nil {
		e := newStandardError(http.StatusBadRequest, "Invalid message body.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	// Injeta dependência
	var service = entity_contact.NewService(contact)

	contact.Id, err = service.Save()
	if err != nil {
		e := newStandardError(http.StatusInternalServerError, "Error in database.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, contact.Id))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

// ****** FIND ALL
// ******
func findAll(w http.ResponseWriter, r *http.Request) {
	// Injeta dependência
	var service = entity_contact.NewService(entity_contact.NewContactEmpty())

	result, err := service.FindAll()
	if err != nil {
		e := newStandardError(http.StatusInternalServerError, "Error in database.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// ****** FIND BY ID
// ******
func findById(w http.ResponseWriter, r *http.Request) {
	id, err := readParamId(r)
	if err != nil {
		e := newStandardError(http.StatusBadRequest, "Invalid parameters.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	// Injeta dependência
	var service = entity_contact.NewService(entity_contact.NewContactEmpty())

	result, err := service.FindById(id)
	// testar primeiro 'error in database'
	if err != nil {
		e := newStandardError(http.StatusInternalServerError, "Error in database.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}
	if result == entity_contact.NewContactEmpty() {
		e := newStandardError(http.StatusNotFound, fmt.Sprintf("Not found. Id: %d", id), err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// ****** UPDATE BY ID
// ******
func updateById(w http.ResponseWriter, r *http.Request) {
	id, err := readParamId(r)
	if err != nil {
		e := newStandardError(http.StatusBadRequest, "Invalid parameters.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	contact, err := entity_contact.NewContact(readBody(w, r))
	if err != nil {
		e := newStandardError(http.StatusBadRequest, "Invalid message body.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	// Injeta dependência
	var service = entity_contact.NewService(contact)

	rows, err := service.UpdateById(id)
	// testar primeiro 'error in database'
	if err != nil {
		e := newStandardError(http.StatusInternalServerError, "Error in database.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	if rows < 1 {
		e := newStandardError(http.StatusNotFound, fmt.Sprintf("Not found. Id: %d", id), err, r.RequestURI)
		responseError(w, r, e)
		return
	}
	contact.Id = id // atualiza para response JSON

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}

// ****** DELETE BY ID
// ******
func deleteById(w http.ResponseWriter, r *http.Request) {
	id, err := readParamId(r)
	if err != nil {
		e := newStandardError(http.StatusBadRequest, "Invalid parameters.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	// Injeta dependência
	var service = entity_contact.NewService(entity_contact.NewContactEmpty())

	rows, err := service.DeleteById(id)
	// testar primeiro 'error in database'
	if err != nil {
		e := newStandardError(http.StatusInternalServerError, "Error in database.", err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	if rows < 1 {
		e := newStandardError(http.StatusNotFound, fmt.Sprintf("Not found. Id: %d", id), err, r.RequestURI)
		responseError(w, r, e)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ****** RESPONSE ERROR
// ******
func responseError(w http.ResponseWriter, r *http.Request, err standardError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}
