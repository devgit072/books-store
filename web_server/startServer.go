package web_server

import (
	"encoding/json"
	"github.com/devgit072/books-store/db"
	"github.com/devgit072/books-store/models"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func StartServer() error {
	router := mux.NewRouter()

	router.HandleFunc("/ping", ping).Methods("GET")

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		glog.Infof("Error while starting server: %s", err)
		return err
	}
	return nil
}

func ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	glog.Info("Getting list of all books..")
	books, err := db.GetBooks()
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infof("Books: %+v", books)
	json.NewEncoder(w).Encode(&books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	glog.Info("Getting a book with id:")
	params := mux.Vars(r) // Get params in form of map.
	id,err := strconv.Atoi(params["id"])
	if err != nil {
		glog.Fatal(err)
	}
	book , err := db.GetBook(id)
	if err != nil {
		glog.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	glog.Info("Adding book into DB")
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		glog.Fatal(err.Error())
	}
	glog.Infof("Book: %+v", book)
	// now insert this book into DB.
	id, err := db.AddBook(&book)
	if err != nil {
		glog.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(id)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	glog.Info("Updating book")
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		glog.Fatal(err)
	}

	rowsAffected, err := db.UpdateBook(&book)
	if err != nil {
		glog.Fatal(err)
	}
	json.NewEncoder(w).Encode(rowsAffected)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	glog.Info("removing book..")
	params := mux.Vars(r)
	id,err := strconv.Atoi(params["id"])
	if err != nil {
		glog.Fatal(err.Error())
	}
	rowsAffected, err := db.RemoveBook(id)
	if err != nil {
		glog.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(rowsAffected)
}
