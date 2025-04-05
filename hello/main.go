package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func welcomeFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"Welcome")

}

func getParams(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "ID: %v\n",id)

}


func main() {
	router := chi.NewRouter()
	
	server := http.Server{
		Addr: ":8080",
		Handler: router,
		
	}
	router.Get("/", welcomeFunc)
	router.Get("/{id}", getParams)

	err := server.ListenAndServe(); if err != nil{
		log.Fatal(err)
	}

}