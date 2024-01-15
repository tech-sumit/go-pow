package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetNonceMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockNonce := "testnonce"
	httpmock.RegisterResponder("GET", serverURL+"/challenge",
		httpmock.NewStringResponder(200, fmt.Sprintf(`{"nonce":"%s"}`, mockNonce)))

	nonce, err := getNonce()
	assert.Nil(t, err)
	assert.Equal(t, mockNonce, nonce)
}

func TestSendSolutionMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockNonce := "testnonce"
	mockSolution := "0000"
	mockQuote := "Knowing yourself is the beginning of all wisdom - Aristotle"

	httpmock.RegisterResponder("POST", serverURL+"/validate",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, gin.H{"quote": mockQuote})
		})

	quote, err := sendSolution(mockNonce, mockSolution)
	assert.Nil(t, err)
	assert.Equal(t, mockQuote, quote)
}
