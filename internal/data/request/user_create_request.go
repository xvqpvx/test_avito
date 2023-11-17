package request

type UserCreateRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
