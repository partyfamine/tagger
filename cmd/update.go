package cmd

import (
	"github.com/partyfamine/tagger/cmd/update"
	"github.com/spf13/cobra"
)

var Update = &cobra.Command{
	Use:   "update",
	Short: "updates id3 data",
	Long:  "updates id3 data",
}

func init() {
	Update.AddCommand(update.TrackData)
	Update.AddCommand(update.FileNames)
}
