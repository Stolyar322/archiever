package cmd

import "github.com/spf13/cobra"

var pack = &cobra.Command{
	Short: `Pack file`,
}

func init() {
	cmd.AddCommand(pack)
}
