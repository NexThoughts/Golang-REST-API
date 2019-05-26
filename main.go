package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lasttname string `json:"lastname"`
}

var books [] Book

func main() {
	fmt.Println("Hello World")
	//Init Router
	r := mux.NewRouter()

	//TODO: Implements Database
	books = append(books, Book{"1", "448730", "Book1", &Author{"Vijay", "Shukla"}})
	books = append(books, Book{"2", "4454546", "Book2", &Author{"Vendetta", "Vendetta"}})

	// Route Handlers //Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

//Delete Books
func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

}

//Update Books
func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

}

//Create Books
func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)

	json.NewEncoder(writer).Encode(book)
}

//Get Books with Id
func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Book{})
}

//Get All Books
func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}
