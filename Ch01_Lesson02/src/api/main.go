package main

import (
	"github.com/dNaszta/ms_go/Ch01_Lesson02/src/api/handlers"
	"github.com/dNaszta/ms_go/Ch01_Lesson02/src/api/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewRepository(
		"localhost:27017",
		"packt",
		"timeZones",
		"root",
		"example",
	)
	defer repo.Close()

	h := handlers.Handlers{
		Repo: repo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/timeZones", h.All).Methods("GET")
	r.HandleFunc("/timeZones/{timeZone}", h.GetByTZ).Methods("GET")

	r.HandleFunc("/timeZones", h.Insert).Methods("POST")
	r.HandleFunc("/timeZones/{timeZone}", h.Delete).Methods("DELETE")
	r.HandleFunc("/timeZones/{timeZone}", h.Update).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8080", r))
}