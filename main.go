package main

import (
	"github.com/MuhammadSuryono/go-helper/db"
	"github.com/joho/godotenv"
	"gtihub.com/MuhammadSuryono/cbt-uploader/controllers"
)

func main() {
	_ = godotenv.Load()
	db.InitConnectionFromEnvironment().CreateNewConnection()

	server := controllers.RunServer()

	controller := controllers.NewExcel()
	api := server.Group("api/v1/excel")
	{
		api.POST("/upload/question/:list_question_id", controller.ImportExcelToDatabase)
	}

	server.Run(":8080")
}
