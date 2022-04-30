package internal

import (
	"github.com/pkg/errors"

	"github.com/gizak/termui/v3/widgets"

	data "github.com/vitzeno/gofetch/internal/data"
)

type CPUView struct {
	Widget *widgets.Paragraph
}

func NewCPUView() (*CPUView, error) {
	cpuInfo, err := data.NewCPUInfo()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load CPU info")
	}

	cpuWidget := widgets.NewParagraph()
	cpuWidget.Title = "CPU"
	for _, cpu := range cpuInfo {
		cpuWidget.Text += cpu.String()
	}

	return &CPUView{
		Widget: cpuWidget,
	}, nil
}
