// +build !linux,!darwin

package notify

func Notify(message string) {
	// if not linux or osx, do nothing
}
