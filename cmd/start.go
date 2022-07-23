package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var watch bool

func init() {
	startCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Watch for changes to docker-compose file and restart when detected")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Docker Composed application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("start watch=%t", watch)
	},
}
