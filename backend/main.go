package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	gemini "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var geminiClient *gemini.Client

func main() {
	router := gin.Default()

	initGemini()

	// router.GET("/upload", uploadHandler)

	// router.GET("/download/tailored_cv.pdf", downloadHandler)

	log.Println("Server running on :8080")
	router.Run(":8080")
}

func initGemini() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Failed to get API key")
	}
	client, err := gemini.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	geminiClient = client
}
