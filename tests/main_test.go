package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	os.Exit(m.Run())
}

func TestAddUserE2E(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("API_BASE_URL")
	orgID := os.Getenv("ORG_ID")

	if apiToken == "" || baseURL == "" || orgID == "" {
		t.Fatal("Missing required environment variables")
	}

	fullURL := baseURL + "identity.user.add"

	resp, err := sendRequest("POST", fullURL, apiToken, fmt.Sprintf(`{
		"emailLoginInvitation": true,
		"inviteToNetworkID": 0,
		"loginName": "teste2eansh2",
		"name": "TestE2EANSH2",
		"orgID": "%s",
		"password": "12345678",
		"passwordChangeRequired": true,
		"pskPassphrase": "8ikyjvq4yt",
		"status": "enabled",
		"userType": "local"
	}`, orgID))

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestGetUserE2E(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("API_BASE_URL")
	orgID := os.Getenv("ORG_ID")

	if apiToken == "" || baseURL == "" || orgID == "" {
		t.Fatal("Missing required environment variables")
	}

	fullURL := baseURL + "identity.user.get"

	resp, err := sendRequest("POST", fullURL, apiToken, fmt.Sprintf(`{
		"loginName": "teste2e",
		"orgID": "%s"
	}`, orgID))

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestUpdateUserE2E(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("API_BASE_URL")
	orgID := os.Getenv("ORG_ID")

	if apiToken == "" || baseURL == "" || orgID == "" {
		t.Fatal("Missing required environment variables")
	}

	fullURL := baseURL + "identity.user.update"

	// Replace with the actual userID of the created user
	userID := "U0197119b-0a8f-7cbd-88a6-005f1a5ec995"

	resp, err := sendRequest("POST", fullURL, apiToken, fmt.Sprintf(`{
		"orgID": "%s",
		"status": "disabled",
		"userID": "%s"
	}`, orgID, userID))

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestUserListE2E(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("API_BASE_URL")
	orgID := os.Getenv("ORG_ID")

	if apiToken == "" || baseURL == "" || orgID == "" {
		t.Fatal("Missing required environment variables")
	}

	fullURL := baseURL + "identity.user.list"

	resp, err := sendRequest("POST", fullURL, apiToken, fmt.Sprintf(`{
		"externalUserGroupID": 0,
		"limit": 10,
		"localUserGroupID": 0,
		"orgID": "%s"
	}`, orgID))

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestDeleteUserE2E(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")
	baseURL := os.Getenv("API_BASE_URL")
	orgID := os.Getenv("ORG_ID")

	if apiToken == "" || baseURL == "" || orgID == "" {
		t.Fatal("Missing required environment variables")
	}

	fullURL := baseURL + "identity.user.delete"

	resp, err := sendRequest("POST", fullURL, apiToken, fmt.Sprintf(`{
		"loginName": "testansh_e2e",
		"orgID": "%s"
	}`, orgID))

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}
