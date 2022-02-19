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
		api.POST("/group/question/create", controller.CreateGroupQuestion)
		api.GET("/group/question/listening/:filename", controller.Listening)
		api.POST("/group/question/update/:id", controller.UpdateGroupQuestion)
	}

	server.Run(":8080")
}
