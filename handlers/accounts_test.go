package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bank-it/models"
	"bank-it/testutils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountHandler(t *testing.T) {
	db := testutils.NewTestDB()
	router := gin.Default()
	RegisterAccountRoutes(router, db)

	t.Run("successfully creates account", func(t *testing.T) {
		body := `{"account_id": 1234, "balance": "500.00"}`
		req, _ := http.NewRequest("POST", "/accounts", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

		var account models.Account
		err := db.First(&account, "account_id = ?", 1234).Error
		assert.NoError(t, err)
		assert.Equal(t, 500.00, account.Balance)
	})

	t.Run("invalid payload", func(t *testing.T) {
		body := `{"account_id": 1234}`
		req, _ := http.NewRequest("POST", "/accounts", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestGetAccountHandler(t *testing.T) {
	db := testutils.NewTestDB()
	router := gin.Default()
	RegisterAccountRoutes(router, db)

	// Seed account
	db.Create(&models.Account{AccountID: 5678, Balance: 300.00})

	t.Run("successfully fetches account", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/accounts/5678", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

		var result map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &result)
		assert.Equal(t, float64(5678), result["account_id"])
		assert.Equal(t, 300.00, result["balance"])
	})

	t.Run("account not found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/accounts/9999", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}
