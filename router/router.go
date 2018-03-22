package router

import (
	"github.com/gorilla/mux"

	"demo_web_app/handler"
)

//set http router
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Host("localHost")
	r.HandleFunc("/users", handler.GetUserList).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users", handler.PostUser).Methods("POST")

	r.HandleFunc("/users/{id:[0-9]+}/relationships", handler.GetRelation).Methods("GET")
	r.HandleFunc("/users/{ida:[0-9]+}/relationships/{idb:[0-9]+}", handler.PutRelation).Methods("PUT")
	return r
}
