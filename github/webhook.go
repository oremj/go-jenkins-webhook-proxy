package github

import (
	"encoding/json"
	"io"
)

type WebhookPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`
}

func ParseWebhookPayload(payload io.Reader) (*WebhookPayload, error) {
	res := new(WebhookPayload)
	err := json.NewDecoder(payload).Decode(res)
	return res, err
}
