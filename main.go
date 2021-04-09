package main

import (
	"github.com/partyfamine/tagger/cmd"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "tagger",
	Short: "tools for managing id3 tags",
	Long:  "tools for managing id3 tags",
}

func init() {
	root.AddCommand(cmd.Update)
}

func main() {
	root.Execute()
}
