// +build linux

package notify

import (
	"log"
	"os/exec"
)

func Notify(foundSong string, addedSong string) {
	title := "ADD CURRENT SONG"
	text := "Found: " + foundSong + "\nAdded: " + addedSong

	cmd := exec.Command("notify-send", "ADD CURRENT SONG\nFound: "+foundSong+"\nAdded: "+addedSong)
	err := cmd.Run()

	if err != nil 
		log.Println("Unable to display Linux notification.")
		log.Println(err.Error())
	}
}
