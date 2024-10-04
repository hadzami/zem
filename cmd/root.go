package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "zem",
	Short: "zem is a CLI tool for engineering productivity",
	Long:  "zem is a CLI tool for engineering productivity for managing your aliases in the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // Show help if no command is provided
	},
}

var addCmd = &cobra.Command{
	Use:   "add [alias] [command]",
	Short: "Add a new alias",
	Long:  "Add a new alias to your .zshrc file.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		command := args[1]

		err := AddAlias(alias, command)
		if err != nil {
			fmt.Printf("Error adding alias: %s\n", err)
		} else {
			fmt.Printf("Alias '%s' added for command '%s'.\n", alias, command)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
