package main

import "github.com/conijnio/aws-iam-user/cmd"

var (
	version = "dev"
	// commit  = "none"
	// date    = "unknown"
)

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
