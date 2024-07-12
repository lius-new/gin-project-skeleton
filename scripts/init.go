package scripts

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ln-scripts",
	Short: "ln-scripts is ln-blog manager control version.",
	Long:  "You can use ln-scripts to run various commands to manage applications. ",
}

func Execute() {
	rootCmd.Execute()
}
