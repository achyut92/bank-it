package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"bank-it/handlers"
	"bank-it/models"
	"bank-it/testutils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionHandler(t *testing.T) {
	db := testutils.NewTestDB()
	router := gin.Default()
	handlers.RegisterTransactionRoutes(router, db)
	handlers.RegisterAccountRoutes(router, db)

	// Seed accounts
	db.Create(&models.Account{AccountID: 1001, Balance: 1000.00})
	db.Create(&models.Account{AccountID: 1002, Balance: 500.00})

	t.Run("successful transaction", func(t *testing.T) {
		body := `{
			"source_account_id": 1001,
			"destination_account_id": 1002,
			"amount": "200.00"
		}`
		req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var src models.Account
		var dst models.Account
		db.First(&src, "account_id = ?", 1001)
		db.First(&dst, "account_id = ?", 1002)

		assert.Equal(t, 800.00, src.Balance)
		assert.Equal(t, 700.00, dst.Balance)

		var txns []models.Transaction
		db.Find(&txns)
		assert.Equal(t, 2, len(txns))
	})

	t.Run("insufficient funds", func(t *testing.T) {
		body := `{
			"source_account_id": 1001,
			"destination_account_id": 1002,
			"amount": "5000.00"
		}`
		req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("invalid payload", func(t *testing.T) {
		body := `{"source_account_id": 1001}`
		req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("invalid destination account", func(t *testing.T) {
		body := `{
			"source_account_id": 1001,
			"destination_account_id": 1005,
			"amount": "50.00"
		}`
		req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}
