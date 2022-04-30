package main

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/vitzeno/gofetch/internal"
)

const borderSize = 1
const GOPHER_IMAGE = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	baseGrid := ui.NewGrid()
	xDim, yDim := ui.TerminalDimensions()

	defer ui.Close()

	memInfo, err := internal.NewMemInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load memory info: %v\n", err)
		os.Exit(1)
	}

	cpuInfo, err := internal.NewCPUInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load CPU info: %v \n", err)
		os.Exit(1)
	}

	diskInfo, err := internal.NewDiskInfo("/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load disk info: %v \n", err)
		os.Exit(1)
	}

	hostInfo, err := internal.NewHostInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load host info: %v \n", err)
		os.Exit(1)
	}

	image, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(GOPHER_IMAGE)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load gopher image: %v \n", err)
		os.Exit(1)
	}
	img := widgets.NewImage(image)

	memWidget := widgets.NewParagraph()
	memWidget.Title = "Memory"
	memWidget.Text = memInfo.String()

	memGauge := widgets.NewGauge()
	memGauge.SetRect(0, 0, xDim/2-borderSize, yDim/2-borderSize)
	memGauge.Percent = int(memInfo.UsedPercent)
	memGauge.BarColor = ui.ColorGreen

	cpuWidget := widgets.NewParagraph()
	cpuWidget.Title = "CPU"
	cpuWidget.Text = cpuInfo[0].String()

	diskWidget := widgets.NewParagraph()
	diskWidget.Title = "Disk"
	diskWidget.Text = diskInfo.String()

	diskPie := widgets.NewGauge()
	diskPie.SetRect(0, 0, xDim/2-borderSize, yDim/2-borderSize)
	diskPie.Percent = int(diskInfo.UsedPercent)
	diskPie.BarColor = ui.ColorGreen

	hostWidget := widgets.NewParagraph()
	hostWidget.Title = "Host"
	hostWidget.Text = hostInfo.String()

	baseGrid.SetRect(0, 0, xDim, yDim)
	baseGrid.Set(
		ui.NewCol(0.5, img),
		ui.NewCol(0.5,
			ui.NewRow(0.2, cpuWidget),
			ui.NewRow(0.2, memWidget),
			ui.NewRow(0.1, memGauge),
			ui.NewRow(0.2, diskWidget),
			ui.NewRow(0.1, diskPie),
			ui.NewRow(0.2, hostWidget),
		),
	)

	ui.Render(baseGrid)

	// wait for key press
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
