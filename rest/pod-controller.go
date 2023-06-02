package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	InputTopic    string `json:"input-topic"`
	ProcessorName string `json:"processor-name"`
	OutputTopic   string `json:"output-topic"`
}

func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	var req Request

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Perform processing logic
	// You can replace the logic below with your desired implementation
	result := fmt.Sprintf("Received request with Input Topic: %s, Processor Name: %s, Output Topic: %s", req.InputTopic, req.ProcessorName, req.OutputTopic)

	// Return the result as a JSON response
	resp := map[string]interface{}{
		"result": result,
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
