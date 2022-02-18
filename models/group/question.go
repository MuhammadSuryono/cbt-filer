package group

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
)

func CreateGroupQuestion(listQuestionId int64, groupName, description, fileGroup, questionGroup string) models.GroupQuestion {
	already, group := alreadyGroupName(listQuestionId, groupName)
	if already {
		return group
	}

	dataGroup := models.GroupQuestion{
		ListQuestionId: listQuestionId,
		Name:           groupName,
		Description:    description,
		FilenameGroup:  fileGroup,
		QuestionGroup:  questionGroup,
	}
	_ = db.Connection.Table("group_question").Create(&dataGroup)
	return dataGroup
}

func alreadyGroupName(listQuestionId int64, groupName string) (already bool, groupQuestion models.GroupQuestion) {
	_ = db.Connection.Table("group_question").Where("name = ?", groupName).Where("list_question_id = ?", listQuestionId).First(&groupQuestion)
	if groupQuestion.Id != 0 {
		already = true
	}
	return
}
