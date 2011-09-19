package ccchanges

import (
    "exp/regexp" // Using perl character classes \d and \s
    "strings"
)

type Person struct {
    name, email string
}

type Change struct {
    author, reviewer, committer Person
    rollout bool
}

// Adapted from webkitpy/common/checkout/changelog.py
var dateLineRE = regexp.MustCompile(`^20\d{2}-\d{2}-\d{2}\s+(.+?)\s+<([^<>]+)>$`)
var reviewerRE = regexp.MustCompile(`Reviewed by (.*?)[\.]`)

func parseCommitter(line string) Person {
    submatches := dateLineRE.FindStringSubmatch(line)
    if submatches == nil {
        return Person{}
    }
    return Person{submatches[1], submatches[2]}
}

func parseReviewer(lines []string) Person {
    for i := 0; i < len(lines); i++ {
        submatches := reviewerRE.FindStringSubmatch(lines[i])
        if submatches != nil {
            return Person{submatches[1], ""}
        }
    }
    return Person{}
}

func ParseEntry(entry string) Change {
    c := Change{}
    lines := strings.Split(entry, "\n")
    c.committer = parseCommitter(lines[0])
    c.author = c.committer
    c.reviewer = parseReviewer(lines[1:])
    return c
}
