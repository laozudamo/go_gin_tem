package models

import (
	"time"

	"gorm.io/gorm"
)

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
