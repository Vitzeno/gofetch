package internal

import (
	"fmt"
	"os/user"

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
	userWidget.Text = fmt.Sprintf(" %s\n %s\n %s", user.Name, user.Username, user.HomeDir)

	return &UserView{
		Widget: userWidget,
	}, nil
}
