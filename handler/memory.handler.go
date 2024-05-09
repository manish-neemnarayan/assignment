package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manish-neemnarayan/assignment/service"
	"github.com/manish-neemnarayan/assignment/types"
)

const BcryptCost = 12

func MemoryDBHandler(svc *service.MemoryDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		switch r.Method {
		// to insert data if method is post
		case http.MethodPost:
			var params *types.CreateUserParams
			if err := json.NewDecoder(body).Decode(&params); err != nil {
				http.Error(w, "Invalid params", http.StatusBadRequest)
				return
			}

			userMemoryParams, err := NewUserFromParams(*params)
			if err != nil {
				http.Error(w, "Invalid params", http.StatusBadRequest)
			}

			response, err := svc.Post(userMemoryParams)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			WriteJson(w, http.StatusOK, response)

		// to get data if method is get
		case http.MethodGet:
			key := r.URL.Query()["email"]
			if len(key) == 0 {
				http.Error(w, "Query key not found inside query params", http.StatusBadRequest)
			}

			response, err := svc.Get(key[0])
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}

			WriteJson(w, http.StatusOK, response)

		// bad http method
		default:
			http.Error(w, fmt.Sprintf("%s method is not allowed", r.Method), http.StatusBadGateway)
		}
	}
}

// /// utility functions --------------------------------
// /// check utility.handler.go
