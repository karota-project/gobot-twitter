package twitter

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestTwitterDriver() *TwitterDriver {
	return NewTwitterDriver(NewTwitterAdaptor("myAdaptor"), "myDriver")
}

func TestTwitterDriverStart(t *testing.T) {
	d := initTestTwitterDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestTwitterDriverHalt(t *testing.T) {
	d := initTestTwitterDriver()
	gobot.Expect(t, d.Halt(), true)
}
