package twitter

import (
	"github.com/hybridgroup/gobot"
)

type TwitterAdaptor struct {
	gobot.Adaptor
}

func NewTwitterAdaptor(name string) *TwitterAdaptor {
	return &TwitterAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"twitter.TwitterAdaptor",
		),
	}
}

func (t *TwitterAdaptor) Connect() bool {
	return true
}

func (t *TwitterAdaptor) Finalize() bool {
	return true
}
