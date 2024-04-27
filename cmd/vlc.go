package cmd

import "github.com/spf13/cobra"

var vlc = &cobra.Command{
	Short: "Vlc algorithm",
	Run:   vlcPack,
}

func vlcPack(cmd *cobra.Command, args []string) {
	print("huy")
}

func init() {
	pack.AddCommand(vlc)
}
