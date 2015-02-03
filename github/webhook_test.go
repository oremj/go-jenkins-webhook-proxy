package github

import (
	"os"
	"testing"
)

func TestParseWebhookPayload(t *testing.T) {
	f, err := os.Open("fixtures/push.json")
	if err != nil {
		t.Fatal(err)
	}

	payload, err := ParseWebhookPayload(f)
	if err != nil {
		t.Fatal(err)
	}

	if payload.Ref != "refs/heads/gh-pages" {
		t.Errorf("payload.Ref is %s", payload.Ref)
	}

	if payload.Repository.FullName != "baxterthehacker/public-repo" {
		t.Errorf("payload.Repository.FullName is %s", payload.Repository.FullName)
	}

}
