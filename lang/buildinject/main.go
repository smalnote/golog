package main

import "fmt"

var (
	commitHash string
	commitDate string
)

type commitInfo struct {
	CommitHash string `json:"commit"`
	CommitDate string `json:"date"`
}

// build with go build -ldflags "-X main.commitHash=$(git rev-parse --short HEAD) -X main.commitDate=$(git log -1 --format=%ct)"
func main() {
	info := commitInfo{
		CommitHash: commitHash,
		CommitDate: commitDate,
	}
	fmt.Printf("%+v", info)
}
