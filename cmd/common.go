package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Data struct {
}

type ApiLoginResponse struct {
	Data string `json:"data"`
	Meta struct {
		CurrentPage int `json:"currentPage"`
		LastPage    int `json:"lastPage"`
		Total       int `json:"total"`
		PerPage     int `json:"perPage"`
	} `json:"meta"`
}

type External interface {
	FetchData(ctx context.Context, id string) (*Data, error)
}

func getSesionId(sesameUrl, sesameUsername, sesamePassword string) (string, error) {
	values := map[string]string{"foo": "bar"}
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPost, sesameUrl+"/api/v3/security/login", bytes.NewBuffer(jsonData))
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	post := &ApiLoginResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post.Data, nil
}
