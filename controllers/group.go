package controllers

import (
	"github.com/gin-gonic/gin"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models"
	"gtihub.com/MuhammadSuryono/cbt-uploader/models/group"
	"path/filepath"
)

func (e *Excel) CreateGroupQuestion(c *gin.Context) {
	var param group.ParameterGroupQuestion
	_ = c.Bind(&param)

	file, _ := c.FormFile("file")
	filename := filepath.Base(file.Filename)

	err := c.SaveUploadedFile(file, filename)
	if err != nil {
		c.JSON(500, models.CommonResponse{
			Code:      500,
			IsSuccess: false,
			Message:   "Failed Upload File",
		})
		return
	}

	data := group.CreateGroupQuestion(param.ListQuestionId, param.Name, param.Description, filename, param.QuestionGroup)

	if data.Id != 0 {
		c.JSON(200, models.CommonResponse{
			Code:      200,
			IsSuccess: true,
			Message:   "Success Upload File",
			Data:      data,
		})
		return
	}

	c.JSON(500, models.CommonResponse{
		Code:      500,
		IsSuccess: false,
		Message:   "Failed Create Group",
	})
}

func (e *Excel) Listening(c *gin.Context) {
	var uri models.Uri
	_ = c.BindUri(&uri)
	c.File(uri.Filename)
}
