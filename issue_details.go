package gong

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func GetBranchName(jiraClient *jira.Client, issueId string, issueType string) string {
	issue, _, _ := jiraClient.Issue.Get(issueId, nil)

	issueTitleSlug := SlugifyTitle(issue.Fields.Summary)
	return fmt.Sprintf("%s/%s-%s", issueType, issueId, issueTitleSlug)
}

func indexOf(status string, data []string) int {
	for k, v := range data {
		if status == v {
			return k
		}
	}
	return -1
}

func StartIssue(jiraClient *jira.Client, issueId string) error {
	allowed := []string{"Ready", "Start"}

	transitions, _, _ := jiraClient.Issue.GetTransitions(issueId)
	nextTransition := transitions[0]

	if indexOf(nextTransition.Name, allowed) > -1 {
		_, err := jiraClient.Issue.DoTransition(issueId, nextTransition.ID)

		if err != nil {
			return err
		}

		_ = StartIssue(jiraClient, issueId)
	}

	return nil
}
