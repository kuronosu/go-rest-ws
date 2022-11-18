package models

type Models[M any] interface {
	GetModel() M
}

func (u User) GetModel() User { return u }

// func (p Post) GetModel() Post { return p }
