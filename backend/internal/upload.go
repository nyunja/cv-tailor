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
	uploadsDir := "uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		os.Mkdir(uploadsDir, 0755)
	}
	cvPath := fmt.Sprintf("uploads/%s", cv.Filename)
	jobDescPath := fmt.Sprintf("uploads/%s", jobDesc.Filename)
	if err := c.SaveUploadedFile(cv, cvPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save CV"})
        return
	}
	if err := c.SaveUploadedFile(jobDesc, jobDescPath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save job description"})
        return
    }

	// Extract text
	cvText, err := extractTextFromPDF(cvPath)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract text from cv"})
        return
    }
	jobDescText, err := extractTextFromPDF(jobDescPath)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract text from job descriptor"})
        return
    }
	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully", "cvText": cvText, "jobDescText": jobDescText})

	// Process the files (implement NLP matching & tailoring)
	tailoredCV := generateTailoredCV(cvText, jobDescText)
	if tailoredCV == "" {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tailored CV"})
        return
    }
	c.JSON(http.StatusOK, gin.H{"message": "CV processed", "tailoredCV": tailoredCV})

	// Save processed CV
	outputPath := "output/tailored_cv.md"
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		os.MkdirAll("output", 0755)
	}
	err = os.WriteFile(outputPath, []byte(tailoredCV), 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tailored CV"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CV processed", "downloadUrl": "/download/tailored_cv.pdf"})
}
