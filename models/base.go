package models

type BaseModel struct {
	Id int64 `json:"id"`
}

type Models interface {
	User
}
