# My Go App Server

**Description**
* Assignment Task to have a in-memory mock-db store and work with csv files 

## Installation
**Dependencies**
1. `go mod download` 

## Running the Server

**Configuration**
* I didn't add any os Env for the sake of simplicity

**Starting**
1. `go mod download`
2. `go build -o ./bin/api`
3. `./bin/api`

## API Endpoints

* `/in-memory` (Methods: POST, GET)
   * This is an in-memory custom store
   * Although a Seed function runs on starting the server and it inserts two users, 
   ``` 
   { Name:user, Email:user@gmail.com, Role:user, Password: user@123}
   { Name:admin, Email:admin@gmail.com, Role:admin, Password: admin@123}
   ```
   * But you can create your own user and admins also.
   ``` 
   POST http://localhost:9005/in-memory
   Json-Body 
   {
	"name": "admin",
	"email": "admin2@gmail.com",
	"password": "admin@123",
	"role": "admin"
   }
   role ---> "user" or "admin"
   GET http://localhost:9005/in-memory?email=admin%40gmail.com

   Query-Param -- you can change it with any email, replace "@" with "%40" in query param
   ```

* `/login` (Methods: POST)
   * Description of login process and response  
   ```
   POST http://localhost:9005/login
   Json-Body 
   {
	"email": "admin@gmail.com",
	"password": "ram123"
   }

   Response
   {
	"user": {
		"name": "admin",
		"email": "admin@gmail.com",
		"role": "admin",
		"encpass": ""
	},
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGdtYWlsLmNvbSIsImV4cGlyZXMiOjE3MTUxNjg4MTQsIm5hbWUiOiJhZG1pbiIsInJvbGUiOiJhZG1pbiJ9.Eff2hJyFglMhstOz35tut6zNDpOASgozG5407DUgWJQ"
    }
   ```
* `/home`(Methods: GET)
   * Description of what happens at this endpoint
   ```
   GET http://localhost:9005/home
   Bearer Token in authorization header
   
   Response
   {
	"data": {
		"bookName": [
			"The Da Vinci Code",
			"Think and Grow Rich",
			"Harry Potter and the Half-Blood Prince",
			"The Catcher in the Rye",
			"The Alchemist",
			"Don Quixote",
			"A Tale of Two Cities",
			"The Lord of the Rings",
			"The Little Prince",
			"Harry Potter and the Sorcerer’s Stone",
			"And Then There Were None",
			"The Dream of the Red Chamber",
			"The Hobbit",
			"She: A History of Adventure",
			"The Lion, the Witch and the Wardrobe"
		]
	}
}

* `/addBook`(Methods: POST)
   * Description of adding books
   ```
   POST http://localhost:9005/addBook
   Bearer Token in authorization header
   Json-Body 
   {
	"bookName": "my book",
	"author": "manish",
	"publicationYear": "2023"
   }

   Response
   {
	"data": "success"
   }
   ```
* `/deleteBook`(Methods: DELETE)
   * Description of deleting books
   ```
   DELETE http://localhost:9005/deleteBook
   Bearer Token in authorization header
   Json-Body 
   {
	"bookName": "my book"
   }

   Response
   {
	"data": "successfully deleted"
   }
   ``` 
