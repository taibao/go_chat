package request

type AddUserParams struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type DelUserParams struct {
	Id uint `json:"id"`
}
