package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Password string `json:"password"`
	Tel      string `json:"tel"`
	UserInfo string `json:"userInfo"`
}

func (User) TableName() string {
	return "user"
}
