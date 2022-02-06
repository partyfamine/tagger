package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/bogem/id3v2"
	"github.com/partyfamine/tagger/cmd/data"
	"github.com/spf13/cobra"
)

var TrackData = &cobra.Command{
	Use:   "track-names",
	Short: "updates track names",
	Long:  "updates track names using json data scraped from rym",
	Run:   trackData,
}

var (
	dataFile string
	dirName  string
)

func init() {
	TrackData.Flags().StringVarP(&dataFile, "file", "f", "", "file containing track data in json format")
	TrackData.Flags().StringVarP(&dirName, "directory", "d", "", "directory containing music files")
	TrackData.MarkFlagRequired("file")
	TrackData.MarkFlagRequired("directory")
}

func trackData(cmd *cobra.Command, args []string) {
	fileBytes, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	albums := make(map[string]data.Album)
	if err := json.Unmarshal(fileBytes, &albums); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		filePath := filepath.Join(dirName, file.Name())
		filepath.Join(dirName, file.Name())
		fmt.Printf("searching %s\n", filePath)
		subFiles, err := ioutil.ReadDir(filePath)
		if err != nil {
			log.Fatal(err)
		}

		for _, subFile := range subFiles {
			if !subFile.IsDir() {
				continue
			}
			album, ok := albums[subFile.Name()]
			if !ok {
				continue
			}

			subFilePath := filepath.Join(dirName, file.Name(), subFile.Name())
			fmt.Printf("searching %s\n", subFilePath)
			subFiles, err := ioutil.ReadDir(subFilePath)
			if err != nil {
				log.Fatal(err)
			}

			for i, trackFile := range subFiles {
				if trackFile.IsDir() {
					continue
				}

				trackFilePath := filepath.Join(dirName, file.Name(), subFile.Name(), trackFile.Name())
				fmt.Printf("setting tags for %s\n", trackFilePath)
				songFile, err := id3v2.Open(trackFilePath, id3v2.Options{Parse: true})
				if err != nil {
					log.Fatal("error opening file", err)
				}

				songFile.SetAlbum(album.Name)
				songFile.SetArtist(album.Artist)
				songFile.AddTextFrame(songFile.CommonID("Band/Orchestra/Accompaniment"), songFile.DefaultEncoding(), album.Artist)
				songFile.SetGenre(album.Genre)
				songFile.SetYear(album.Year)
				songFile.SetTitle(album.TrackNames[i])
				songFile.AddTextFrame(songFile.CommonID("Track number/Position in set"), songFile.DefaultEncoding(), fmt.Sprintf("%d/%d", i+1, len(album.TrackNames)))

				if err = songFile.Save(); err != nil {
					fmt.Printf("Error while saving a tag: %s\n", err.Error())
					break
				}
			}
		}
	}
}
