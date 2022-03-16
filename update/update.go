package update

import (
	"github.com/spf13/cobra"
)

var Update = &cobra.Command{
	Use:   "update",
	Short: "updates id3 data",
	Long:  "updates id3 data",
}

func init() {
	Update.AddCommand(TrackData)
	Update.AddCommand(FileNames)
}
