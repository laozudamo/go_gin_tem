package models

import (
	"time"

	"gorm.io/gorm"
)

// 标题
type Topic struct {
	Title     string `json:"title"`
	Text      string
	Tag       string
	CreatByID uint
	Status    TopicType // 话题的状态，草稿、待审核、、已结束
	gorm.Model
}

type TopicType int

type Vote struct {
	VoteCount      int     // 投票数量
	VoteStatus     int     // 投票状态  1未开始、2进行中、3已结束
	TopicID        uint    `gorm:"foreignkey:TopicID"`
	VoteType       int     // 0单选、1多选
	VoteWay        VoteWay // 0匿名投票、1实名投票、2限定人群投票
	VoteStartTime  time.Time
	VoteEndTime    time.Time
	VoteLimit      int  // 投票限制数量 1,2,3
	IsPublicResult bool // 投票结果是否公开
	VoteOptions    string
	gorm.Model
}

type VoteWay int

const (
	AnonymousVote VoteWay = iota // 匿名投票
	RealNameVote                 // 实名投票
	LimitedVote                  // 限定人群投票
)

func (Vote) TableName() string {
	return "vote"
}

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

type VoteOption struct {
	ID        int
	Text      string
	VoteCount int
}
