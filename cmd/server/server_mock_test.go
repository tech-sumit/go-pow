package main

import (
	"bytes"
	"encoding/json"
	"gopow/pkg"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/challenge", func(c *gin.Context) {
		nonce := pkg.GenerateNonce()
		c.JSON(http.StatusOK, gin.H{"nonce": nonce})
	})
	r.POST("/validate", func(c *gin.Context) {
		var req struct {
			Nonce string `json:"nonce"`
			Hash  string `json:"hash"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if pkg.ValidatePoW(req.Nonce, req.Hash) {
			c.JSON(http.StatusOK, gin.H{"quote": randomQuote()})
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid PoW"})
		}
	})
	return r
}

func TestChallengeEndpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/challenge", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "nonce")
}

func TestValidateEndpoint(t *testing.T) {
	router := setupRouter()

	// Step 1: Get nonce from server
	nonce := pkg.GenerateNonce()

	// Step 2: Solve PoW
	solution, err := pkg.SolvePoW(nonce)
	if err != nil {
		t.Errorf("Error solving PoW: %v", err)
		return
	}

	// Step 3: Send solution and get quote
	w := httptest.NewRecorder()
	jsonData := map[string]string{
		"nonce": nonce,
		"hash":  solution,
	}
	reqBody, _ := json.Marshal(jsonData)
	req, _ := http.NewRequest("POST", "/validate", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response: %v", err)
		return
	}

	if _, ok := response["quote"]; !ok {
		t.Errorf("Response does not contain 'quote' key")
	}
}
