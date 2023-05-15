package test

import (
	"admin-server/admin/controllers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetEmails(t *testing.T) {
	// Create a new Fiber instance
	app := fiber.New()

	// Register a handler function for GET /emails
	app.Get("/emails", controllers.GetEmails)

	// Create a new HTTP request for GET /emails
	req := httptest.NewRequest(http.MethodGet, "/emails", nil)

	// Perform the request and record the response
	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// Verify that the status code is correct
	if res.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: expected %d, but got %d", http.StatusOK, res.StatusCode)
	}

	// Parse the response body
	var body struct {
		Error  bool        `json:"error"`
		Msg    interface{} `json:"msg"`
		Count  int         `json:"count"`
		Emails []string    `json:"emails"`
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Verify that the response body is correct
	if body.Error != false {
		t.Errorf("Unexpected error value: expected false, but got %v", body.Error)
	}

	if body.Msg != nil {
		t.Errorf("Unexpected error message: %v", body.Msg)
	}

	if body.Count != 0 {
		t.Errorf("Unexpected email count: expected 0, but got %d", body.Count)
	}

	if body.Emails != nil {
		t.Errorf("Unexpected email list: expected nil, but got %v", body.Emails)
	}
}
