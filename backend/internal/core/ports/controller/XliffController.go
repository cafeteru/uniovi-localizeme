package controller

import (
	"net/http"
)

type XliffController interface {
	Read(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}
