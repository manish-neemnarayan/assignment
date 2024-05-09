package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/manish-neemnarayan/assignment/service"
	"github.com/manish-neemnarayan/assignment/types"
)

func AuthHandler(svc *service.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// to insert data if method is post
		case http.MethodPost:
			// log.Fatal(r.Header.Get("role"))
			var (
				params *types.AuthParams
				body   io.ReadCloser
			)
			body = r.Body
			// fmt.Println(r.Header.Get("role"))
			if err := json.NewDecoder(body).Decode(&params); err != nil {
				http.Error(w, "Invalid params", http.StatusBadRequest)
				return
			}

			response, err := svc.Login(params)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			WriteJson(w, http.StatusOK, response)

		// bad http method
		default:
			http.Error(w, fmt.Sprintf("%s method is not allowed", r.Method), http.StatusBadGateway)
		}
	}
}
