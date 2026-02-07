package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "eksdoctor",
	Short: "EKS wiring snapshot, diff and diagnostics tool",
}

func Execute() error {
	return rootCmd.Execute()
}
