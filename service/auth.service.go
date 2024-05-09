package service

import (
	"fmt"

	"github.com/manish-neemnarayan/assignment/types"
)

type AuthService struct {
	memoryDB *MemoryDB
}

func NewAuthService(svc *MemoryDB) *AuthService {
	return &AuthService{
		memoryDB: svc,
	}
}

func (a *AuthService) Login(params *types.AuthParams) (*types.AuthReponse, error) {
	user, err := a.memoryDB.Get(params.Email)
	if err != nil {
		return nil, err
	}

	if ok := IsValidPassword(user.EncPassword, params.Password); !ok {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &types.AuthReponse{
		User: types.User{
			Name:        user.Name,
			Email:       user.Email,
			EncPassword: "",
			Role:        user.Role,
		},
		Token: CreateTokenFromUser(user),
	}, nil
}

// // utility functions
// // check the utility.service.go
