package jenkins

import (
	"encoding/json"
	"io"
	"net/http"
)

type JenkinsApi struct {
	Username string
	Password string
	URL      string // e.g., https://deploy.jenkins.com

	Client *http.Client
}

func NewJenkinsApi(username, password, url string) *JenkinsApi {
	return &JenkinsApi{
		Username: username,
		Password: password,
		URL:      url,
		Client:   &http.Client{},
	}
}

type JenkinsApiJobListResponse struct {
	Jobs []struct {
		Name     string `json:"name"`
		Property []struct {
			Parameters []struct {
				Name     string `json:"name"`
				Defaults struct {
					Value string `json:"value"`
				} `json:"defaultParameterValue"`
			} `json:"parameterDefinitions"`
		} `json:"property"`
	}
}

func (j *JenkinsApi) BuildURL(path string) string {
	return j.URL + path
}

func (j *JenkinsApi) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(j.Username, j.Password)
	return j.Client.Do(req)
}

func (j *JenkinsApi) Get(v interface{}, path string) error {
	req, err := http.NewRequest("GET", j.BuildURL(path), nil)
	if err != nil {
		return err
	}

	return j.doRequest(v, req)
}

func (j *JenkinsApi) Post(v interface{}, path string, body io.Reader) error {
	req, err := http.NewRequest("POST", j.BuildURL(path), body)
	if err != nil {
		return err
	}

	return j.doRequest(v, req)
}

func (j *JenkinsApi) doRequest(v interface{}, req *http.Request) error {
	resp, err := j.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

func (j *JenkinsApi) FetchJobList() (*JenkinsApiJobListResponse, error) {
	resp := new(JenkinsApiJobListResponse)

	err := j.Get(resp, "/api/json?pretty=true&tree=jobs[name,property[parameterDefinitions[name,defaultParameterValue[value]]]]")
	return resp, err
}
