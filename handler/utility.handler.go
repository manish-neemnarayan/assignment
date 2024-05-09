package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/manish-neemnarayan/assignment/types"
	"golang.org/x/crypto/bcrypt"
)

func WriteJson(w http.ResponseWriter, statusCode int, msg any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}

func NewUserFromParams(params types.CreateUserParams) (*types.User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), BcryptCost)
	if err != nil {
		return nil, err
	}
	return &types.User{
		Name:        params.Name,
		Email:       params.Email,
		Role:        params.Role,
		EncPassword: string(encpw),
	}, nil
}

func isValidYear(year string) error {
	currentYear := time.Now().Year()

	intYear, err := strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("incorrect year provided")
	}

	if currentYear < intYear {
		return fmt.Errorf("incorrect year provided")
	}

	return nil
}
