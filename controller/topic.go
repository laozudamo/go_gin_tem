package controller

import (
	response "goGinTem/Response"
	"goGinTem/dao"
	"goGinTem/forms"
	"goGinTem/utils"

	"github.com/gin-gonic/gin"
)

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

	_, err := dao.CreatTopic(topic, createId)
	if err != nil {
		response.Err(c, 200, 500, "服务器内部错误", err)
	}
	response.Success(c, 200, "创建成功", "")

}

// func DelTopic(c *gin.Context) {

// }

// func UpdateTopic(c *gin.Context) {

// }
