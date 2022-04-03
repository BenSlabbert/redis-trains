package main

import "fmt"

// GitCommit is set during compilation
var GitCommit string

func main() {
	fmt.Printf("GitCommit: %s", GitCommit)
	// listen to train stream and send notifications of train errors/changes
}
