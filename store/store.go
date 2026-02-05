package store


type LoginStruct struct {
	HashedPass   string
	SessionToken string
	CSRFToken    string
}

// Fake in-memory DB
var Users = map[string]LoginStruct{}