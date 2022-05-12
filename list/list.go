package list

import (
	"fmt"
	"log"

	"github.com/bogem/id3v2/v2"
	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "lists tags for a song file",
	Long:  "lists tags for a song file",
	Run:   list,
}

var fileName string

func init() {
	List.Flags().StringVarP(&fileName, "file", "f", "", "song file")
	List.MarkFlagRequired("file")
}

func list(*cobra.Command, []string) {
	song, err := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("error opening file", err)
	}
	defer song.Close()

	for id, f := range song.AllFrames() {
		for _, frame := range f {
			if txtFrame, ok := frame.(id3v2.TextFrame); ok {
				fmt.Printf("%s: %s\n", id, txtFrame.Text)
				continue
			}
			if cmtFrame, ok := frame.(id3v2.CommentFrame); ok {
				fmt.Printf("%s: %s\n", id, cmtFrame.Text)
			}
		}
	}
}
