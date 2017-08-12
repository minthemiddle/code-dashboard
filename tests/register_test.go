package tests

import (
	"net/http"
	"strings"
	"testing"

	"encoding/json"
	"strconv"
)

func TestRegister(t *testing.T) {
	body := strings.NewReader(`{"email": "test@localhost", "password": "test123"}`)
	req, err := http.NewRequest("POST", "http://localhost:3000/register", body)
	if err != nil {
		t.Error("Failed to make request: " + err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Failed to make request: " + err.Error())
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(nil)
	if resp.StatusCode != 200 {
		t.Error("Did not receive Status 200. Got: " + strconv.Itoa(resp.StatusCode))
		return
	}
}
