package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("API_BASE_URL")
	orgID := os.Getenv("ORG_ID")

	endpoint := "identity.user.get"
	fullURL := strings.TrimRight(baseURL, "/") + "/" + endpoint

	// Create request body
	// Sprintf - for multiline formatted string
	requestBody := fmt.Sprintf(`{
	"loginName": "ansh",
	"orgID": "%s",
	"userID": "AAK"
	}`, orgID)

	// Creating a new request (type,url,response format)
	// ------------------------------------------------------------------------
	// requestBody is a []byte, so we wrap it with bytes.NewBuffer() to turn it into a readable stream for http.NewRequest
	// req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(requestBody))
	// byte slice format is not needed here
	// ------------------------------------------------------------------------
	req, err := http.NewRequest("POST", fullURL, strings.NewReader(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// ------------------------------------------------------------------------
	// defer means "run this function later, when the current function returns."
	defer resp.Body.Close()

	// Read raw response
	respBody, _ := io.ReadAll(resp.Body)

	// pretty print
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, respBody, "", "  ")
	if err != nil {
		// fallback to raw response
		fmt.Printf("Status: %d\nRaw Response:\n%s\n", resp.StatusCode, respBody)
	} else {
		fmt.Printf("Status: %d\nPretty JSON Response:\n%s\n", resp.StatusCode, prettyJSON.String())
	}
}
