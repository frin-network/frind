package protowire

import (
	"github.com/frin-network/frind/app/appmessage"
	"github.com/pkg/errors"
)

func (x *frindMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "frindMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest is nil")
	}
	return &appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage{}, nil
}

func (x *frindMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest) fromAppMessage(_ *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage) error {
	x.NotifyVirtualSelectedParentBlueScoreChangedRequest = &NotifyVirtualSelectedParentBlueScoreChangedRequestMessage{}
	return nil
}

func (x *frindMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "frindMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse is nil")
	}
	return x.NotifyVirtualSelectedParentBlueScoreChangedResponse.toAppMessage()
}

func (x *frindMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse) fromAppMessage(message *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyVirtualSelectedParentBlueScoreChangedResponse = &NotifyVirtualSelectedParentBlueScoreChangedResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyVirtualSelectedParentBlueScoreChangedResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyVirtualSelectedParentBlueScoreChangedResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *frindMessage_VirtualSelectedParentBlueScoreChangedNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "frindMessage_VirtualSelectedParentBlueScoreChangedNotification is nil")
	}
	return x.VirtualSelectedParentBlueScoreChangedNotification.toAppMessage()
}

func (x *frindMessage_VirtualSelectedParentBlueScoreChangedNotification) fromAppMessage(message *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage) error {
	x.VirtualSelectedParentBlueScoreChangedNotification = &VirtualSelectedParentBlueScoreChangedNotificationMessage{
		VirtualSelectedParentBlueScore: message.VirtualSelectedParentBlueScore,
	}
	return nil
}

func (x *VirtualSelectedParentBlueScoreChangedNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "VirtualSelectedParentBlueScoreChangedNotificationMessage is nil")
	}
	return &appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage{
		VirtualSelectedParentBlueScore: x.VirtualSelectedParentBlueScore,
	}, nil
}
