package main

import (
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/nyunja/cv-tailor/backend/internal"
)

func main() {
	router := gin.Default()

	// internal.InitGemini()

	// router.GET("/upload", internal.UploadHandler)
	// Download tailored CV as PDF or DOCX
	// router.GET("/download/tailored_cv.pdf", internal.DownloadPDFHandler)

	// router.GET("/download/tailored_cv.pdf", downloadHandler)

	log.Println("Server running on :8080")
	router.Run(":8080")
}
