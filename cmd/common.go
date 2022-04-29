package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ApiLoginResponse struct {
	Data string `json:"data"`
}

type ApiMeResponse struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

func getSessionId(sesameUrl, sesameUsername, sesamePassword string) (string, error) {
	values := map[string]string{"origin": "web"}
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPost, sesameUrl+"/api/v3/security/login", bytes.NewBuffer(jsonData))
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("error: could not connect to server")
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("error: unexpected server response")
	}
	defer res.Body.Close()
	post := &ApiLoginResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		return "", errors.New("error: unable to process server response")
	}
	return post.Data, nil
}

func getUserId(sesameUrl, sessionId string) (string, error) {
	values := map[string]string{"origin": "web"}
	jsonData, _ := json.Marshal(values)
	req, _ := http.NewRequest(http.MethodPost, sesameUrl+"/api/v3/security/me", bytes.NewBuffer(jsonData))
	req.Header.Add("Authorization", "Bearer "+sessionId)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("error: could not connect to server")
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("error: unexpected server response")
	}
	defer res.Body.Close()
	post := &ApiMeResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		return "", errors.New("error: unable to process server response")
	}
	return post.Data[0].ID, nil
}
