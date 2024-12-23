package http

import (
    "github.com/gorilla/mux"
)

func NewRouter(userHandler *UserHandler) *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/users", userHandler.CreateUserHandler).Methods("POST")
    r.HandleFunc("/users", userHandler.GetAllUsersHandler).Methods("GET")
    r.HandleFunc("/users/{id}", userHandler.GetUserByIDHandler).Methods("GET")

    return r
}
