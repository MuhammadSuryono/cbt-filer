package question

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"gtihub.com/MuhammadSuryono/cbt-uploader/file/excel"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
)

func AddExamQuestion(dataExcel excel.TemplateQuestion, groupId int64, listQuestionId int64) {
	dataExam := models.ExamQuestion{
		GroupQuestionId: groupId,
		ListQuestionId:  listQuestionId,
		Question:        dataExcel.Question,
		Option1:         dataExcel.Option1,
		Option2:         dataExcel.Option2,
		Option3:         dataExcel.Option3,
		Option4:         dataExcel.Option4,
		Option5:         dataExcel.Option5,
		Answer:          dataExcel.Answer,
	}

	_ = db.Connection.Table("exam_question").Create(&dataExam)
}
