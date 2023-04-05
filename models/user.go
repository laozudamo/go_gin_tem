package models

import "time"

type UserInfo struct {
	Address string `json:"address"`
	// HeadUrl  string     `json:"head_url"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Username string     `json:"username"`
	Desc     string     `json:"desc"`
	Gender   string     `json:"gender"`
	Role     int        `json:"role"`
	Tel      string     `json:"tel"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Password string `json:"password"`
	Tel      string `json:"tel"`
	UserInfo string `json:"userInfo"`
}

func (User) TableName() string {
	return "user"
}
