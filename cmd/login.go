/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github-io/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	token string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if token == "" {
			fmt.Println("Need token github, please input your token github!")
			return
		}

		config.InsertToken(token)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVar(&token, "token", "", "Need token github!")
}
