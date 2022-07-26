package cmd

import (
	"github.com/mads-hartmann/docker-compose-runner/pkg/repository"
	"github.com/spf13/cobra"
)

var watch bool

func init() {
	startCmd.Flags().BoolVarP(&watch, "watch", "w", false, "Watch for changes to docker-compose file and restart when detected")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start <git-url>",
	Short: "Start a Docker Composed application",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := repository.Repository{
			GitUrl: args[0],
			Watch:  watch,
		}
		repo.Clone()
	},
}
