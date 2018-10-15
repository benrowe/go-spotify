package spotify

import (
	"fmt"

	"github.com/everdev/mack"
)

const StatusPlaying = 1
const StatusStopped = 2

// Song details
type Song struct {
	URI      string
	Track    string
	Album    string
	Artist   string
	Duration uint
}

// Status of player
type Status struct {
	State    uint
	Song     Song
	Position uint
	Volume   uint
}

func run(cmd string) (string, error) {
	return mack.Tell("Spotify", cmd)
}

// GetStatus of spotify
func GetStatus() (Status, error) {
	return Status{}, nil
}

// Play spotify
func Play() (Song, error) {
	_, err := run("play")
	return Song{}, err
}

// PlayUri
func PlayUri(uri string) (Song, error) {
	_, err := run("play track uri")
	return Song{}, err
}

// Next track
func Next() error {
	_, err := run("next track")
	return err
}

// Pause current track
func Pause() error {
	_, err := run("pause")
	return err
}

func Previous() error {
	return nil
}

func Mute() error {
	_, err := run("set sound volume to 0")
	return err
}

func Unmute() error {
	_, err := run("set sound volume to 100")
	return err
}

func Quit() error {
	_, err := run("quit")
	return err
}

// CurrentSong
func CurrentSong() (Song, error) {
	result := Song{}
	fields := map[string]string{
		"Artist":   "artist of current track as string",
		"Album":    "album of current track as string",
		"Track":    "name of current track as string",
		"Duration": "duration of current track",
	}
	// var song Song
	for _, val := range fields {
		output, _ := run(val)
		fmt.Println(output)
		// song[key] = run(val)
	}
	return Song{
		Album:    "",
		Artist:   "",
		Track:    "",
		Duration: 0,
	}, nil
}
