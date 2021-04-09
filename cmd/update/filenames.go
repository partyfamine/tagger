package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FileNames = &cobra.Command{
	Use:   "file-names",
	Short: "updates file names",
	Long:  "updates file names using existing id3 data",
	Run:   fileNames,
}

func fileNames(cmd *cobra.Command, args []string) {
	fmt.Println("unimplemented")
}
