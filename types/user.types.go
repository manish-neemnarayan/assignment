package types

const Secret = "sfsjlks8798slkdjflk" //can get from os.getenv()
type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthReponse struct {
	User  `json:"user"`
	Token string `json:"token"`
}

type HomeResponse struct {
	BookName []string `json:"bookName"`
}

type DeleteParam struct {
	BookName string `json:"bookName"`
}

type AddBook struct {
	BookName        string `json:"bookName"`
	Author          string `json:"author"`
	PublicationYear string `json:"publicationYear"`
}
type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	EncPassword string `json:"encpass"`
}
