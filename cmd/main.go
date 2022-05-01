package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/pkg/errors"

	view "github.com/vitzeno/gofetch/internal/view"
)

const (
	WIN_LOGO_PATH      = "../images/win.png"
	LINUX_LOGO_PATH    = "../images/linux.png"
	UBUNTU_LOGO_PATH   = "../images/ubuntu.png"
	APPLE_LOGO_PATH    = "../images/apple.png"
	GOPHER_LOGO_BASE64 = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`

	UPDATE_INTERVAL = time.Second * 1
)

func updateViews(view ...view.View) error {
	for _, v := range view {
		err := v.Update()
		if err != nil {
			return errors.Wrap(err, "Failed to update view")
		}
	}

	return nil
}

func getOSLogo() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return WIN_LOGO_PATH
	case "darwin":
		return APPLE_LOGO_PATH
	case "linux":
		return UBUNTU_LOGO_PATH
	default:
		fmt.Printf("%s.\n", os)
	}

	return GOPHER_LOGO_BASE64
}

func main() {
	ticker := time.NewTicker(UPDATE_INTERVAL)
	done := make(chan bool)
	args := os.Args[1:]

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer ui.Close()

	memView, err := view.NewMemeView()
	if err != nil {
		log.Fatalf("Failed to load memory view: %v", err)
		os.Exit(1)
	}

	cpuView, err := view.NewCPUView()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load CPU view: %v\n", err)
		os.Exit(1)
	}

	diskView, err := view.NewDiskView()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load disk view: %v\n", err)
		os.Exit(1)
	}

	hostView, err := view.NewHostView()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load host view: %v\n", err)
		os.Exit(1)
	}

	var asciiImage *view.AsciiImageView
	if len(args) > 0 {
		asciiImage, err = view.NewAsciiImageView(args[0], view.WithReversed(true))
	} else {
		asciiImage, err = view.NewAsciiImageView(getOSLogo(), view.WithReversed(true))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load ascii image view: %v\n", err)
		os.Exit(1)
	}

	userView, err := view.NewUserView()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load user view: %v\n", err)
		os.Exit(1)
	}

	baseGrid := ui.NewGrid()
	xDim, yDim := ui.TerminalDimensions()

	baseGrid.SetRect(0, 0, xDim, yDim)
	baseGrid.Set(
		ui.NewCol(0.55,
			ui.NewRow(0.8, asciiImage.Image),
			ui.NewRow(0.2, userView.Widget),
		),
		ui.NewCol(0.45,
			ui.NewRow(0.25,
				ui.NewRow(0.8, cpuView.Widget),
				ui.NewRow(0.2, cpuView.Gauge)),
			ui.NewRow(0.25,
				ui.NewRow(0.8, memView.Widget),
				ui.NewRow(0.2, memView.Gauge)),
			ui.NewRow(0.25,
				ui.NewRow(0.8, diskView.Widget),
				ui.NewRow(0.2, diskView.Gauge)),
			ui.NewRow(0.25, hostView.Widget),
		),
	)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				xDim, yDim := ui.TerminalDimensions()
				baseGrid.SetRect(0, 0, xDim, yDim)

				err := updateViews(cpuView, memView, diskView, hostView, userView)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to update views: %v\n", err)
					os.Exit(1)
				}

				ui.Render(baseGrid)
			}
		}
	}()

	// wait for key press
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}

	ticker.Stop()
	done <- true
}
