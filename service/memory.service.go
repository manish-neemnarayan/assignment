package service

import (
	"fmt"

	"github.com/manish-neemnarayan/assignment/types"
	"golang.org/x/crypto/bcrypt"
)

type MemoryDBer interface {
	Post(*types.User) (*types.User, error)
	Get(string) (*types.User, error)
}

type MemoryDB struct {
	store map[string]*types.User
}

func NewMemoryDB() *MemoryDB {
	fmt.Println("In-memory DB is started")
	return &MemoryDB{
		store: make(map[string]*types.User),
	}
}

func (m *MemoryDB) Post(data *types.User) (*types.User, error) {
	m.store[data.Email] = data
	return &types.User{
		Name:        data.Name,
		Email:       data.Email,
		Role:        data.Role,
		EncPassword: data.EncPassword,
	}, nil
}

func (m *MemoryDB) Get(key string) (*types.User, error) {
	data, ok := m.store[key]
	if !ok {
		return &types.User{}, fmt.Errorf("user not found")
	}

	return data, nil
}

func (m *MemoryDB) Seed() error {

	userpass, _ := bcrypt.GenerateFromPassword([]byte("user@123"), 12)
	adminpass, _ := bcrypt.GenerateFromPassword([]byte("admin@123"), 12)
	users := []*types.User{
		{
			Name:        "user",
			Email:       "user@gmail.com",
			Role:        "user",
			EncPassword: string(userpass),
		},
		{
			Name:        "admin",
			Email:       "admin@gmail.com",
			Role:        "admin",
			EncPassword: string(adminpass),
		},
	}

	for _, user := range users {
		m.store[user.Email] = user
		fmt.Printf("%+v\n", *user)
	}
	fmt.Printf("seed data for user and admin is added\n")
	return nil
}
