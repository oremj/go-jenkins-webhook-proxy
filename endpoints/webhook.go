package endpoints

import (
	"log"
	"net/http"
	"strings"

	"github.com/oremj/go-jenkins-api"
	"github.com/oremj/go-jenkins-webhook-proxy/github"
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

	if !payload.Created || strings.HasPrefix(payload.Ref, "refs/tags/") {
		http.Error(w, "not a tag", http.StatusOK)
		return
	}

	jobList, err := h.Jenkins.FetchJobList()
	if err != nil {
		http.Error(w, "could not parse payload", http.StatusBadRequest)
		log.Print("HandleWebhook: ", err)
		return
	}

	jobs := jobList.FilterByPropertyFunc(func(key, val string) bool {
		return strings.HasSuffix(key, "UpdateOnTag") && strings.HasPrefix(val, payload.Repository.FullName)
	})

	for _, j := range jobs {
		// TODO: build these jobs
		println("Update: ", j.Name)
		println("Tag: ", payload.Ref)
		println("PuppetGitRef: ", "refs/heads/master")
		println("SvcopRef: ", "refs/heads/master")
	}
}
