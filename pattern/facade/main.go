package main

import (
	"facade/devices"
	"fmt"
)

// The Facade pattern simplifies interaction with a complex system by providing a unified interface
type Facade struct {
	tv    *devices.TV
	dvd   *devices.DVDPlayer
	sound *devices.SoundSystem
}

func newFacade(tv *devices.TV, dvd *devices.DVDPlayer, sound *devices.SoundSystem) *Facade {
	return &Facade{tv: tv, dvd: dvd, sound: sound}
}

func (f *Facade) watchMovie(movie string) {
	fmt.Println("Starting to watch")
	f.tv.On()
	f.tv.SetInput(movie)
	f.dvd.On()
	f.dvd.Play(movie)
	f.sound.On()
	f.sound.SetVolume(10)
	fmt.Println("Setting done")
}

func main() {
	facade := newFacade(&devices.TV{}, &devices.DVDPlayer{}, &devices.SoundSystem{})
	facade.watchMovie("The Matrix")
}
