package cmd

import (
	"github-io/internal/defaults"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Initializing configuration...")
		if err := createDir(configPath, "Main"); err != nil {
			log.Fatalln("Failed to inititalize main directory: ", err)
		}

		if err := writeFile(filepath.Join(configPath, "main.yml"), "main.yml", defaults.MainFile); err != nil {
			log.Fatalln("Failed to initialize main file: %w", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createDir(path, name string) error {
	if info, err := os.Stat(path); err == nil {
		if !info.IsDir() {
			return fmt.Errorf("%s path exists but is not a directory: %s", name, path)
		}
		fmt.Printf("%s directory already exists", name)
		return nil
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create %s directory %s: %w", name, path, err)
	}

	fmt.Printf("%s directory created", name)

	return nil
}

func writeFile(path, name, content string) error {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("%s already exists", name)

		return nil
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create %s %s: %w", name, path, err)
	}

	fmt.Printf("%s created", name)

	return nil
}
