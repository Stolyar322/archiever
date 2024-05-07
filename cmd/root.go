package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: `Root command`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ErrorHandle(err)
	}
}

func ErrorHandle(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
