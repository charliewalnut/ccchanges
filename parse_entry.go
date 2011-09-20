package main

import (
    "exp/regexp" // Using perl character classes \d and \s
    "time"
)

type Change struct {
    author, reviewer, committer string
    paths []string
    rollout bool
    date *time.Time
}

// Adapted from webkitpy/common/checkout/changelog.py
var dateLineRE = regexp.MustCompile(`^(20\d{2}-\d{2}-\d{2})\s+(.+?)\s+<[^<>]+>$`)
var reviewerRE = regexp.MustCompile(`Reviewed by (.*?)[\.]`)

var rolloutRE = regexp.MustCompile(`Unreviewed, rolling out r(\d+)[\.]`)
var pathRE = regexp.MustCompile(`\* ([\w/\.]+):`)

func parseDateAndCommitter(line string) (*time.Time, string) {
    submatches := dateLineRE.FindStringSubmatch(line)
    if submatches == nil {
        return nil, ""
    }
    t, _ := time.Parse("2006-01-02", submatches[1])
    return t, submatches[2]
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

func parseEntry(entry []string) Change {
    c := Change{}
    c.date, c.committer = parseDateAndCommitter(entry[0])
    c.author = c.committer
    c.reviewer = parseReviewer(entry[1:])
    c.rollout = parseRollout(entry[1:])
    c.paths = parsePaths(entry[1:])
    return c
}
