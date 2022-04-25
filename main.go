package main

import (
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/vitzeno/gofetch/internal"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	_, err := internal.NewMemInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load memory info: %v\n", err)
		os.Exit(1)
	}

	_, err = internal.NewCPUInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load CPU info: %v \n", err)
		os.Exit(1)
	}

	_, err = internal.NewDiskInfo("/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load disk info: %v \n", err)
		os.Exit(1)
	}

	_, err = internal.NewHostInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load host info: %v \n", err)
		os.Exit(1)
	}

	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)
}
