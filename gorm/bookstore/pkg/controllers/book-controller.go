package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go/gorm/bookstore/pkg/models"
	"github.com/go/gorm/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()

	resp, err := json.Marshal(books)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		encoder := json.NewEncoder(w)
		encoder.Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdVar := vars["bookId"]

	bookId, err := strconv.ParseInt(bookIdVar, 0, 0)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	book := models.GetBookById(bookId)
	mBook, err := json.Marshal(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mBook)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdVar := vars["bookId"]

	bookId, err := strconv.ParseInt(bookIdVar, 0, 0)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var updatingBook models.Book

	err = utils.ParseBody(r, &updatingBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	models.UpdateBookById(bookId, &updatingBook)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdVar := vars["bookId"]

	bookId, err := strconv.ParseInt(bookIdVar, 0, 0)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	models.DeleteBookById(bookId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	err := utils.ParseBody(r, &newBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		encoder := json.NewEncoder(w)
		encoder.Encode(err.Error())
		return
	}

	createdBook := newBook.CreateBook()
	mCreatedBook, err := json.Marshal(createdBook)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		encoder := json.NewEncoder(w)
		encoder.Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mCreatedBook)
}
