package controllers

import "github.com/gin-gonic/gin"

type IExcel interface {
	ImportExcelToDatabase(c *gin.Context)
}

type Excel struct {
}

func NewExcel() IExcel {
	return &Excel{}
}
