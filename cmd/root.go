package cmd

import (
	"fmt"
	"github.com/conijnio/aws-iam-user/pkg/adapters"
	"github.com/conijnio/aws-iam-user/pkg/models"
	"github.com/spf13/cobra"
	"os"
)

var (
	Debug   bool
	osExit  = os.Exit
	profile string
	region  = "eu-west-1" // TODO: Read the default region from the configuration
	adapter adapters.IUserAdapter
)

var rootCmd = &cobra.Command{
	Use:     "aws-iam-user",
	Short:   "aws-iam-user - Sample cli tool implementation",
	PreRun:  preRun,
	Example: "aws-iam-user",
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := adapter.LoadUser()

		if err != nil {
			return err
		}

		renderOutput(user)
		return nil
	},
}

func renderOutput(user *models.User) {
	fmt.Printf("Found the following IAM User using the '%s' profile\n\n", profile)
	fmt.Printf("Account: %s\n", user.Account)
	fmt.Printf("UserId: %s\n", user.UserId)
	fmt.Printf("UserName: %s\n", user.Name)
	fmt.Printf("Type: %s\n", user.Type)
	fmt.Printf("Arn: %s\n", user.Arn)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "verbose logging")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "default", "The AWS profile used")
	rootCmd.PersistentFlags().StringVar(&region, "region", region, "The AWS profile used")
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
	adapter = adapters.LoadUserAdapter(region, profile)
}
