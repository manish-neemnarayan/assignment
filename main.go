package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/manish-neemnarayan/assignment/handler"
	mw "github.com/manish-neemnarayan/assignment/middleware"
	"github.com/manish-neemnarayan/assignment/service"
)

// user struct with roles: user & admin
func main() {
	var (
		svc     = service.NewMemoryDB() //it creates a temp in-memory db store
		authSVC = service.NewAuthService(svc)
	)

	if err := svc.Seed(); err != nil {
		fmt.Println("Not able to seed data")
	}

	http.HandleFunc("/in-memory", handler.MemoryDBHandler(svc)) //this route accepts two method only: POST & GET
	http.HandleFunc("/login", handler.AuthHandler(authSVC))
	http.Handle("/home", mw.AuthenticateToken(http.HandlerFunc(handler.HomeHandler())))
	http.Handle("/addBook", mw.AuthenticateToken(mw.IsAdmin(http.HandlerFunc(handler.AddBookHandler()))))
	http.Handle("/deleteBook", mw.AuthenticateToken(mw.IsAdmin(http.HandlerFunc(handler.DeleteBookHandler()))))

	fmt.Println("server is running on port http://localhost:9005")
	log.Fatal(http.ListenAndServe(":9005", nil))
}
