package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	TopicID uint   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
	Content string `gorm:"not null"`
	Visible bool   `gorm:"not null;default:true"`
	Likes   uint   `gorm:"not null;default:0"`
	gorm.Model
}

type CommentReply struct {
	CommentID uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	ParentID  uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	Visible   bool   `gorm:"not null;default:true"`
	Likes     uint   `gorm:"not null;default:0"`
	gorm.Model
}
