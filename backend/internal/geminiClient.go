package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	gemini "github.com/google/generative-ai-go/genai"
)

var geminiClient *gemini.Client

func InitGemini() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Failed to get API key")
	}
	ctx = context.Background()
	client, err := gemini.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	geminiClient = client
}

// AI Function to Tailor CV
func generateTailoredCV(cvText, jobDescText string) string {
	if geminiClient == nil {
		log.Fatal("Gemini API client is not initialized")
	}
	prompt := fmt.Sprintf(`
		You are an expert in professional resume writing.
		Given the following candidate's CV:
		---
		%s
		---
		And the job description:
		---
		%s
		---

		Rewrite the CV to be **one page** and optimized for **Applicant Tracking Systems (ATS)**.
		Make sure it is:
		- **Concise**
		- **Well-structured with bullet points**
		- **Highlights only the most relevant skills and experiences**
		- **Uses strong action verbs**
		- **Formatted properly for easy readability**
		
		### **Example Structure for the Updated CV**
		
		**[Candidate Name]**  
		_Professional Title_  
		📍 Location | 📧 Email | 📞 Phone | 🌐 LinkedIn  

		### **Professional Summary**  
		A brief 2-3 sentence overview emphasizing expertise relevant to the job.

		### **Key Skills**  
		- Skill 1 | Skill 2 | Skill 3  
		- Skill 4 | Skill 5 | Skill 6  

		### **Experience**  
		**Job Title – Company** _(Years of Experience)_  
		- Key achievement 1  
		- Key achievement 2  
		- Key achievement 3  

		**Previous Job Title – Company** _(Years of Experience)_  
		- Key achievement 1  
		- Key achievement 2  

		### **Education & Certifications**  
		- Degree – University Name _(Year)_  
		- Relevant Certification _(Year)_  
		
		Please generate the tailored CV in **Markdown format** so it maintains structure.
	`, cvText, jobDescText)

	// Call Gemini API
	model := geminiClient.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", resp.Candidates[0])
}
