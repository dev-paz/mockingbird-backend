package dto

//User struct does a thing
type User struct {
	ID    string
	Name  string
	Email string
}

type SignInRequest struct {
	TokenID string
}
