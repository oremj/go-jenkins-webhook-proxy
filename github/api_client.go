package github

import (
	"encoding/json"
	"net/http"
)

const GITHUB_API_URL = "https://api.github.com"

type ApiClient struct {
	Username string
	Password string
}

func (c *ApiClient) Do(req *http.Request, v interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}
