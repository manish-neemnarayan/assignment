package handler

import (
	"fmt"
	"net/http"

	"github.com/manish-neemnarayan/assignment/service"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			role := r.Header.Get("role")

			val, err := service.HomeService(role)
			if err != nil {
				fmt.Printf("error while attaching home service to home handler %s", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			WriteJson(w, http.StatusOK, map[string]any{"data": val})

		default:
			http.Error(w, fmt.Sprintf("%s method not allowed", r.Method), http.StatusBadGateway)
		}
	}
}
