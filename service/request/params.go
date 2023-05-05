package request

type AddUserParams struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
