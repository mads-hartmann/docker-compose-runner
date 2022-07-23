package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-compose-runner",
	Short: "A CLI for managing Docker Compose applications",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WIP")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
