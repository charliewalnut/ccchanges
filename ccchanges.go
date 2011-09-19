package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "exp/regexp"
)

type bufIOReaderLineGetter struct {
    reader *bufio.Reader
}

func (b *bufIOReaderLineGetter) GetLine() (string, os.Error, bool) {
    line, err := b.reader.ReadString('\n')
    if len(line) > 1 {
        line = line[:len(line)-1]
    }
    return line, err, err == nil && dateLineRE.MatchString(line)
}

var path = flag.String("path", "", "path regexp filter to use")
var committer = flag.String("committer", "", "committer regexp filter to use")
var reviewer = flag.String("reviewer", "", "reviewer regexp filter to use")
var reviewed = flag.Bool("reviewed", false, "reviewed patches only")

func ezMatch(pattern, s string) bool {
    m, _ := regexp.MatchString(pattern, s)
    return m
}

func matchesFilter(change Change) bool {
    pathMatch := false
    for i := 0; !pathMatch && i < len(change.paths); i++ {
        pathMatch = ezMatch(*path, change.paths[i])
    }
    if !pathMatch {
        return false
    }
    if *reviewed && change.reviewer == "" {
        return false
    }
    return ezMatch(*committer, change.committer) && ezMatch(*reviewer, change.reviewer)
}

func main() {
    flag.Parse()
    lg := &bufIOReaderLineGetter{bufio.NewReader(os.Stdin)}
    changes := parseLog(lg)
    commits := make(map[string] int)
    reviews := make(map[string] int)
    for _, change := range changes {
        if matchesFilter(change) {
            commits[change.committer]++
            if change.reviewer != "" {
                reviews[change.reviewer]++
            }
        }
    }
    for committer, count := range commits {
        fmt.Printf("commits %d %s\n", count, committer)
    }
    for reviewer, count := range reviews {
        fmt.Printf("reviews %d %s\n", count, reviewer)
    }
}

