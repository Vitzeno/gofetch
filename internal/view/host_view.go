package internal

import (
	"github.com/pkg/errors"

	"github.com/gizak/termui/v3/widgets"

	data "github.com/vitzeno/gofetch/internal/data"
)

type HostView struct {
	Widget *widgets.Paragraph
}

func NewHostView() (*HostView, error) {
	hostInfo, err := data.NewHostInfo()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load host info")
	}

	hostWidget := widgets.NewParagraph()
	hostWidget.Title = "Host"
	hostWidget.Border = false
	hostWidget.Text = hostInfo.String()

	return &HostView{
		Widget: hostWidget,
	}, nil
}
