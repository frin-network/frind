package protowire

import (
	"github.com/frin-network/frind/app/appmessage"
	"github.com/pkg/errors"
)

func (x *frindMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "frindMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *frindMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
