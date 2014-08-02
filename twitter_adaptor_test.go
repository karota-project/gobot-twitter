package twitter

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestTwitterAdaptor() *TwitterAdaptor {
	return NewTwitterAdaptor("myAdaptor")
}

func TestTwitterAdaptorConnect(t *testing.T) {
	a := initTestTwitterAdaptor()
	gobot.Expect(t, a.Connect(), true)
}

func TestTwitterAdaptorFinalize(t *testing.T) {
	a := initTestTwitterAdaptor()
	gobot.Expect(t, a.Finalize(), true)
}
