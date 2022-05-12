package generate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/partyfamine/tagger/data"
	"github.com/spf13/cobra"
)

var Dirs = &cobra.Command{
	Use:   "dirs",
	Short: "generates directory structure",
	Long:  "generates directory structure",
	Run:   dirs,
}

var (
	dataFile string
	rootDir  string
)

func init() {
	Dirs.Flags().StringVarP(&dataFile, "file", "f", "", "file containing album data in json format")
	Dirs.Flags().StringVarP(&rootDir, "directory", "d", "", "root directory to generate directories")
	Dirs.MarkFlagRequired("file")
}

func dirs(cmd *cobra.Command, args []string) {
	fileBytes, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	albums := make(map[string]data.Album)
	if err := json.Unmarshal(fileBytes, &albums); err != nil {
		log.Fatal(err)
	}

	if rootDir == "" {
		rootDir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, album := range albums {
		artistDir := filepath.Join(rootDir, album.Artist)
		_, err := os.Stat(artistDir)
		if os.IsNotExist(err) {
			if err := os.Mkdir(artistDir, os.ModeDir); err != nil {
				log.Fatal(err)
			}
		} else if err != nil {
			log.Fatal(err)
		}
		albumDir := filepath.Join(rootDir, album.Artist, album.Name)
		if err := os.Mkdir(albumDir, os.ModeDir); err != nil {
			log.Fatal(err)
		}
	}
}
