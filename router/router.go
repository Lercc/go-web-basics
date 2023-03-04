package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Create route...")
	})

	return router
}
