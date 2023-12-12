package protowire

import (
	"github.com/frin-network/frind/app/appmessage"
	"github.com/pkg/errors"
)

func (x *frindMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "frindMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *frindMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
