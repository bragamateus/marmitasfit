package main

import ( "net/http"
	"github.com/gorilla/mux" )

func createHandler() (handler *mux.Router) {

	// creats router
	handler = mux.NewRouter()

	// associate register user route
	handler.HandleFunc(USER_PATH, RegisterUser).Methods(http.MethodPost)
	handler.HandleFunc(USER_PATH, SearchUserEmail).Methods(http.MethodGet)
	handler.HandleFunc(USER_PATH, UpdateUser).Methods(http.MethodPut)

	// returns handle
	return
}

