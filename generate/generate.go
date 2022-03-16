package generate

import "github.com/spf13/cobra"

var Generate = &cobra.Command{
	Use:   "generate",
	Short: "generate files",
	Long:  "generate files",
}

func init() {
	Generate.AddCommand(Dirs)
}
