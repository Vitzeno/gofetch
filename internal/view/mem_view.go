package internal

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/pkg/errors"

	data "github.com/vitzeno/gofetch/internal/data"
)

type MemView struct {
	Widget *widgets.Paragraph
	Gauge  *widgets.Gauge
}

func NewMemeView() (*MemView, error) {
	memInfo, err := data.NewMemInfo()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load memory info")
	}

	memWidget := widgets.NewParagraph()
	memWidget.Border = false
	memWidget.Text = memInfo.String()

	memGauge := widgets.NewGauge()
	memGauge.Percent = int(memInfo.UsedPercent)
	memGauge.BarColor = ui.ColorGreen

	return &MemView{
		Widget: memWidget,
		Gauge:  memGauge,
	}, nil
}

func (m MemView) Update() error {
	memInfo, err := data.NewMemInfo()
	if err != nil {
		return errors.Wrap(err, "Failed to load memory info")
	}

	m.Widget.Text = memInfo.String()
	m.Gauge.Percent = int(memInfo.UsedPercent)
	if memInfo.UsedPercent > 80 {
		m.Gauge.BarColor = ui.ColorRed
	} else {
		m.Gauge.BarColor = ui.ColorGreen
	}

	return nil
}
