# Jira Blockers

When working through a complex project in Jira the question I always ask is "what are the blockers, and what are the blocker's blockers." There are some plugins that answer this question, but maybe not in the way I want.

This project presents the blockers and their child blockers in a tree-like format. This doesn't really catch the possible complexity of a set of Jira issues, which are more of a graph. However, for my uses it's a quick way to get a handle on what needs to be done first.

In order to use it you'll need to generate a user token for yourself.
```
go run . -user=username@example.com -token=asdf1234 -issue=ISS-123
```

You may also pass a `-depth` flag if you'd like to shorten or lengthen the tree depth from the default `3`.

Caveat emptor, there's very little error handling and the JQL query is fragile to say the least.

