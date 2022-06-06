package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/leehaowei/go-dbapp/src/models"
	"github.com/leehaowei/go-dbapp/src/utils"
	"net/http"
)

type ResponseResult struct {
	Result string `json:"result"`
}

func setSameHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "ListBooks",
	})
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "GetBook",
	})
}

//func CreateBook(w http.ResponseWriter, r *http.Request) {
//	setSameHeader(w)
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(ResponseResult{
//		Result: "CreateBook",
//	})
//}

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
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "DeleteBook",
	})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "UpdateBook",
	})
}
