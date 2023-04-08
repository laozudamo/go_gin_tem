package forms

type TopicForm struct {
	Title          string       `json:"title"`
	Text           string       `json:"text"`
	Tag            string       `json:"tag"`
	VoteType       int          `json:"voteType"` // 0单选、1多选
	VoteWay        int          `json:"voteWay"`  // 0匿名投票、1实名投票、2限定人群投票
	VoteStartTime  string       `json:"startTime"`
	VoteEndTime    string       `json:"endTime"`
	VoteLimit      int          `json:"voteLimit"` // 投票限制数量 1,2,3
	IsPublicResult bool         `json:"isPublic"`  // 投票结果是否公开
	VoteOptions    []VoteOption `json:"options"`
}

type VoteOption struct {
	ID   int
	Text string
}
