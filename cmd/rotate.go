package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rotateCmd = &cobra.Command{
	Use:     "rotate",
	Short:   "Rotate the CLI credentials from the given profile.",
	PreRun:  preRun,
	Example: "aws-iam-user",
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := adapter.LoadUser()

		if err != nil {
			return err
		}

		if err = adapter.RotateCredentials(user); err == nil {
			fmt.Println("The credentials has been rotated!")
		}

		return err
	},
}
