package user

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/exam"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/question"
	type_exam "gtihub.com/MuhammadSuryono/cbt-uploader/models/type"
	"gtihub.com/MuhammadSuryono/cbt-uploader/pagination"
	"time"
)

type UserParticipant struct {
	Id               int64     `json:"id"`
	SessionId        int64     `json:"session_id"`
	FullName         string    `json:"full_name"`
	PlaceOfBirth     string    `json:"place_of_birth"`
	DateOfBirth      string    `json:"date_of_birth"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	Address          string    `json:"address"`
	DateRegister     string    `json:"date_register"`
	NumberOfRegister string    `json:"number_of_register"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type UserParticipantResponse struct {
	Id               int64     `json:"id"`
	SessionId        int64     `json:"session_id"`
	FullName         string    `json:"full_name"`
	PlaceOfBirth     string    `json:"place_of_birth"`
	DateOfBirth      string    `json:"date_of_birth"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	Address          string    `json:"address"`
	DateRegister     string    `json:"date_register"`
	NumberOfRegister string    `json:"number_of_register"`
	TotalScore       int       `json:"total_score"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func GetAllParticipant(page int, perPage int) *pagination.Paginator {
	var participants []UserParticipantResponse
	query := db.Connection.Table("user_peserta").
		Select("user_peserta.*, session_access.session_name").
		Joins("join session_access ON user_peserta.session_id = session_access.id")

	var offset = 0
	if page > 1 {
		offset = (page - 1) * perPage
	}

	if perPage == 0 {
		perPage = 10
	}

	paginator := pagination.Paging(&pagination.Param{
		DB:      query,
		Limit:   perPage,
		Offset:  offset,
		OrderBy: []string{"id desc"},
	}, &participants)

	for _, participant := range participants {
		participant = getResultExam(participant)
	}

	paginator.Records = participants
	return paginator
}

func getResultExam(user UserParticipantResponse) UserParticipantResponse {
	var totalCount = 0
	dataTypes := type_exam.GetAllTypeExam()
	totalType := len(dataTypes)

	for _, dataType := range dataTypes {
		totalCorrect := question.GetTotalCorrect(user.NumberOfRegister, dataType.ID)
		scoreCovert := exam.DictRewardExam[dataType.ID][totalCorrect]
		totalCount += scoreCovert
	}
	user.TotalScore = exam.CountTotalScore(totalCount, totalType)
	return user
}
