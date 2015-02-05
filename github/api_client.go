package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const GITHUB_API_URL = "https://api.github.com"

type ApiClient struct {
	Username string
	Password string
}

func (c *ApiClient) Do(req *http.Request, v interface{}) error {
	if c.Username != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("non-200 status code: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
