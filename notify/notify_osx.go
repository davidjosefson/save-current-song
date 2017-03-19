// +build darwin

package notify

import (
	"log"
	"os/exec"
)

func Notify(foundSong string, addedSong string) {
	title := "ADD CURRENT SONG"
	text := "Found: " + foundSong + "\nAdded: " + addedSong

	cmd := exec.Command("/usr/bin/osascript", "-e", "display notification \""+text+"\" with title \""+title+"\"")
	err := cmd.Run()

	if err != nil {
		log.Println("Unable to display OS X notification.")
		log.Println(err.Error())
	}
}
