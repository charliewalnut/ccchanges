package main

import (
    "bufio"
    "exp/regexp"
    "flag"
    "fmt"
    "os"
    "time"
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
var mincommits = flag.Int("mincommits", 0, "minimum commits (-1 to filter all commits)")
var minreviews = flag.Int("minreviews", 0, "minimum reviews (-1 to filter all reviews)")
var startdate = flag.String("startdate", "", "start date of time limit")
var enddate = flag.String("enddate", "", "end date of time limit")
var reviewed = flag.Bool("reviewed", false, "reviewed patches only")

type Filter struct {
    path, reviewer, committer *regexp.Regexp
    startdate, enddate *time.Time
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
    if f.startdate != nil {
        if change.date == nil || change.date.Seconds() < f.startdate.Seconds() {
            return false
        }
    }
    if f.enddate != nil {
        if change.date == nil || change.date.Seconds() < f.enddate.Seconds() {
            return false
        }
    }
    return f.committer.MatchString(change.committer) && f.reviewer.MatchString(change.reviewer)
}

func main() {
    flag.Parse()

    lg := &bufIOReaderLineGetter{bufio.NewReader(os.Stdin)}
    start, _ := time.Parse("2006-01-02", *startdate)
    end, _ := time.Parse("2006-01-02", *enddate)
    filter := &Filter{regexp.MustCompile(*path),
                      regexp.MustCompile(*reviewer),
                      regexp.MustCompile(*reviewer),
                      start, end,
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
    if *mincommits != -1 {
        for committer, count := range commits {
            if count > *mincommits {
                fmt.Printf("commits %d %s %.1f%% of total, %.1f%% of reviewed\n", count, committer, 100.0*float64(count)/float64(numCommits), 100.0*float64(count)/float64(numReviewedCommits))
            }
        }
    }
    if *minreviews != -1 {
        for reviewer, count := range reviews {
            if count > *minreviews {
                fmt.Printf("reviews %d %s %.1f%%\n", count, reviewer, 100.0*float64(count)/float64(numReviewedCommits))
            }
        }
    }
}

