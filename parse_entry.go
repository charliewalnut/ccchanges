package ccchanges

import (
    "exp/regexp" // Using perl character classes \d and \s
    "strings"
)

type Change struct {
    author, reviewer, committer string
    paths []string
    rollout bool
}

// Adapted from webkitpy/common/checkout/changelog.py
var dateLineRE = regexp.MustCompile(`^20\d{2}-\d{2}-\d{2}\s+(.+?)\s+<[^<>]+>$`)
var reviewerRE = regexp.MustCompile(`Reviewed by (.*?)[\.]`)

var rolloutRE = regexp.MustCompile(`Unreviewed, rolling out r(\d+)[\.]`)
var pathRE = regexp.MustCompile(`\* ([\w/\.]+):`)

func parseCommitter(line string) string {
    submatches := dateLineRE.FindStringSubmatch(line)
    if submatches == nil {
        return ""
    }
    return submatches[1]
}

func parseReviewer(lines []string) string {
    for _, l := range lines {
        submatches := reviewerRE.FindStringSubmatch(l)
        if submatches != nil {
            return submatches[1]
        }
    }
    return ""
}

func parseRollout(lines []string) bool {
    for _, l := range lines {
        submatches := rolloutRE.FindStringSubmatch(l)
        if submatches != nil {
            return true
        }
    }
    return false
}

func parsePaths(lines []string) []string {
    paths := make([]string, 0)
    for _, l := range lines {
        submatches := pathRE.FindStringSubmatch(l)
        if submatches != nil {
            paths = append(paths, submatches[1])
        }
    }
    return paths
}

func ParseEntry(entry string) Change {
    c := Change{}
    lines := strings.Split(entry, "\n")
    c.committer = parseCommitter(lines[0])
    c.author = c.committer
    c.reviewer = parseReviewer(lines[1:])
    c.rollout = parseRollout(lines[1:])
    c.paths = parsePaths(lines[1:])
    return c
}
