package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Books Attributes
type Books struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
	BookID    string `json:"bookid"`
	Quantity  string `json:"quantity"`
}

// Book list
var Book []Books

func homepage(w http.ResponseWriter, r *http.Request) {

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Book)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, value := range Book {
		// Check for Specific Id
		if value.BookID == params["bookid"] {
			json.NewEncoder(w).Encode(value)
		}
	}
	json.NewEncoder(w).Encode("")
}

func main() {

	// Append Books
	Book = append(Book, Books{"JavaScript", "Pike", "Oxford", "The Untitled Story", "105", "200"}, Books{"Java", "Ken", "Abc", "Object Oriented Programming", "106", "200"}, Books{"Golang", "Robert", "HPH", "Evolution In Programming", "107", "200"})

	// init Router
	r := mux.NewRouter()

	// Handle URL
	r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/getBooks", getBooks).Methods("GET")
	r.HandleFunc("/getBook/{id}", getBook).Methods("GET")

	// Start Server On Port 8080
	log.Fatal(http.ListenAndServe(":8094", r))

}
