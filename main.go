package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"save-current-song/notify"
	"strings"
)

var conf Conf

func main() {
	var err error

	var dirPath string
	var currentSong CurrentSong
	var token string

	// INITIALIZATION
	dirPath, err = getDirPath()
	handleError(err, "trying to create an os.Executable")
	initializeLogging(dirPath)
	readConf(dirPath)

	// REFRESH SPOTIFY TOKEN
	log.Println("Refreshing Spotify token..")
	token, err = refreshSpotifyToken()
	handleError(err, "refreshing Spotify token")

	// GET CURRENT SONG FROM SPOTIFY
	log.Println("Fetching currently playing song from Spotify..")
	currentSong, err = getCurrentSong(token)
	handleError(err, "fetching current song from Spotify")

	// ADD SONG TO PLAYLIST
	log.Println("Adding song to playlist..")
	err = addSongToPlaylist(token, currentSong.SpotifyItem.Uri)
	handleError(err, "adding song to playlist")
	log.Println("Saved: " + currentSong.SpotifyItem.Artists[0].Name + " - " + currentSong.SpotifyItem.Name)

	// SUCCESS
	log.Println("Succeeded to add song to playlist!")

	// SEND NOTIFICATION
	addedSong := currentSong.SpotifyItem.Artists[0].Name + " - " + currentSong.SpotifyItem.Name
	notify.Notify(addedSong)
}

func getDirPath() (string, error) {
	ex, err := os.Executable()
	return path.Dir(ex), err
}

func initializeLogging(dirPath string) {
	file, err := os.OpenFile(dirPath+"/log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err == nil {
		log.SetOutput(file)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	log.Println("-------------------------")
}

func readConf(dirPath string) {
	file, err := os.Open(dirPath + "/conf.json")
	handleError(err, "trying to open conf.json")

	decoder := json.NewDecoder(file)
	conf = Conf{}
	err = decoder.Decode(&conf)
	handleError(err, "decode conf.json")
}

func getCurrentSong(token string) (CurrentSong, error) {
	var err error
	var currentSong CurrentSong

	url := "https://api.spotify.com/v1/me/player/currently-playing"

	// CREATING REQUEST
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	handleError(err, "create request to fetch currently playing song")

	// SENDING REQUEST
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	handleError(err, "sending request to add song to playlist")

	// READING RESPONSE BODY
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	handleError(err, "trying to read body after getting current song")

	// UNMARSHALLING BODY TO JSON
	err = json.Unmarshal(body, &currentSong)
	handleError(err, "unmarshalling body after getting current song")

	// SUCCESS
	log.Println("Succeeded to get current song!")
	log.Println("Found: " + currentSong.SpotifyItem.Artists[0].Name + " - " + currentSong.SpotifyItem.Name)

	return currentSong, err
}

func refreshSpotifyToken() (string, error) {
	var err error
	var spotifyToken SpotifyToken

	// CREATING REQUEST
	spotifyUrl := "https://accounts.spotify.com/api/token"
	formValues := url.Values{}
	formValues.Set("grant_type", "refresh_token")
	formValues.Set("refresh_token", conf.SPOTIFY_REFRESH_TOKEN)
	formValues.Set("client_id", conf.SPOTIFY_CLIENT_ID)
	formValues.Set("client_secret", conf.SPOTIFY_CLIENT_SECRET)

	// SENDING REQUEST AND RECEIVING RESPONSE
	resp, err := http.PostForm(spotifyUrl, formValues)
	handleError(err, "POSTing to refresh Spotify token")

	// READING RESPONSE BODY
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	handleError(err, "trying to read body after POSTing to refresh Spotify token")

	// UNMARSHALLING BODY TO JSON
	err = json.Unmarshal(body, &spotifyToken)
	handleError(err, "unmarshalling body after refreshing Spotify token")

	// SUCCESS
	log.Println("Succeeded to refresh token!")
	log.Println("TOKEN: " + spotifyToken.Token)

	return spotifyToken.Token, err
}

func replaceSpacesWithPlus(textWithSpaces string) string {
	return strings.Replace(textWithSpaces, " ", "+", -1)
}

func addSongToPlaylist(token string, currentSongUri string) error {
	var err error

	spotifyUrl := "https://api.spotify.com/v1/users/" + conf.SPOTIFY_USERNAME + "/playlists/" + conf.SPOTIFY_PLAYLIST_ID + "/tracks?uris=" + currentSongUri

	// CREATING REQUEST
	client := &http.Client{}
	req, err := http.NewRequest("POST", spotifyUrl, nil)
	handleError(err, "create request to add song to playlist")

	// SENDING REQUEST
	req.Header.Add("Authorization", "Bearer "+token)
	_, err = client.Do(req)
	handleError(err, "sending request to add song to playlist")

	return err
}

func handleError(err error, context string) {
	if err != nil {
		log.Fatal("Error while "+context+": ", err)
	}
}
