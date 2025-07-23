package labels

import "fyne.io/fyne/v2/widget"

type Label struct {
	*widget.Label
}

func NewLabel(title string) *Label {
	lb := widget.NewLabel(title)

	return &Label{lb}
}
