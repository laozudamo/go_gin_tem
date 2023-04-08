package models

import (
	"gorm.io/gorm"
)

// 标题
type Topic struct {
	Title     string `json:"title"`
	Text      string
	Tag       string
	CreatByID uint
	Status    TopicType // 话题的状态，草稿、待审核、已发布、已结束
	gorm.Model
}

type TopicType int

// 话题的状态，包括草稿、待审核、已发布、已结束
const (
	Draft TopicType = iota
	PendingPub
	Publish
	Ending
)

func (Topic) TableName() string {
	return "topic"
}
