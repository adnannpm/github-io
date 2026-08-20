/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github-io/internal/api"
)

var (
	name string
)

var findUserCmd = &cobra.Command{
	Use:   "find-user",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("Need username github")
			return
		}

		api.Username(name)
	},
}

func init() {
	rootCmd.AddCommand(findUserCmd)

	findUserCmd.Flags().StringVar(&name, "name", "", "Need username github...")
}
