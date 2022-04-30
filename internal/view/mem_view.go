package internal

import (
	"fmt"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	data "github.com/vitzeno/gofetch/internal/data"
)

type MemView struct {
	Widget *widgets.Paragraph
	Gauge  *widgets.Gauge
}

func NewMemeView() (*MemView, error) {
	memInfo, err := data.NewMemInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load memory info: %v\n", err)
		os.Exit(1)
	}

	memWidget := widgets.NewParagraph()
	memWidget.Title = "Memory"
	memWidget.Border = false
	memWidget.Text = memInfo.String()

	memGauge := widgets.NewGauge()
	memGauge.Percent = int(memInfo.UsedPercent)
	memGauge.BarColor = ui.ColorGreen
	//memGauge.Border = false

	return &MemView{
		Widget: memWidget,
		Gauge:  memGauge,
	}, nil
}
