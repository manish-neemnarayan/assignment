package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manish-neemnarayan/assignment/service"
	"github.com/manish-neemnarayan/assignment/types"
)

func AddBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var params *types.AddBook
			if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
				http.Error(w, "invalid parameters", http.StatusBadRequest)
				return
			}

			if err := isValidYear(params.PublicationYear); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			val, err := service.AddBookService(params)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			WriteJson(w, http.StatusOK, map[string]any{"data": val})

		default:
			http.Error(w, fmt.Sprintf("%s method not allowed", r.Method), http.StatusBadGateway)
			return
		}
	}
}

func DeleteBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			var bookName *types.DeleteParam
			if err := json.NewDecoder(r.Body).Decode(&bookName); err != nil {
				http.Error(w, "invalid parameters", http.StatusBadRequest)
				return
			}

			err := service.DeleteBookService(bookName.BookName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			WriteJson(w, http.StatusOK, map[string]any{"data": "successfully deleted"})

		default:
			http.Error(w, fmt.Sprintf("%s method not allowed", r.Method), http.StatusBadGateway)
			return
		}
	}
}

//utility function
//check utility.handler.go
