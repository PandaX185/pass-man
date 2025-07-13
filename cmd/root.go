package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "passman",
	Short: "Password Manager CLI",
	Long:  `This is a command line interface for managing passwords securely.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
