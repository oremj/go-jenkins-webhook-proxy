package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/oremj/go-jenkins-webhook-proxy/endpoints"
	"github.com/oremj/go-jenkins-webhook-proxy/jenkins"
)

var jenkinsUserName = flag.String("jenkins-user", "", "Jenkins Username")
var jenkinsPassword = flag.String("jenkins-pass", "", "Jenkins Password")
var jenkinsBaseURL = flag.String("jenkins-base-url", "", "Example: https://deploy.jenkins.com")

var serverAddr = flag.String("addr", ":8080", "Listen address")

func main() {

	flag.Parse()

	jenkinsApi := jenkins.NewApi(*jenkinsUserName, *jenkinsPassword, *jenkinsBaseURL)

	webhookHandler := &endpoints.WebhookHandler{
		Jenkins: jenkinsApi,
	}

	mux := http.NewServeMux()
	mux.Handle("/webhook/", webhookHandler)

	log.Fatal(http.ListenAndServe(*serverAddr, mux))
}
