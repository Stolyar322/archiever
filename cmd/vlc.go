package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var vlc = &cobra.Command{
	Use:   "vlc",
	Short: "Vlc algorithm",
	Run:   vlcPack,
}

const packExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func vlcPack(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		ErrorHandle(ErrEmptyPath)
	}

	filePath := args[0]

	fileConnection, err := os.Open(filePath)
	if err != nil {
		ErrorHandle(err)
	}
	defer func(fileConnection *os.File) {
		err := fileConnection.Close()
		if err != nil {
			ErrorHandle(err)
		}
	}(fileConnection)

	packFileName := setPackedFileName(filePath)
	_, err = fmt.Fprintln(os.Stderr, packFileName)
	if err != nil {
		ErrorHandle(err)
	}
}

func setPackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packExtension
}

func init() {
	pack.AddCommand(vlc)
}
