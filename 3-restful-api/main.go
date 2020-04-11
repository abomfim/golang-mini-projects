package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Book struct {
	ID     uuid.UUID `json:"id"`
	Isbn   string    `json:"isbn"`
	Title  string    `json:"title"`
	Author *Author   `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, http.StatusOK, books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.FromString(mux.Vars(r)["id"])

	for _, item := range books {
		if item.ID == id {
			responseJSON(w, http.StatusOK, item)
			return
		}
	}
	responseJSON(w, http.StatusNotFound, nil)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = uuid.NewV4()
	books = append(books, book)
	responseJSON(w, http.StatusCreated, book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.FromString(mux.Vars(r)["id"])

	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = id
			books = append(books, book)
			responseJSON(w, http.StatusOK, book)
			return
		}
	}
	responseJSON(w, http.StatusBadRequest, nil)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.FromString(mux.Vars(r)["id"])
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	responseJSON(w, http.StatusNoContent, books)
}

func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func main() {
	router := mux.NewRouter().
		PathPrefix("/api").
		Subrouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5555", router))
}
