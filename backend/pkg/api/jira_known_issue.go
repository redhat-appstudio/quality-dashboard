package api

import (
	"fmt"
	"log"
	"net/http"

	jira "github.com/andygrunwald/go-jira"
	client "github.com/flacatus/qe-dashboard-backend/pkg/api/apis/jira"
)

const (
	KNOWN_E2E_ISSUE_LABEL = "appstudio-e2e-tests-known-issues"
)

// Jira godoc
// @Summary Jira
// @Description returns a list of jira issues which contain the label appstudio-e2e-tests-known-issues
// @Tags Version API
// @Produce json
// @Router /api/jira/e2e-known/get [get]
// @Success 200 {object} api.MapResponse
func (s *Server) getE2eKnownIssues(w http.ResponseWriter, r *http.Request) {
	factory := client.NewTotClientFactory()
	jiraClient, err := factory.NewJiraClient()
	if err != nil {
		log.Fatal(err)
	}
	var issues []jira.Issue

	// append the jira issues to []jira.Issue
	appendFunc := func(i jira.Issue) (err error) {
		issues = append(issues, i)
		return err
	}

	// In this example, we'll search for all the issues with the provided JQL filter and Print the Story Points
	err = jiraClient.Issue.SearchPages(fmt.Sprintf("labels in (%s) AND status not in (resolved, closed)", KNOWN_E2E_ISSUE_LABEL), nil, appendFunc)
	if err != nil {
		log.Fatal(err)
	}
	s.JSONResponse(w, r, issues)
}
