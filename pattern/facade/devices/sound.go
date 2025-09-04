package devices

import "fmt"

type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Turning on the Sound System...")
}

func (s *SoundSystem) SetVolume(volume int) {
	fmt.Printf("Setting sound system volume to %d...\n", volume)
}
