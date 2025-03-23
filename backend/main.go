package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nyunja/cv-tailor/internal"
)

func main() {
	router := gin.Default()

	internal.InitGemini()

	router.POST("/upload", internal.UploadHandler)
	// Download tailored CV as PDF or DOCX
	router.GET("/download/tailored_cv.pdf", internal.DownloadPDFHandler)

	log.Println("Server running on :8080")
	router.Run(":8080")
}
