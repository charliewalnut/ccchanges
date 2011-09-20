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
var startdate = flag.String("startdate", "", "start date of time limit")
var enddate = flag.String("enddate", "", "end date of time limit")

type Filter struct {
    path, reviewer, committer *regexp.Regexp
    startdate, enddate string
    reviewed bool
}

func (f *Filter) Matches(change *Change) bool {
    pathMatch := false
    for i := 0; !pathMatch && i < len(change.paths); i++ {
        pathMatch = f.path.MatchString(change.paths[i])
    }
    if !pathMatch {
        return false
    }
    if f.reviewed && change.reviewer == "" {
        return false
    }
    return f.committer.MatchString(change.committer) && f.reviewer.MatchString(change.reviewer)
}

func main() {
    flag.Parse()

    lg := &bufIOReaderLineGetter{bufio.NewReader(os.Stdin)}
    filter := &Filter{regexp.MustCompile(*path),
                      regexp.MustCompile(*reviewer),
                      regexp.MustCompile(*reviewer),
                      *startdate,
                      *enddate,
                      *reviewed}
    changes := parseLog(lg, filter)
    commits := make(map[string] int)
    reviews := make(map[string] int)
    numCommits := 0
    numReviewedCommits := 0
    for _, change := range changes {
        commits[change.committer]++
        numCommits++
        if change.reviewer != "" {
            reviews[change.reviewer]++
            numReviewedCommits++
        }
    }
    for committer, count := range commits {
        fmt.Printf("commits %d %s %.1f%% of total, %.1f%% of reviewed\n", count, committer, 100.0*float64(count)/float64(numCommits), 100.0*float64(count)/float64(numReviewedCommits))
    }
    for reviewer, count := range reviews {
        fmt.Printf("reviews %d %s %.1f%%\n", count, reviewer, 100.0*float64(count)/float64(numReviewedCommits))
    }
}

