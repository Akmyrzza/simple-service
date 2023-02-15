package entities

type UserSt struct {
	Id    int `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserCUSt struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}
