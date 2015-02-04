package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oremj/go-jenkins-webhook-proxy/github"
	"github.com/oremj/go-jenkins-webhook-proxy/jenkins"
)

type WebhookHandler struct {
	Jenkins *jenkins.Api
}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	payload, err := github.ParseWebhookPayload(req.Body)
	if err != nil {
		http.Error(w, "could not parse payload", http.StatusBadRequest)
		log.Print("HandleWebhook: ", err)
		return
	}
	fmt.Println(payload.Ref)
}
