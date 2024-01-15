package main

import (
	"github.com/gin-gonic/gin"
	"gopow/pkg"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"Knowing yourself is the beginning of all wisdom - Aristotle",
	"The only true wisdom is in knowing you know nothing - Socrates",
	"The fool doth think he is wise, but the wise man knows himself to be a fool - Shakespeare",
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

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

	r.Run() // Listen and serve on 0.0.0.0:8080
}

func randomQuote() string {
	rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))]
}
