package internal

import (
	"fmt"
	"os/user"
	"time"

	"github.com/gizak/termui/v3/widgets"
	"github.com/pkg/errors"
)

type UserView struct {
	Widget *widgets.Paragraph
}

func NewUserView() (*UserView, error) {
	user, err := user.Current()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load user info")
	}

	userWidget := widgets.NewParagraph()
	userWidget.Border = false
	userWidget.Text = fmt.Sprintf(" %s\n %s\n %s\n %s", user.Name, user.Username, user.HomeDir, time.Now().Format("15:04:05 2006-01-02"))

	return &UserView{
		Widget: userWidget,
	}, nil
}

func (u UserView) Update() error {
	user, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "Failed to load user info")
	}

	u.Widget.Text = fmt.Sprintf(" %s\n %s\n %s\n %s", user.Name, user.Username, user.HomeDir, time.Now().Format("15:04:05 2006-01-02"))

	return nil
}
