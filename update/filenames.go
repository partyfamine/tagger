package update

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/spf13/cobra"
)

var FileNames = &cobra.Command{
	Use:   "file-names",
	Short: "updates file names",
	Long:  "updates file names using existing id3 data",
	Run:   fileNames,
}

var renameDir string

func init() {
	FileNames.Flags().StringVarP(&renameDir, "directory", "d", "", "directory containing music files to rename")
}

func fileNames(cmd *cobra.Command, args []string) {
	if renameDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		renameDir = wd
	}
	songDirs := findSongDirs(renameDir)

	for _, dir := range songDirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				fullFilePath := filepath.Join(dir, file.Name())
				tag, err := id3v2.Open(fullFilePath, id3v2.Options{Parse: true})
				if err != nil {
					log.Fatal(err)
				}

				title := tag.Title()
				title = strings.ReplaceAll(title, ":", "-")
				fullTrackNum := tag.GetTextFrame(tag.CommonID("Track number/Position in set")).Text

				splitTrackNums := strings.Split(fullTrackNum, "/")
				trackNum := mustFormatTrackStr(splitTrackNums[0])

				newFileName := filepath.Join(dir, trackNum+" "+title+".mp3")
				fmt.Printf("renaming %s\n      to %s\n", fullFilePath, newFileName)

				if err := tag.Close(); err != nil {
					fmt.Printf("err: %s\n", err.Error())
				}
				if err := os.Rename(fullFilePath, newFileName); err != nil {
					fmt.Printf("err: %s\n", err.Error())
				}
			}
		}
	}
}

func mustFormatTrackStr(s string) string {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%02d", i)
}

func findSongDirs(startDir string) []string {
	var songDirs []string
	foundSongDirs := make(map[string]struct{})

	filepath.Walk(filepath.Clean(startDir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			parentDir := filepath.Dir(path)
			if _, ok := foundSongDirs[parentDir]; !ok {
				foundSongDirs[parentDir] = struct{}{}
				songDirs = append(songDirs, parentDir)
			}
		}
		return nil
	})

	return songDirs
}
