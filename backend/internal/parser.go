package internal

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// Extract text from PDF using pdfcpu
func extractTextFromPDF(filePath string) string {
	tmpFile := "extract_text.txt"
	defer os.Remove(tmpFile)

	err := api.ExtractContentFile(filePath, tmpFile, nil, nil)
	if err != nil {
		log.Println("Error extracting content from PDF:", err)
		return ""
	}
	content, err := os.ReadFile(tmpFile)
	if err != nil {
		log.Println("Error reading temp file:", err)
		return ""
	}
	return strings.TrimSpace(string(content))
}

// Handlers to Download Files
func DownloadPDFHandler(c *gin.Context) {
	c.File("output/tailored_cv.pdf")
}
