package main

import "github.com/conijnio/golang-template/cmd"

var (
	version = "dev"
	// commit  = "none"
	// date    = "unknown"
)

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
