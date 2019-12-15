package controllers

import (
	"net/http"
)

// RegisterControllers registers controllers to handle routes
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
