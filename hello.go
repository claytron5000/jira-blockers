package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	userEmail := flag.String("user", "", "Your email to access Jira")
	userToken := flag.String("token", "", "Your Jira token")
	issueID := flag.String("issue", "", "The issue id against which we're querying")
	flag.Parse()
	// fmt.Println(*issueID)
	// fmt.Println(*userEmail)
	// fmt.Println(*userToken)

	queryString := generateQueryString(*issueID)
	urlString := "https://bsn.atlassian.net/rest/api/latest/search?jql=" + queryString
	// fmt.Println(urlString)

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlString, nil)
	req.SetBasicAuth(DerefString(userEmail), DerefString(userToken))
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

	fmt.Println(s)
	var response JiraResponse
	json.Unmarshal([]byte(s), &response)
	for i := 0; i < len(response.Issues); i++ {
		fmt.Println(response.Issues[i].Key)
	}

}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	fmt.Println("nil argument")

	return ""
}

func generateQueryString(i string) string {

	templatePreface := " SYM AND issueBlocks "
	encodedPreface := url.PathEscape(templatePreface)
	// puking emoji face
	prefacePreface := "project%20="

	fmt.Println(i)

	return prefacePreface + encodedPreface + "=" + i
}

type Issue struct {
	Key string
}
type JiraResponse struct {
	Issues []Issue
}
