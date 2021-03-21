package main

import (
	"flag"
	"fmt"

	"github.com/xlab/treeprint"
)

func main() {

	issueID := flag.String("issue", "", "The issue id against which we're querying")
	userEmail := flag.String("user", "", "Your email to access Jira")
	userToken := flag.String("token", "", "Your Jira token")
	depth := flag.Int("depth", 3, "How deep do we build the tree")
	flag.Parse()
	// fmt.Println(*issueID)
	// fmt.Println(*userEmail)
	// fmt.Println(*userToken)
	fetcher := Fetcher{
		userEmail: *userEmail,
		userToken: *userToken,
	}

	tree := treeprint.New()

	var done = make(chan bool)

	go recurseTreeFetching(fetcher, tree, *issueID, done, *depth)
	<-done

	fmt.Println(tree.String())

}

func recurseTreeFetching(fetcher Fetcher, tree treeprint.Tree, issueID string, done chan bool, depth int) {
	issues := fetcher.fetchIssues(issueID)
	if len(issues) == 0 {
		return
	}
	depth--
	if depth == 0 {
		done <- true
		return
	}

	for i := 0; i < len(issues); i++ {
		currIssueID := issues[i].Key
		currBranch := tree.AddBranch(currIssueID)
		go recurseTreeFetching(fetcher, currBranch, currIssueID, done, depth)
	}
}

func Crawl(fetcher Fetcher, issueID string, ch chan int, tree treeprint.Tree) {

	issues := fetcher.fetchIssues(issueID)
	fmt.Println("number issues", len(issues))
	ch <- len(issues)
	if len(issues) > 0 {
		fmt.Println(len(issues))
		tree.AddBranch(issueID)
		for _, iss := range issues {
			fmt.Println(iss)
			go Crawl(fetcher, iss.Key, ch, tree)
		}
	} else {
		fmt.Println("else")
		tree.AddNode(issueID)
	}

}

// func generateQueryString(i string) string {

// 	templatePreface := " SYM AND issueBlocks "
// 	encodedPreface := url.PathEscape(templatePreface)
// 	// puking emoji face
// 	prefacePreface := "project%20="

// 	// fmt.Println(i)

// 	return prefacePreface + encodedPreface + "=" + i
// }

type Issue struct {
	Key string
}
type JiraResponse struct {
	Issues []Issue
}

// func fetchBlockingIssues(issueID string, userEmail string, userToken string) []Issue {
// 	queryString := generateQueryString(issueID)
// 	urlString := "https://bsn.atlassian.net/rest/api/latest/search?jql=" + queryString
// 	// fmt.Println(urlString)

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", urlString, nil)
// 	req.SetBasicAuth(userEmail, userToken)
// 	req.Header.Add("Content-Type", "application/json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	bodyText, err := ioutil.ReadAll(resp.Body)
// 	s := string(bodyText)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var response JiraResponse
// 	json.Unmarshal([]byte(s), &response)
// 	return response.Issues
// }

type Res struct {
	issueID  string
	blockers []Issue
}

// func recurseFetch(ch chan Res, issueID string, userEmail string, userToken string) {
// 	issues := fetchBlockingIssues(issueID, userEmail, userToken)
// 	ch <- Res{issueID: issueID, blockers: issues}
// }

// func Crawl(issue Issue, depth int, ch chan Res, errs chan error, issueID string, userEmail string, userToken string) {
// 	issues := fetchBlockingIssues(issueID, userEmail, userToken)

// 	newUrls := 0
// 	if depth > 1 {
// 		for _, u := range issues {
// 			newUrls++
// 			go Crawl(u, depth-1, ch, errs, issueID, userEmail, userToken)

// 		}
// 	}

// 	// Send the result along with number of urls to be fetched
// 	ch <- Res{issueID: issueID, blockers: issues}

// 	return
// }
