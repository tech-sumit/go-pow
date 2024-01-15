package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gopow/pkg"
	"io/ioutil"
	"net/http"
)

const (
	serverURL = "http://localhost:8080"
)

func main() {
	// Setup the logger
	logger, err := pkg.SetupLogger()
	if err != nil {
		fmt.Println("Error setting up logger:", err)
		return
	}
	defer logger.Sync()

	// Log a starting message
	logger.Info("Client started")

	// Call the actual client logic
	if err := runClient(logger); err != nil {
		logger.Errorf("Client error: %v", err)
	}
}

func runClient(log *zap.SugaredLogger) error {
	// Step 1: Get nonce from server
	nonce, err := getNonce()
	if err != nil {
		log.Errorf("Error getting nonce: %v", err)
		return err
	}

	log.Infof("Received nonce: %s", nonce)

	// Step 2: Solve PoW
	solution, err := pkg.SolvePoW(nonce)
	if err != nil {
		log.Errorf("Error solving PoW: %v", err)
		return err
	}

	log.Infof("Solved PoW: %s", solution)

	// Step 3: Send solution and get quote
	quote, err := sendSolution(nonce, solution)
	if err != nil {
		log.Errorf("Error sending solution: %v", err)
		return err
	}

	log.Infof("Received quote: %s", quote)

	return nil
}

func getNonce() (string, error) {
	resp, err := http.Get(serverURL + "/challenge")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["nonce"], nil
}

func sendSolution(nonce, solution string) (string, error) {
	solutionData := map[string]string{
		"nonce": nonce,
		"hash":  solution,
	}
	jsonData, _ := json.Marshal(solutionData)

	resp, err := http.Post(serverURL+"/validate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result["quote"], nil
}
