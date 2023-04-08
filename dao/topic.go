package dao

import (
	"encoding/json"
	"fmt"
	"goGinTem/forms"
	"goGinTem/global"
	"goGinTem/models"
	"goGinTem/utils"

	"gorm.io/gorm"
)

func CreatTopic(topic *forms.TopicForm, creatId uint) (bool, error) {
	newTopic := models.Topic{}
	newTopic.Title = topic.Title
	newTopic.Text = topic.Text
	newTopic.Tag = topic.Tag
	newTopic.CreatByID = creatId
	newTopic.Status = 0
	if err := global.DB.Create(&newTopic).Error; err != nil {
		panic(err)
	}
	isOK, err := CreateTopicAndVote(topic, newTopic.ID)
	return isOK, err
}

func CreateTopicAndVote(topic *forms.TopicForm, creatId uint) (bool, error) {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		newTopic := models.Topic{
			Title:     topic.Title,
			Text:      topic.Text,
			Tag:       topic.Tag,
			CreatByID: creatId,
			Status:    0,
		}
		if err := tx.Create(&newTopic).Error; err != nil {
			return err
		}

		vote := models.Vote{
			TopicID:        newTopic.ID,
			VoteStatus:     0,
			VoteCount:      0,
			VoteType:       topic.VoteType,
			VoteWay:        0,
			VoteLimit:      topic.VoteLimit,
			IsPublicResult: topic.IsPublicResult,
		}

		startTime, ok := utils.ParseTime(topic.VoteStartTime)
		if !ok {
			return fmt.Errorf("invalid vote start time")
		}
		endTime, ok := utils.ParseTime(topic.VoteEndTime)
		if !ok {
			return fmt.Errorf("invalid vote end time")
		}

		voteOptionsStr, err := json.Marshal(topic.VoteOptions)
		if err != nil {
			return err
		}

		vote.VoteStartTime = *startTime
		vote.VoteOptions = string(voteOptionsStr)
		vote.VoteEndTime = *endTime

		if err := tx.Create(&vote).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
