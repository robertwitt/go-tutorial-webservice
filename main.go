package main

import (
	"net/http"

	"github.com/robertwitt/go-tutorial-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
