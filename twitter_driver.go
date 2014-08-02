package twitter

import (
	"github.com/hybridgroup/gobot"
)

type TwitterDriver struct {
	gobot.Driver
}

type TwitterInterface interface {
}

func NewTwitterDriver(a *TwitterAdaptor, name string) *TwitterDriver {
	return &TwitterDriver{
		Driver: *gobot.NewDriver(
			name,
			"twitter.TwitterDriver",
			a,
		),
	}
}

func (t *TwitterDriver) adaptor() *TwitterAdaptor {
	return t.Driver.Adaptor().(*TwitterAdaptor)
}

func (t *TwitterDriver) Start() bool { return true }
func (t *TwitterDriver) Halt() bool  { return true }
