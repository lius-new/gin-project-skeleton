package scripts

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/lius-new/ln-blog/internal/app"
	_ "github.com/lius-new/ln-blog/internal/init"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start applications.",

	Run: func(cmd *cobra.Command, args []string) {
		app.Aapplication.StartWithSupportStop(":8080", time.Second*3)
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true // hide Completion command
	rootCmd.AddCommand(startCmd)
}
