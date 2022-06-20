package files

import (
	rice "github.com/GeertJohan/go.rice"
)

var Static *rice.Box

func init() {
	Static = rice.MustFindBox("templates")
}
