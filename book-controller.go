package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vladigumnw/go-bookstore-mysql/pkg/models"
	"github.com/vladigumnw/go-bookstore-mysql/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    // Parse the URL parameter "bookId"
    vars := mux.Vars(r)
    bookId := vars["bookId"]

    // Log the value of bookId for debugging
    fmt.Println("Received bookId:", bookId)

    // Parse the bookId into an integer (int64) for use in the database lookup
    ID, err := strconv.ParseInt(bookId, 10, 64) // Explicitly set base to 10 and bit size to 64
    if err != nil {
        // Log the error and send a detailed message back to the client
        fmt.Println("Error parsing bookId:", err)
        http.Error(w, fmt.Sprintf("Invalid book ID: %s", bookId), http.StatusBadRequest) // Send client a bad request error
        return
    }

    // Call the model function to fetch book details by ID
    bookDetails, _ := models.GetBookById(ID)
    if err != nil {
        fmt.Println("Error fetching book details:", err)
        http.Error(w, "Book not found", http.StatusNotFound) // Handle error for missing book
        return
    }

    // Marshal book details into JSON
    res, err := json.Marshal(bookDetails)
    if err != nil {
        fmt.Println("Error marshalling book details:", err)
        http.Error(w, "Error marshalling book details", http.StatusInternalServerError)
        return
    }

    // Set the response content type to JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
		var updateBook = &models.Book{}
		utils.ParseBody(r, updateBook)
		vars := mux.Vars(r)
		bookId := vars["bookId"]
		ID, err := strconv.ParseInt(bookId, 0, 0)
		if err != nil{
			fmt.Println("error while parsing")
		}
		bookDetails, db := models.GetBookById(ID)
		if updateBook.Name != ""{
			bookDetails.Name = updateBook.Name
		}
		if updateBook.Author != ""{
			bookDetails.Author = updateBook.Author
		}
		if updateBook.Publication != "" {
			bookDetails.Publication = updateBook.Publication
		}
		db.Save(&bookDetails)
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}