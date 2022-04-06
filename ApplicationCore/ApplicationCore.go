package ApplicationCore

import (
	"os"
)

func (ApplicationCore *ApplicationCore) Start() (Error error) {
	_, Error = os.Stat("Settings.json")
	if Error != nil {

		return Error
	}
	return Error
}
