package main

import (
	"github.com/partyfamine/tagger/generate"
	"github.com/partyfamine/tagger/list"
	"github.com/partyfamine/tagger/update"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "tagger",
	Short: "tools for managing id3 tags",
	Long:  "tools for managing id3 tags",
}

func init() {
	root.AddCommand(update.Update)
	root.AddCommand(generate.Generate)
	root.AddCommand(list.List)
}

func main() {
	root.Execute()
}
