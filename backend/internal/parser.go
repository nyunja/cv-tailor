package internal

import (
	"bytes"
	"fmt"
	"log"

	"github.com/dslipak/pdf"
	"github.com/gin-gonic/gin"
)

// Extract text from PDF using pdfcpu
func extractTextFromPDF(filePath string) (string, error) {
	file, err := pdf.Open(filePath)
	if err != nil {
		log.Printf("Error opening PDF file %s: %v", filePath, err)
		return "", fmt.Errorf("error opening PDF file %s: %v", filePath, err)
	}
	var buf bytes.Buffer
	b, err := file.GetPlainText()
	if err != nil {
		return "", fmt.Errorf("error extracting text from PDF file %s: %v", filePath, err)
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

// Handlers to Download Files
func DownloadPDFHandler(c *gin.Context) {
	c.File("output/tailored_cv.pdf")
}
