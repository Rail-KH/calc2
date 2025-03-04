package orchestrator

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {
	orch := NewOrchestrator()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orch.calculateHandler(w, r)
	})

	reqBody := `{"expression": "(42-52)*37"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rr.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding the response: %v", err)
	}
	if id, ok := resp["id"]; !ok || id == "" {
		t.Errorf("Expected valid id in the response, got: %v", resp)
	}
}
