package main

import (
    "strings"
)

func parseLog(log string) []Change {
    lines := strings.Split(log, "\n")
    i := 0
    changes := make([]Change, 0)
    for i < len(lines) {
        entry := make([]string, 0)
        entry = append(entry, lines[i])
        for i++; i < len(lines) && !dateLineRE.MatchString(lines[i]); i++ {
            entry = append(entry, lines[i])
        }
        changes = append(changes, parseEntry(entry))
    }
    return changes
}
