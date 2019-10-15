package main

import (
	//"fmt"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct Details
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Details
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init Books Struct
var books []Book

// Get All Book Record
func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Particular Book Record
func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	param := mux.Vars(r)

	for _, items := range books {
		if items.ID == param["id"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}
	json.NewEncoder(w).Encode("Not Found")
}

// Add New Book Record
func createbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode("1")
}

// Update Book Record
func updatebook(w http.ResponseWriter, r *http.Request) {
	flag := 0
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	for val := range books {
		if books[val].ID == book.ID {
			if books[val].Isbn != book.Isbn {
				flag = 1
				books[val].Isbn = book.Isbn
			}

			if books[val].Title != book.Title {
				flag = 1
				books[val].Title = book.Title
			}

			if books[val].Author.Firstname != book.Author.Firstname {
				flag = 1
				books[val].Author.Firstname = book.Author.Firstname
			}

			if books[val].Author.Lastname != book.Author.Lastname {
				flag = 1
				books[val].Author.Lastname = book.Author.Lastname
			}
			break
		}
	}
	if flag == 1 {
		json.NewEncoder(w).Encode("OK")
	} else {
		json.NewEncoder(w).Encode("Not Done")
	}

}

// Delete Book Record
func deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for val, items := range books {
		if items.ID == param["id"] {
			books = append(books[:val], books[val+1:]...)
		}
	}
	json.NewEncoder(w).Encode("1")
}

func main() {

	// Append Value to books
	books = append(books, Book{ID: "1", Isbn: "871937", Title: "ATOB", Author: &Author{Firstname: "Nilesh", Lastname: "Yadav"}})
	books = append(books, Book{ID: "2", Isbn: "192384", Title: "123LOST", Author: &Author{Firstname: "Akash", Lastname: "Rai"}})
	books = append(books, Book{ID: "3", Isbn: "283894", Title: "CTOD", Author: &Author{Firstname: "Suraj", Lastname: "Tiwari"}})
	books = append(books, Book{ID: "4", Isbn: "365987", Title: "ATOS", Author: &Author{Firstname: "Ashutosh", Lastname: "Singh"}})
	books = append(books, Book{ID: "5", Isbn: "987654", Title: "STOV", Author: &Author{Firstname: "Vishal", Lastname: "Singh"}})

	//init Router
	r := mux.NewRouter()

	//Route Handlers / Endpoint
	r.HandleFunc("/api/books", getbooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getbook).Methods("GET")
	r.HandleFunc("/api/book", createbook).Methods("POST")
	r.HandleFunc("/api/booku", updatebook).Methods("PUT")
	r.HandleFunc("/api/bookd/{id}", deletebook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
