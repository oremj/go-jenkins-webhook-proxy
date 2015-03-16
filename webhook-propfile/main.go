package main

import (
	"flag"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/oremj/go-jenkins-webhook-proxy/github"
)

var webhookData = flag.String("data", "", "full webhook json blob")
var startPrefix = "WEBHOOK_"

func printField(prefix, fieldName string, field reflect.Value) {
	kind := field.Type().Kind()
	if kind == reflect.String {
		fmt.Printf("%s%s=%s\n", prefix, fieldName, field.String())
		return
	}

	if kind == reflect.Bool {
		if field.Bool() {
			fmt.Printf("%s%s=TRUE\n", prefix, fieldName)
		}
		fmt.Printf("%s%s=FALSE\n", prefix, fieldName)
		return
	}

	if kind == reflect.Struct {
		printStruct(prefix+fieldName+"_", field)
		return
	}
}

func printStruct(prefix string, s reflect.Value) {
	sType := s.Type()
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		fieldName := field.Name
		tag := field.Tag
		if tag != "" {
			fieldName = tag.Get("json")
		}
		printField(prefix, strings.ToUpper(fieldName), s.Field(i))
	}
}

func main() {

	flag.Parse()
	if *webhookData == "" {
		log.Fatal("-data flag cannot be empty.")
	}

	payload, err := github.ParseWebhookPayload(strings.NewReader(*webhookData))
	if err != nil {
		log.Fatalf("Parse error: %s", err)
	}

	printStruct(startPrefix, reflect.ValueOf(*payload))
}
