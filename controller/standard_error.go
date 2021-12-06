package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jairhdev/go-api-contact/util"
)

type standardError struct {
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	Err       string    `json:"err"`
	Path      string    `json:"path"`
	TimeStamp time.Time `json:"timeStamp"`
}

func newStandardError(status int, message string, err error, path string) standardError {
	// Code defensivo. Previne "invalid memory address or nil pointer dereference"
	var e string
	if status == http.StatusNotFound && err == nil {
		e = "Not found."
	} else {
		e = err.Error()
	}

	se := standardError{
		Status:    status,
		Message:   message,
		Err:       e,
		Path:      path,
		TimeStamp: getTimeStamp(),
	}

	l, err := json.Marshal(se) // prepara log
	if err != nil {
		util.NewLog(fmt.Sprintf("%v | %v", se, err))
	} else {
		util.NewLog(string(l))
	}
	return se
}

func getTimeStamp() time.Time {
	const tf string = "2006-01-02T15:04:05Z"

	t, err := time.Parse(time.RFC3339, time.Now().UTC().Format(tf))
	if err != nil {
		util.NewLog(err.Error())
	}
	return t
}
