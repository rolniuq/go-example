package containers

import (
	"fyne.io/fyne/v2"
)

type Container struct {
	*fyne.Container
}

func NewContainer() *Container {
	return &Container{}
}
