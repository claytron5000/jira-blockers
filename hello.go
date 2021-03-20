package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/xlab/treeprint"
)

func main() {

	issueID := flag.String("issue", "", "The issue id against which we're querying")
	userEmail := flag.String("user", "", "Your email to access Jira")
	userToken := flag.String("token", "", "Your Jira token")
	flag.Parse()
	// fmt.Println(*issueID)
	// fmt.Println(*userEmail)
	// fmt.Println(*userToken)
	tree := treeprint.New()

	top := tree.AddBranch(*issueID)

	// fmt.Println(tree.String())
	recurseTreeFetching(top, *issueID, *userEmail, *userToken)

	fmt.Println(tree.String())
}

func generateQueryString(i string) string {

	templatePreface := " SYM AND issueBlocks "
	encodedPreface := url.PathEscape(templatePreface)
	// puking emoji face
	prefacePreface := "project%20="

	// fmt.Println(i)

	return prefacePreface + encodedPreface + "=" + i
}

type Issue struct {
	Key string
}
type JiraResponse struct {
	Issues []Issue
}

func fetchBlockingIssues(issueID string, userEmail string, userToken string) []Issue {
	queryString := generateQueryString(issueID)
	urlString := "https://bsn.atlassian.net/rest/api/latest/search?jql=" + queryString
	// fmt.Println(urlString)

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlString, nil)
	req.SetBasicAuth(userEmail, userToken)
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

func recurseTreeFetching(tree treeprint.Tree, issueID string, userEmail string, userToken string) bool {
	issues := fetchBlockingIssues(issueID, userEmail, userToken)
	if len(issues) == 0 {
		return true
	}
	for i := 0; i < len(issues); i++ {
		currIssueID := issues[i].Key
		currBranch := tree.AddBranch(currIssueID)
		recurseTreeFetching(currBranch, currIssueID, userEmail, userToken)
	}
	return false
}
