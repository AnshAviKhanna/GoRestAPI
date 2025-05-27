// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"
// )

// // func main() {
// // 	// Load environment variables from .env file
// // 	err := godotenv.Load()
// // 	if err != nil {
// // 		fmt.Println("Error loading .env file")
// // 		return
// // 	}

// // 	apiToken := os.Getenv("API_TOKEN")
// // 	baseURL := os.Getenv("API_BASE_URL")
// // 	orgID := os.Getenv("ORG_ID")

// // 	AddUser(apiToken, baseURL+"identity.user.add", orgID)
// // 	// GetUser(apiToken, baseURL+"identity.user.get", orgID)
// // 	// UpdateUser(apiToken, baseURL+"identity.user.update", orgID)
// // 	// UserList(apiToken, baseURL+"identity.user.list", orgID)
// // 	// DeleteUser(apiToken, baseURL+"identity.user.delete", orgID)
// // }

// // GetUser sends a request to the "identity.user.get" endpoint
// func GetUser(apiToken, fullURL, orgID string) {
// 	// endpoint := "identity.user.get"
// 	// fullURL := strings.TrimRight(baseURL, "/") + "/" + endpoint

// 	// Create request body
// 	// Sprintf - for multiline formatted string
// 	requestBody := fmt.Sprintf(`{
// 		"loginName": "testansh2",
// 		"orgID": "%s"
// 	}`, orgID)

// 	// Send request using helper function
// 	resp, err := sendRequest("POST", fullURL, apiToken, requestBody)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}

// 	printResponse(resp)
// }

// // AddUser sends a request to the "identity.user.add" endpoint
// func AddUser(apiToken, fullURL, orgID string) {
// 	// endpoint := "identity.user.add"
// 	// fullURL := strings.TrimRight(baseURL, "/") + "/" + endpoint

// 	// Create request body
// 	requestBody := fmt.Sprintf(`{
// 	"emailLoginInvitation": true,
// 	"inviteToNetworkID": 0,
// 	"loginName": "teste2eansh",
// 	"name": "TestE2EANSH",
// 	"orgID": "%s",
// 	"password": "12345678",
// 	"passwordChangeRequired": true,
// 	"pskPassphrase": "8lkyjvq4yt",
// 	"status": "enabled",
// 	"userType": "local"
// 	}`, orgID)

// 	// Send request using helper function
// 	resp, err := sendRequest("POST", fullURL, apiToken, requestBody)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}

// 	printResponse(resp)
// }

// // DeleteUser sends a request to the "identity.user.delete" endpoint
// func DeleteUser(apiToken, fullURL, orgID string) {
// 	// endpoint := "identity.user.delete"
// 	// fullURL := strings.TrimRight(baseURL, "/") + "/" + endpoint

// 	// Create request body
// 	requestBody := fmt.Sprintf(`{
// 		"loginName": "testansh3",
// 		"orgID": "%s"
// 	}`, orgID)

// 	// Send request using helper function
// 	resp, err := sendRequest("POST", fullURL, apiToken, requestBody)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}

// 	printResponse(resp)
// }

// // UserList sends a request to the "identity.user.list" endpoint
// func UserList(apiToken, fullURL, orgID string) {
// 	// endpoint := "identity.user.list"
// 	// fullURL := strings.TrimRight(baseURL, "/") + "/" + endpoint

// 	// Create request body
// 	requestBody := fmt.Sprintf(`{
// 		"externalUserGroupID": 0,
// 		"limit": 0,
// 		"localUserGroupID": 0,
// 		"orgID": "%s"
// 	}`, orgID)

// 	// Send request using helper function
// 	resp, err := sendRequest("POST", fullURL, apiToken, requestBody)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}

// 	printResponse(resp)
// }

// // UpdateUser sends a request to the "identity.user.update" endpoint
// func UpdateUser(apiToken, fullURL, orgID string) {
// 	// endpoint := "identity.user.update"
// 	// fullURL := strings.TrimRight(baseURL, "/") + "/" + endpoint

// 	// Create request body
// 	requestBody := fmt.Sprintf(`{
// 		"orgID": "%s",
// 		"status": "disabled",
//   		"userID": "U01970bdc-f54f-7443-8544-c5887406b17a"
// 	}`, orgID)

// 	// Send request using helper function
// 	resp, err := sendRequest("POST", fullURL, apiToken, requestBody)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}

// 	printResponse(resp)
// }

// // sendRequest is a reusable helper to create and send HTTP requests
// func sendRequest(method, url, token, body string) (*http.Response, error) {
// 	// Creating a new request (type, url, body)
// 	// ------------------------------------------------------------------------
// 	// requestBody is a []byte, so we wrap it with bytes.NewBuffer() to turn it into a readable stream for http.NewRequest
// 	// req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(requestBody))
// 	// byte slice format is not needed here
// 	// ------------------------------------------------------------------------
// 	req, err := http.NewRequest(method, url, strings.NewReader(body))
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating request: %w", err)
// 	}

// 	// Set request headers
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "application/json")

// 	// Send request
// 	client := &http.Client{}
// 	return client.Do(req)
// }

// // printResponse reads and pretty-prints the response
// func printResponse(resp *http.Response) {
// 	// ------------------------------------------------------------------------
// 	// defer means "run this function later, when the current function returns."
// 	defer resp.Body.Close()

// 	// Read raw response
// 	respBody, _ := io.ReadAll(resp.Body)

// 	// Pretty print
// 	var prettyJSON bytes.Buffer
// 	err := json.Indent(&prettyJSON, respBody, "", "  ")
// 	if err != nil {
// 		// fallback to raw response
// 		fmt.Printf("Status: %d\nRaw Response:\n%s\n", resp.StatusCode, respBody)
// 	} else {
// 		fmt.Printf("Status: %d\nPretty JSON Response:\n%s\n", resp.StatusCode, prettyJSON.String())
// 	}
// }
