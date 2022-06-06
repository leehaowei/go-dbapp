package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leehaowei/go-dbapp/src/models"
	"github.com/leehaowei/go-dbapp/src/utils"
	"net/http"
	"strconv"
)

type ResponseResult struct {
	Result string `json:"result"`
}

func setSameHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		fmt.Println(err)
	}
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		book, _ := models.GetBook(id)
		res, err := json.Marshal(book)
		if err != nil {
			fmt.Println(err)
		}
		setSameHeader(w)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseRequestBody(r, book)
	fmt.Println(book)
	newBook := book.CreateBook()
	fmt.Println(newBook)
	res, _ := json.Marshal(newBook)
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		book := models.DeleteBook(id)
		res, _ := json.Marshal(book)
		setSameHeader(w)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseRequestBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		book, db := models.GetBook(id)
		if updateBook.Title != "" {
			book.Title = updateBook.Title
		}
		if updateBook.Price > -1 {
			book.Price = updateBook.Price
		}
		if updateBook.Author != "" {
			book.Author = updateBook.Author
		}
		fmt.Println(book)

		db.Save(&book)
		res, _ := json.Marshal(book)
		setSameHeader(w)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
