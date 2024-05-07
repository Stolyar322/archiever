package cmd

import "github.com/spf13/cobra"

var pack = &cobra.Command{
	Use:   "pack",
	Short: `Pack file`,
}

func init() {
	rootCmd.AddCommand(pack)
}
