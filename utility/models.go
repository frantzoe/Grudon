package utlt

var FileName = "data.json"

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Data struct {
	Users []User `json:"users"`
}
