package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var ctx context.Context

// Upload handler
func UploadHandler(c *gin.Context) {
	cv, err := c.FormFile("cv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CV upload failed"})
		return
	}

	jobDesc, err := c.FormFile("jobDesc")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job Description upload failed"})
		return
	}

	// Save files temporarily
	cvPath := fmt.Sprintf("uploads/%s", cv.Filename)
	jobDescPath := fmt.Sprintf("uploads/%s", jobDesc.Filename)

	_ = c.SaveUploadedFile(cv, cvPath)
	_ = c.SaveUploadedFile(jobDesc, jobDescPath)

	// Extract text
	cvText := extractTextFromPDF(cvPath)
	jobDescText := extractTextFromPDF(jobDescPath)

	// Process the files (implement NLP matching & tailoring)
	tailoredCV := generateTailoredCV(cvText, jobDescText)

	// Save processed CV
	outputPath := "output/tailored_cv.pdf"
	err = os.WriteFile(outputPath, []byte(tailoredCV), 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tailored CV"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CV processed", "downloadUrl": "/download/tailored_cv.pdf"})
}
