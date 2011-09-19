package main

import (
    "os"
)

type LineGetter interface {
    GetLine() (string, os.Error, bool)
}

func parseLog(lineGetter LineGetter) []Change {
    changes := make([]Change, 0)
    line, err, dateline := lineGetter.GetLine()
    for err == nil {
        entry := make([]string, 0)
        entry = append(entry, line)
        for line, err, dateline = lineGetter.GetLine();
            err == nil && !dateline;
            line, err, dateline = lineGetter.GetLine() {
            entry = append(entry, line)
        }
        changes = append(changes, parseEntry(entry))
    }
    return changes
}
