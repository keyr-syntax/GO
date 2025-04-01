package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


var findById = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, existing := vars["id"]
	if existing  {
		fmt.Fprintf(w, "User ID: %s", id)
	} else {
		fmt.Fprint(w,"Invalid ID")
	}

}

var welcome = func (w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Welcome")
}



func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/home/{id}", findById).Methods("GET")
	fmt.Println("Server is running on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil{
		log.Fatal(err)
	}

}