package cmd

import (
	"fmt"
	"github.com/conijnio/aws-iam-user/pkg/core"
	"github.com/spf13/cobra"
	"os"
)

var (
	Debug  bool
	osExit = os.Exit
)

var rootCmd = &cobra.Command{
	Use:     "aws-iam-user",
	Short:   "aws-iam-user - Sample cli tool implementation",
	PreRun:  preRun,
	Example: "aws-iam-user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.MainRoutine()
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "verbose logging")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		osExit(1)
	}
}

func SetVersion(version string) {
	rootCmd.Version = version
}

func preRun(cmd *cobra.Command, args []string) {
	toggleDebug()
}
