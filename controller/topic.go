package controller

import (
	"fmt"
	response "goGinTem/Response"
	"goGinTem/dao"
	"goGinTem/forms"
	"goGinTem/utils"

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

	_, err := dao.FindTopic(uint(checkTopicForm.ID))

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
