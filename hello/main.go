package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func parseForm(db *sql.DB,w http.ResponseWriter, r *http.Request){

	r.ParseForm()
	title := r.Form.Get("title")
	content := r.Form.Get("content")
	isPublishedString := r.Form.Get("isPublished")
	isPublished  := 0

	if isPublishedString == "true"{
		isPublished = 1
	}


	query := "INSERT INTO blog(id, title, content, isPublished) VALUES(NULL, ?, ?, ?)"
	_, err := db.Exec(query,title,content,isPublished)
	if err != nil {
		log.Fatal("Error inserting data:", err)
	}
	resultQuery := "SELECT id, title, content, isPublished FROM blog"
    rows, err := db.Query(resultQuery)
    if err != nil {
        log.Fatal("Error retrieving data:", err)
    }
    defer rows.Close()

	//fmt.Fprintf(w,"Rows: %v\n", rows)

    for rows.Next() {
        var id int
        var title string
        var content string
        var isPublished int
        if err := rows.Scan(&id, &title, &content, &isPublished); err != nil {
            log.Fatal("Error scanning row:", err)
        }
        fmt.Fprintf(w, "Blog: id=%d, title=%s, content=%s, isPublished=%d\n", id, title, content, isPublished)
    }
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home")

}


func connectDB() (*sql.DB, error) {
	dsn := "root:keyr@tcp(localhost:3306)/godb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MySQL database successfully!")
	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/",home).Methods("GET")
	router.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		parseForm(db, w, r)
	} ).Methods("GET")
	
	fmt.Println("Server is running on 8080")
	if err := http.ListenAndServe(":8080", router); err !=nil{
		log.Fatal("Error connecting to server")
	} 

}