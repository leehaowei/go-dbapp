package controllers

import (
	"encoding/json"
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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "CreateBook",
	})
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
