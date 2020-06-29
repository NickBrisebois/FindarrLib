package findarrlib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	_ "net/http"
)
import _ "errors"

type TVInfo struct  {

}

func NewTVInfo() *TVInfo {
	newTV := &TVInfo{}
	return newTV
}

// GetJWT gets a JWT token to authorize TVdb api requests. Requires a username, userKey and apiKey
// which can all be obtained from thetvdb.com
func (tv *TVInfo) GetJWT(username string, userKey string, apiKey string) (string, error) {
	loginURL := "https://api.thetvdb.com/swagger/login"

	// Make request to TVdb API
	jsonValues := map[string]string {
		"apikey": apiKey,
		"userkey": userKey,
		"username": username,
	}
	reqJSON, _ := json.Marshal(jsonValues)

	loginResp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	defer loginResp.Body.Close()
	var respData map[string]interface{}
	err = json.NewDecoder(loginResp.Body).Decode(&respData)
	if err != nil {
		return "", err
	}

	if respData["Error"] != nil {
		// Convert error to string
		errorValue := fmt.Sprint(respData["Error"])
		err = errors.New(errorValue)
		return "", err
	}

	return "", err
}

