package data

type Album struct {
	Name       string   `json:"name"`
	Artist     string   `json:"artist"`
	Year       string   `json:"year"`
	Genre      string   `json:"genre"`
	TrackNames []string `json:"trackNames"`
}
