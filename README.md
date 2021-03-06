# SAVE CURRENT SONG
Adds currently playing Spotify song to a pre-defined playlist by using the Last.FM API (which provides a "currently-playing"-feature) and then searching for the song title and artist using the Spotify API.

Sends notifications in Linux (notify-send) and OS X.

## Conf.json
The program requires a conf.json-file. A sample is included in the repo.

## Build and run
Build and run:

```bash
go build . && ./save-current-song
```

## Build for Linux from OS X
Build for linux:

```bash
env GOOS=linux GOARCH=amd64 go -o save-current-song-linux build -v .
```

## Logging
Logging is done to ./log.txt

To filter only found/saved songs: 

```bash
cat log.txt | grep -e "-----\|Found\|Saved"
```


