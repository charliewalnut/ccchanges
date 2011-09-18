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
var dateLineRegexp = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})\s+(.+?)\s+<([^<>]+)>$`)

func parseCommitter(line string) Person {
    submatches := dateLineRegexp.FindStringSubmatch(line)
    if submatches == nil {
        return Person{}
    }
    return Person{submatches[2], submatches[3]}
}

func ParseEntry(entry string) Change {
    c := Change{}
    lines := strings.Split(entry, "\n")
    c.committer = parseCommitter(lines[0])
    return c
}
