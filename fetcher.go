package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type IFetcher interface {
	fetchIssues(issueID string) []Issue
}

type Fetcher struct {
	userEmail string
	userToken string
}

func (f Fetcher) fetchIssues(issueID string) []Issue {
	templatePreface := " SYM AND issueBlocks "
	encodedPreface := url.PathEscape(templatePreface)
	// puking emoji face
	prefacePreface := "project%20="
	queryString := prefacePreface + encodedPreface + "=" + issueID
	urlString := "https://bsn.atlassian.net/rest/api/latest/search?jql=" + queryString

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlString, nil)
	req.SetBasicAuth(f.userEmail, f.userToken)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	if err != nil {
		log.Fatal(err)
	}

	var response JiraResponse
	json.Unmarshal([]byte(s), &response)
	return response.Issues

}
