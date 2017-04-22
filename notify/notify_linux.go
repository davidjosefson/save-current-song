// +build linux

package notify

import (
	"log"
	"os/exec"
)

func Notify(addedSong string) {
	cmd := exec.Command("notify-send", "SAVE CURRENT SONG\nAdded: "+addedSong)
	err := cmd.Run()

	if err != nil {
		log.Println("Unable to display Linux notification.")
		log.Println(err.Error())
	}
}
