package test

import (
	"admin-server/admin/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type fakeDB struct {
	notices []models.Notice
	emails []models.Email
}

func (f *fakeDB) GetNotices() ([]models.Notice, error) {
	return f.notices, nil
}

func (f *fakeDB) GetNotice(id int) (models.Notice, error) {
	for _, n := range f.notices {
		if n.NoticeId == id {
			return n, nil
		}
	}
	return models.Notice{}, fmt.Errorf("notice with ID %d not found", id)
}

func (f *fakeDB) CreateNotice(n *models.NoticeUploadDto) error {
	return nil
}

func TestGetNotices(t *testing.T) {
	fakeNotices := []models.Notice{
		{
			NoticeId:       1,
			NoticeTitle:    "Test Notice 1",
			NoticeContents: "This is the first test notice",
			ProfileId:      uuid.New(),
		},
		{
			NoticeId:       2,
			NoticeTitle:    "Test Notice 2",
			NoticeContents: "This is the second test notice",
			ProfileId:      uuid.New(),
		},
	}

	db := &fakeDB{notices: fakeNotices}
	app := fiber.New()
	app.Get("/notices", func(c *fiber.Ctx) error {
		notices, err := db.GetNotices()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(notices)
	})

	req := httptest.NewRequest(http.MethodGet, "/notices", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, resp.StatusCode)
	}

	var result []models.Notice
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	if !reflect.DeepEqual(result, fakeNotices) {
		t.Errorf("Expected %v but got %v", fakeNotices, result)
	}

}

func TestGetNotice(t *testing.T) {
	app := fiber.New()

	mockNotice := models.Notice{
		NoticeId:       1,
		NoticeTitle:    "Mock notice",
		NoticeContents: "This is the content of the mock notice.",
	}

	// Define a fakeDB that returns the mock notice
	db := &fakeDB{
		notices: []models.Notice{mockNotice},
	}

	app.Get("/notice", func(c *fiber.Ctx) error {
		notices, err := db.GetNotice(1)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(notices)
	})

	req := httptest.NewRequest("GET", fmt.Sprintf("/notice?id=%d", mockNotice.NoticeId), nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, resp.StatusCode)
	}

}
