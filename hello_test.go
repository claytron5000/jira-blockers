package main

import "testing"

func TestGenerateQueryString(t *testing.T) {
	ans := generateQueryString("SYM-123")
	expect := "project%20=%20SYM%20AND%20issueBlocks%20=SYM-123"
	if ans != expect {
		t.Errorf("generateQueryString(\"SYM-123\") = %s; want %s", ans, expect)
	}
}
