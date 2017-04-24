// +build linux

package notify

import (
	"log"
	"os/exec"
)

func Notify(message string) {
	cmd := exec.Command("notify-send", "SAVE CURRENT SONG\n"+message)
	err := cmd.Run()

	if err != nil {
		log.Println("Unable to display Linux notification.")
		log.Println(err.Error())
	}
}
