package router

import (
	"github.com/Sambit99/Go-Basics/MongoAPI/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PATCH")
	r.HandleFunc("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")

	return r
}
