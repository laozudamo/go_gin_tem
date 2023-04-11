package models

import (
	"time"

	"gorm.io/gorm"
)

// 标题
type Topic struct {
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Tag       string    `json:"tag"`
	CreatByID uint      `json:"creatorId"`
	Status    TopicType `json:"status"` // 话题的状态，草稿、待审核、、已结束
	Vote      Vote      `json:"vote"`
	gorm.Model
}

type TopicType int

type Vote struct {
	VoteCount      int       `json:"voteCount"`  // 投票数量
	VoteStatus     int       `json:"voteStatus"` // 投票状态  1未开始、2进行中、3已结束
	TopicID        uint      `json:"topicId" gorm:"foreignkey:TopicID"`
	VoteType       int       `json:"voteType"` // 0单选、1多选
	VoteWay        VoteWay   `json:"voteWay"`  // 0匿名投票、1实名投票、2限定人群投票
	VoteStartTime  time.Time `json:"voteStartTime"`
	VoteEndTime    time.Time `json:"voteEndTime"`
	VoteLimit      int       `json:"voteLimit"`      // 投票限制数量 1,2,3
	IsPublicResult bool      `json:"isPublicResult"` // 投票结果是否公开
	VoteOptions    string    `json:"voteOptions"`
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
