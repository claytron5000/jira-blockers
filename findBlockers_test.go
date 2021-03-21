package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/xlab/treeprint"
)

// func TestGenerateQueryString(t *testing.T) {
// 	ans := generateQueryString("SYM-123")
// 	expect := "project%20=%20SYM%20AND%20issueBlocks%20=SYM-123"
// 	if ans != expect {
// 		t.Errorf("generateQueryString(\"SYM-123\") = %s; want %s", ans, expect)
// 	}
// }

type fakeUnEvenFetcher struct {
}

func (f fakeUnEvenFetcher) fetchIssues(issueID string) []Issue {
	var array []Issue
	// initial return
	if issueID == "a" {
		b := Issue{Key: "b"}
		array = append(array, b)
		c := Issue{Key: "c"}
		array = append(array, c)
	}
	// branch b quickly returns three results
	if issueID == "b" {
		d := Issue{Key: "d"}
		array = append(array, d)
		e := Issue{Key: "e"}
		array = append(array, e)

	}
	// branch c returns two returns after 3 seconds
	if issueID == "c" {
		time.Sleep(3 * time.Second)
		f := Issue{Key: "f"}
		array = append(array, f)
		g := Issue{Key: "g"}
		array = append(array, g)
		h := Issue{Key: "h"}
		array = append(array, h)
	}

	return array
}

func TestUnEvenFetch(t *testing.T) {
	fetcher := fakeUnEvenFetcher{}
	tree := treeprint.New()
	var ch = make(chan int)

	go RecurseTreeFetching(fetcher, tree, "a", ch, 3)
	childs := 1
	for childs > 0 {
		childs += <-ch
	}

	fmt.Println(tree.String())
}
