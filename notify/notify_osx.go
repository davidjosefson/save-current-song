// +build darwin

package notify

import (
	"log"
	"os/exec"
)

func Notify(foundSong string, savedSong string) {
	title := "SAVE CURRENT SONG"
	text := "Found: " + foundSong + "\nSaved: " + savedSong

	cmd := exec.Command("/usr/bin/osascript", "-e", "display notification \""+text+"\" with title \""+title+"\"")
	err := cmd.Run()

	if err != nil {
		log.Println("Unable to display OS X notification.")
		log.Println(err.Error())
	}
}
