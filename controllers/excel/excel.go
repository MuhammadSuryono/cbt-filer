package controllers

import "github.com/gin-gonic/gin"

type IExcel interface {
	ImportExcelToDatabase(c *gin.Context)
	CreateGroupQuestion(c *gin.Context)
	Listening(c *gin.Context)
	UpdateGroupQuestion(c *gin.Context)
	ListUserParticipant(c *gin.Context)
}

type Excel struct {
}

func NewExcel() IExcel {
	return &Excel{}
}
