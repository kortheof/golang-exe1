package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/trivago/tgo/tcontainer"
)

func main() {
	jiraClient, err := jira.NewClient(nil, "https://upstreamsystems.atlassian.net/")
	if err != nil {
		panic(err)
	}

	res, err := jiraClient.Authentication.AcquireSessionCookie("prdsuser", "#####")
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	//Handling the custom field of project ID in our Jira. Library tcontainer is used by Go-Jira.
	//https://github.com/trivago/tgo/blob/master/tcontainer/marshalmap.go
	customfield_11002 := tcontainer.NewMarshalMap()
	//map[string]interface{}
	customfield_11002["value"] = "All Projects"
	i := jira.Issue{
		Fields: &jira.IssueFields{
			Reporter: &jira.User{
				Name: "fanis.korlos",
			},
			Description: "Test Golang Go-Jira",
			//Mandatory field, it's a PRDS ticket
			Project: jira.Project{
				Key: "PRDS",
			},
			//Mandatory field, Title of the ticket
			Summary: "Golang Demo Ticket",
			//Mandatory field, Issue type of the ticket
			Type: jira.IssueType{
				Name: "General Task",
			},
			//Mandatory field, ProjectID -> represented by customfield_11002 in our Jira
			//https://upstreamsystems.atlassian.net/rest/api/2/issue/PRDS-2099
			Unknowns: customfield_11002,
		},
	}
	issue, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
