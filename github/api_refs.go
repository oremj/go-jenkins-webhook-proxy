package github

import (
	"fmt"
	"net/http"
)

type ApiRefResponse struct {
	Ref    string `json:"ref"`
	Url    string `json:"url"`
	Object struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"object"`
}

func (c *ApiClient) Refs(owner, repo, ref string) (*ApiRefResponse, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/git/refs/%s", GITHUB_API_URL, owner, repo, ref)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := new(ApiRefResponse)
	err = c.Do(req, res)

	return res, err
}
