package tests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"github.com/fronbasal/TanteMateLaden-go/tests"
)

// Login logs in a user and returns the token
func Login(username, password string, t *testing.T) string {

	body := strings.NewReader(`{"username": "` + username + `", "password": "` + password + `"}`)
	req, err := http.NewRequest("POST", "http://localhost:3000/login", body)
	if err != nil {
		t.Error("Failed to create request")
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Could not submit request!")
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Error("Did not return status 200, returned: " + strconv.Itoa(resp.StatusCode) + "\n" + string(b))
		return ""
	}
	var r tests.TokenResponse
	json.NewDecoder(resp.Body).Decode(&r)
	return r.Token
}

// Query runs an authorized query
func Query(token, method, url string, t *testing.T) io.ReadCloser {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Error("Could not create request")
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Could not send request")
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Error("Non-200 status code: " + strconv.Itoa(resp.StatusCode))
		return nil
	}
	return resp.Body
}

// PostQuery runs a query
func PostQuery(data, url, token string, t *testing.T) io.ReadCloser {
	body := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		t.Error("Could not create request")
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Could not submit request: " + err.Error())
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Error("Did not return status 200. Got: " + strconv.Itoa(resp.StatusCode))
		return nil
	}
	return resp.Body
}
