package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, err := jira.NewClient(nil, "https://upstreamsystems.atlassian.net/")
	if err != nil {
		panic(err)
	}
	jiraClient.Authentication.SetBasicAuth("prdsuser", "PRD$123#")

	issue, _, err := jiraClient.Issue.Get("PRDS-4571", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
