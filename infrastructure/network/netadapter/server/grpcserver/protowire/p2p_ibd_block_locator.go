package protowire

import (
	"github.com/waglayla/waglaylad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *WaglayladMessage_IbdBlockLocator) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "WaglayladMessage_IbdBlockLocator is nil")
	}
	return x.IbdBlockLocator.toAppMessage()
}

func (x *IbdBlockLocatorMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "IbdBlockLocatorMessage is nil")
	}
	targetHash, err := x.TargetHash.toDomain()
	if err != nil {
		return nil, err
	}
	blockLocatorHash, err := protoHashesToDomain(x.BlockLocatorHashes)
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgIBDBlockLocator{
		TargetHash:         targetHash,
		BlockLocatorHashes: blockLocatorHash,
	}, nil
}

func (x *WaglayladMessage_IbdBlockLocator) fromAppMessage(message *appmessage.MsgIBDBlockLocator) error {
	x.IbdBlockLocator = &IbdBlockLocatorMessage{
		TargetHash:         domainHashToProto(message.TargetHash),
		BlockLocatorHashes: domainHashesToProto(message.BlockLocatorHashes),
	}
	return nil
}
