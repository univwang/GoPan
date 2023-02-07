package models

type UserBasic struct {
	Id       int    `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
