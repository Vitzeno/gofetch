package internal

import (
	"github.com/pkg/errors"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	data "github.com/vitzeno/gofetch/internal/data"
)

type DiskView struct {
	Widget *widgets.Paragraph
	Gauge  *widgets.Gauge
}

func NewDiskView() (*DiskView, error) {
	diskInfo, err := data.NewDiskInfo("/")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load disk info")
	}

	diskWidget := widgets.NewParagraph()
	diskWidget.Title = "Disk"
	diskWidget.Text = diskInfo.String()

	diskGauge := widgets.NewGauge()
	diskGauge.Percent = int(diskInfo.UsedPercent)
	diskGauge.BarColor = ui.ColorGreen

	return &DiskView{
		Widget: diskWidget,
		Gauge:  diskGauge,
	}, nil
}