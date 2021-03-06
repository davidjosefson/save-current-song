// +build darwin

package notify

import (
	"log"
	"os/exec"
)

func Notify(message string) {
	title := "SAVE CURRENT SONG"
	text := message

	cmd := exec.Command("/usr/bin/osascript", "-e", "display notification \""+text+"\" with title \""+title+"\"")
	err := cmd.Run()

	if err != nil {
		log.Println("Unable to display OS X notification.")
		log.Println(err.Error())
	}
}
