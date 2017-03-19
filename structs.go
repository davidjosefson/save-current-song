package main

type Conf struct {
	LASTFM_API_KEY        string `json:"LASTFM_API_KEY"`
	LASTFM_USERNAME       string `json:"LASTFM_USERNAME"`
	SPOTIFY_REFRESH_TOKEN string `json:"SPOTIFY_REFRESH_TOKEN"`
	SPOTIFY_CLIENT_ID     string `json:"SPOTIFY_CLIENT_ID"`
	SPOTIFY_CLIENT_SECRET string `json:"SPOTIFY_CLIENT_SECRET"`
	SPOTIFY_USERNAME      string `json:"SPOTIFY_USERNAME"`
	SPOTIFY_PLAYLIST_ID   string `json:"SPOTIFY_PLAYLIST_ID"`
}

type CurrentSong struct {
	RecentTracks *RecentTracks `json:"recenttracks"`
}

type RecentTracks struct {
	Tracks []Track `json:"track"`
}

type Track struct {
	Artist *Artist `json:"artist"`
	Track  string  `json:"name"`
}

type Artist struct {
	Name string `json:"#text"`
}

type SpotifySong struct {
	SpotifyTracks *SpotifyTracks `json:"tracks"`
}

type SpotifyTracks struct {
	SpotifyItems []SpotifyItem `json:"items"`
}

type SpotifyItem struct {
	Id      string          `json:"id"`
	Uri     string          `json:"uri"`
	Name    string          `json:"name"`
	Artists []SpotifyArtist `json:"artists"`
}

type SpotifyArtist struct {
	Name string `json:"name"`
}

type SpotifyToken struct {
	Token string `json:"access_token"`
}
