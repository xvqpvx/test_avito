package response

type UserResponse struct {
	IdUser    int    `json:"idUser"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
