package ccchanges

type Person struct {
    name, email string
}

type Change struct {
    author, reviewer, committer Person
    rollout bool
}

func ParseEntry(entry string) Change {
    c := Change{}
    return c
}
