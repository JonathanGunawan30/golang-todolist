package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"todolist-v1/config"
	"todolist-v1/modules/activity/entities"
	activityHandler "todolist-v1/modules/activity/handler"
	"todolist-v1/modules/activity/models"
	activityRepo "todolist-v1/modules/activity/repository"
	activityUsecase "todolist-v1/modules/activity/usecase"
	"todolist-v1/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ActivityTestSuite struct {
	suite.Suite
	app *fiber.App
	db  *database.PostgresDB
}

func (suite *ActivityTestSuite) SetupSuite() {
	cfg, err := config.LoadConfig()
	if err != nil {
		suite.T().Fatalf("Failed to load config: %v", err)
	}

	suite.db = database.NewPostgresDatabase()
	if err := suite.db.Connect(cfg.Database.URL); err != nil {
		suite.T().Fatalf("Failed to connect to test database: %v", err)
	}

	suite.app = fiber.New()
	repo := activityRepo.NewActivityRepository(suite.db.GetDB())
	usecase := activityUsecase.NewActivityUsecase(repo)
	handler := activityHandler.NewActivityHttpHandler(suite.app, usecase)
	handler.RegisterRoutes()
}

func (suite *ActivityTestSuite) TearDownTest() {
	suite.db.GetDB().Exec("TRUNCATE TABLE activities RESTART IDENTITY")
}

func TestActivityAPI(t *testing.T) {
	suite.Run(t, new(ActivityTestSuite))
}

func (suite *ActivityTestSuite) createSeedActivity() models.ActivityResponse {
	repo := activityRepo.NewActivityRepository(suite.db.GetDB())
	usecase := activityUsecase.NewActivityUsecase(repo)

	seed, _ := usecase.Create(entities.Activity{
		Title:        "Seed Task",
		Category:     "TASK",
		Description:  "A pre-existing task",
		ActivityDate: time.Now(),
	})

	return models.ActivityResponse{
		Id:           seed.Id,
		Title:        seed.Title,
		Category:     seed.Category,
		Description:  seed.Description,
		ActivityDate: seed.ActivityDate,
		Status:       seed.Status,
	}
}

func (suite *ActivityTestSuite) TestCreateActivity_Success() {
	body := bytes.NewBufferString(`{
		"title": "Integration Test Task",
		"category": "TASK",
		"description": "A task created from an integration test",
		"activity_date": "2025-10-10T10:00:00Z"
	}`)

	req, _ := http.NewRequest("POST", "/api/activities", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusCreated, resp.StatusCode)

	respBody, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	assert.Equal(suite.T(), "Activity created successfully", result["message"])
	data := result["data"].(map[string]interface{})
	assert.Equal(suite.T(), "Integration Test Task", data["title"])
	assert.Equal(suite.T(), "NEW", data["status"])
}

func (suite *ActivityTestSuite) TestCreateActivity_ValidationError() {
	body := bytes.NewBufferString(`{"title": ""}`)

	req, _ := http.NewRequest("POST", "/api/activities", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusBadRequest, resp.StatusCode)
}

func (suite *ActivityTestSuite) TestGetAllActivities_Success() {
	suite.createSeedActivity()

	req, _ := http.NewRequest("GET", "/api/activities", nil)
	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	respBody, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	data := result["data"].([]interface{})
	assert.Len(suite.T(), data, 1)
}

func (suite *ActivityTestSuite) TestGetAllActivities_Empty() {
	req, _ := http.NewRequest("GET", "/api/activities", nil)
	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	respBody, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	data := result["data"]
	assert.Nil(suite.T(), data, "The 'data' field should be nil when the list is empty")
}

func (suite *ActivityTestSuite) TestUpdateActivity_Success() {
	seed := suite.createSeedActivity()

	jsonBody := `{
		"title": "Updated Title",
		"category": "EVENT",
		"description": "Updated description",
		"activity_date": "2026-11-11T11:00:00Z",
		"status": "ON PROGRESS"
	}`
	updateBody := bytes.NewBufferString(jsonBody)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/activities/%d", seed.Id), updateBody)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode, "Status code should be 200 OK")

	if resp.StatusCode == fiber.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		var result map[string]interface{}
		json.Unmarshal(respBody, &result)

		assert.NotNil(suite.T(), result["data"], "Response data should not be nil")
		data := result["data"].(map[string]interface{})
		assert.Equal(suite.T(), "Updated Title", data["title"])
		assert.Equal(suite.T(), "ON PROGRESS", data["status"])
	}
}

func (suite *ActivityTestSuite) TestUpdateActivity_NotFound() {
	updateBody := bytes.NewBufferString(`{
		"title": "This is a valid title",
		"category": "TASK",
		"description": "This is a valid description",
		"activity_date": "2026-11-11T11:00:00Z",
		"status": "ON PROGRESS"
	}`)

	req, _ := http.NewRequest("PUT", "/api/activities/9999", updateBody)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusNotFound, resp.StatusCode)
}

func (suite *ActivityTestSuite) TestDeleteActivity_Success() {
	seed := suite.createSeedActivity()

	reqDelete, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/activities/%d", seed.Id), nil)
	respDelete, err := suite.app.Test(reqDelete)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, respDelete.StatusCode)

	reqGet, _ := http.NewRequest("GET", fmt.Sprintf("/api/activities/%d", seed.Id), nil)
	_, err = suite.app.Test(reqGet)
	assert.NoError(suite.T(), err)
}

func (suite *ActivityTestSuite) TestDeleteActivity_NotFound() {
	nonExistentID := "9999"

	req, _ := http.NewRequest("DELETE", "/api/activities/"+nonExistentID, nil)
	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusNotFound, resp.StatusCode)
}
