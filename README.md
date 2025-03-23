# CV Tailor: AI-Powered CV Optimization

CV Tailor is a powerful tool that leverages the Gemini AI model to tailor your CV to specific job descriptions. It analyzes your existing CV and the job description, then rewrites your CV to highlight the most relevant skills and experiences, making it more likely to pass through Applicant Tracking Systems (ATS) and impress hiring managers.

## Features

* **AI-Driven Tailoring:** Uses the Gemini AI model to intelligently rewrite your CV.
* **ATS Optimization:** Formats your CV to be easily parsed by Applicant Tracking Systems.
* **Concise and Relevant:** Focuses on the most relevant skills and experiences for each job.
* **Markdown Output:** Generates a well-structured CV in Markdown format.
* **PDF Extraction**: Extracts text from PDF files.
* **Downloadable Output**: Provides a download link for the tailored CV.

## Prerequisites

* **Go:** Ensure you have Go installed on your system.
* **Gemini API Key:** You'll need a Gemini API key from Google. Set it as an environment variable: `GEMINI_API_KEY`.
* **Dependencies**: Run `go mod tidy` to install the dependencies.

## Installation and Setup

1.  **Clone the Repository:**
    ```bash
    git clone <your-repository-url>
    cd <your-repository-directory>
    ```

2.  **Set the Gemini API Key:**
    ```bash
    export GEMINI_API_KEY="your_gemini_api_key"
    ```
    Replace `"your_gemini_api_key"` with your actual API key.

3.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Run the Application:**
    ```bash
    go run main.go
    ```
    This will start the server on `http://localhost:8080`.

## Usage

### Uploading CV and Job Description

To tailor your CV, you need to upload both your CV and the job description to the server. You can do this using `curl` as follows:

```bash
curl -X POST http://localhost:8080/upload \
  -F "cv=@your_cv.pdf" \
  -F "jobDesc=@job_description.pdf"
```

## Author

John Paul
