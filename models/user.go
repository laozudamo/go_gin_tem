package models

import "time"

type User struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Password string     `json:"password"`
	Username string     `json:"username"`
	HeadUrl  string     `json:"head_url"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Address  string     `json:"address"`
	Desc     string     `json:"desc"`
	Gender   string     `json:"gender"`
	Role     int        `json:"role"`
	Tel      string     `json:"tel"`
}

func (User) TableName() string {
	return "user"
}
