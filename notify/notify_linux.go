// +build linux

package notify

import (
	"log"
	"os/exec"
)

func Notify(foundSong string, savedSong string) {
	cmd := exec.Command("notify-send", "SAVE CURRENT SONG\nFound: "+foundSong+"\nSaved: "+savedSong)
	err := cmd.Run()

	if err != nil {
		log.Println("Unable to display Linux notification.")
		log.Println(err.Error())
	}
}
