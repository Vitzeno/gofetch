package internal

import (
	"github.com/pkg/errors"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	data "github.com/vitzeno/gofetch/internal/data"
)

type CPUView struct {
	Widget *widgets.Paragraph
	Gauge  *widgets.Gauge
}

func NewCPUView() (*CPUView, error) {
	cpuInfo, err := data.NewCPUInfo()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load CPU info")
	}

	cpuWidget := widgets.NewParagraph()
	cpuWidget.TitleStyle = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	cpuWidget.Title = "CPU"
	cpuWidget.Border = false

	for _, cpu := range cpuInfo {
		cpuWidget.Text += cpu.String()
	}

	cpuGauge := widgets.NewGauge()
	cpuGauge.Percent = int(cpuInfo[0].Usage)
	cpuGauge.BarColor = ui.ColorGreen

	return &CPUView{
		Widget: cpuWidget,
		Gauge:  cpuGauge,
	}, nil
}

func (c CPUView) Update() error {
	cpuInfo, err := data.NewCPUInfo()
	if err != nil {
		return errors.Wrap(err, "Failed to load CPU info")
	}

	c.Widget.Text = ""
	for _, cpu := range cpuInfo {
		c.Widget.Text += cpu.String()
	}

	c.Gauge.Percent = int(cpuInfo[0].Usage)

	return nil
}
