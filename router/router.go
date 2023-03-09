package router

import (
	"crud/handlers"

	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index)

	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", handlers.IndexUsers).Methods("GET")
	userRouter.HandleFunc("/create", handlers.CreateUser).Methods("GET")
	userRouter.HandleFunc("/store", handlers.StoreUser).Methods("POST")
	userRouter.HandleFunc("/delete", handlers.DeleteUser).Methods("GET")
	userRouter.HandleFunc("/edit", handlers.EditUser).Methods("GET")
	userRouter.HandleFunc("/update", handlers.UpdateUser).Methods("POST")

	return router
}
