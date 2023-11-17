package request

type UserUpdateRequest struct {
	IdUser    int    `json:"idUser"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
