package main

import (
	"fmt"
	"os"

	"github.com/benrowe/spotify/spotify"
)

func main() {

	if len(os.Args) <= 1 {
		help()
		return
	}

	cmd := os.Args[1:2][0]
	// @todo tidy this
	switch cmd {
	case "play":
		spotify.Play()
	case "pause":
		spotify.Pause()
	case "":
		uri := os.Args[2:3][0]
		spotify.PlayUri(uri)
	case "next":
		spotify.Next()
	case "current":
		curr, _ := spotify.CurrentSong()
		fmt.Println(curr)
	case "mute":
		spotify.Mute()
	case "unmute":
		spotify.Unmute()

	}
}

func help() {

}
