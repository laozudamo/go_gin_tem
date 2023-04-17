package models

import (
	"time"

	"gorm.io/gorm"
)

type VoteTopic struct {
	TopicID   int
	UserID    int
	VotedTime time.Time
	gorm.Model
}
