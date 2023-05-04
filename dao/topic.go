package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"goGinTem/forms"
	"goGinTem/global"
	"goGinTem/models"
	"goGinTem/utils"

	"gorm.io/gorm"
)

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

		sli := []models.VoteOption{}
		for _, v := range topic.VoteOptions {
			options := models.VoteOption{
				ID:        v.ID,
				Text:      v.Text,
				VoteCount: 0,
			}
			sli = append(sli, options)
		}

		voteOptionsStr, err := json.Marshal(sli)
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

func UpdateTopicStatus(topicCheckForm *forms.ReviewTopicForm) (bool, error) {
	err := global.DB.Model(&models.Topic{}).Where("id = ?", topicCheckForm.ID).Update("status", topicCheckForm.Status).Error
	return errors.Is(err, gorm.ErrRecordNotFound), err
}

func FindTopic(id int64) (*models.Topic, error) {
	topic := &models.Topic{}
	err := global.DB.Preload("Vote").First(&topic, id).Error
	return topic, err
}

func Vote(topicId int64, optionId int64, userId int64) error {
	// Start a database transaction
	tx := global.DB.Begin()

	// Defer a function to handle the transaction commit or rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	vote := models.Vote{}
	if err := tx.First(&vote, "topic_id = ?", topicId).Error; err != nil {
		return err
	}

	options := []models.VoteOption{}
	if err := json.Unmarshal([]byte(vote.VoteOptions), &options); err != nil {
		return err
	}

	for i := range options {
		if options[i].ID == int(optionId) {
			options[i].VoteCount++
			break
		}
	}

	str, err := json.Marshal(&options)
	if err != nil {
		return err
	}

	if err := tx.Model(&vote).Where("topic_id = ?", topicId).UpdateColumn("vote_options", string(str)).Error; err != nil {
		return err
	}
	// Commit the transaction if everything succeeded
	return tx.Commit().Error
}

func GetAllTopic() ([]*models.Topic, error) {
	var topics []*models.Topic
	err := global.DB.Preload("Vote").Find(&topics).Error
	return topics, err
}
