package controller

import (
	"encoding/json"
	"fmt"
	response "goGinTem/Response"
	"goGinTem/dao"
	"goGinTem/forms"
	"goGinTem/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 创建话题
func CreatTopic(c *gin.Context) {

	createId := c.GetUint("userId")

	topic := &forms.TopicForm{}
	if err := c.ShouldBind(topic); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	// 时间验证
	_, ok := utils.ParseTime(topic.VoteStartTime)
	if !ok {
		response.Err(c, 200, 500, "时间格式错误", "startTime")
		return
	}
	_, b := utils.ParseTime(topic.VoteEndTime)
	if !b {
		response.Err(c, 200, 500, "时间格式错误", "endTime")
		return
	}

	_, err := dao.CreateTopicAndVote(topic, createId)
	if err != nil {
		response.Err(c, 200, 500, "服务器内部错误", err)
	}
	response.Success(c, 200, "创建成功", "")

}

func ReviewTopic(c *gin.Context) {
	checkTopicForm := &forms.ReviewTopicForm{}
	if err := c.ShouldBindJSON(checkTopicForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	fmt.Printf("checkTopicForm: %v\n", checkTopicForm)

	_, err := dao.FindTopic(int64(checkTopicForm.ID))

	if err != nil {
		response.Err(c, 200, 400, "请传入正确的id", err)
		return
	}

	isFail, err := dao.UpdateTopicStatus(checkTopicForm)
	if err != nil {
		response.Err(c, 200, 500, "服务器内部错误", err)
		return
	}
	if !isFail {
		response.Success(c, 200, "审批成功", nil)
		return
	}

}

func GetTopic(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Err(c, 200, 500, "id错误", err)
	}

	topic, err := dao.FindTopic(id)

	if err != nil {
		response.Err(c, 200, 500, "查询错误", err)
		return
	}

	data := &forms.QueryTopicForm{}
	options := []forms.Option{}

	if err := json.Unmarshal([]byte(topic.Vote.VoteOptions), &options); err != nil {
		response.Err(c, 200, 500, "解析出错", err)
		return
	}
	data.CreatByID = topic.CreatByID
	data.ID = topic.ID
	data.Status = int(topic.Status)
	data.IsPublicResult = topic.Vote.IsPublicResult
	data.Text = topic.Text
	data.Title = topic.Title
	data.VoteLimit = topic.Vote.VoteLimit
	data.VoteCount = topic.Vote.VoteCount
	data.VoteWay = int(topic.Vote.VoteWay)
	data.Tag = topic.Tag
	data.VoteStartTime = topic.Vote.VoteStartTime.Format("2006-01-02 15:04:05")
	data.VoteEndTime = topic.Vote.VoteEndTime.Format("2006-01-02 15:04:05")
	data.VoteOptions = options
	data.VoteType = topic.Vote.VoteType

	response.Success(c, 200, "获取成功", data)
}

func Vote(c *gin.Context) {
	userID := c.GetUint("userId")

	voteForm := &forms.VoteForm{}
	if err := c.ShouldBind(voteForm); err != nil {
		response.Err(c, 200, 400, "", err)
		return
	}

	topic, err := dao.FindTopic(int64(voteForm.TopicId))

	if err != nil {
		response.Err(c, 404, 500, "传入topicID错误", err)
		return
	}

	if !time.Now().After(topic.Vote.VoteStartTime) {
		response.Err(c, 200, 500, "投票时间尚未开始", nil)
		return
	}

	if !time.Now().Before(topic.Vote.VoteEndTime) {
		response.Err(c, 200, 500, "投票已经失效", nil)
		return
	}

	if err := dao.Vote(int64(voteForm.TopicId), int64(voteForm.OptionID), int64(userID)); err != nil {
		response.Err(c, 200, 500, "查询数据出错", err)
		return
	}

	response.Success(c, 200, "投票成功", nil)
}

// 获取所有话题
