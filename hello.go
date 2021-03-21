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

	var ch = make(chan int)

	go recurseTreeFetching(fetcher, tree, *issueID, ch, *depth)
	childs := 1
	for childs > 0 {
		childs += <-ch
	}

	fmt.Println(tree.String())

}

func recurseTreeFetching(fetcher IFetcher, tree treeprint.Tree, issueID string, ch chan int, depth int) {
	issues := fetcher.fetchIssues(issueID)
	if len(issues) == 0 {
		// delete this child
		ch <- -1
		return
	}
	depth--
	if depth == 0 {
		// delete this child
		ch <- -1
		return
	}
	// add the number of child issues, minus the current one
	ch <- len(issues) - 1

	for i := 0; i < len(issues); i++ {
		currIssueID := issues[i].Key
		currBranch := tree.AddBranch(currIssueID)
		go recurseTreeFetching(fetcher, currBranch, currIssueID, ch, depth)
	}
}

type Issue struct {
	Key string
}
type JiraResponse struct {
	Issues []Issue
}
