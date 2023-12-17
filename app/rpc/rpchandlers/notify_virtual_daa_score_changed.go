package rpchandlers

import (
	"github.com/frin-network/frind/app/appmessage"
	"github.com/frin-network/frind/app/rpc/rpccontext"
	"github.com/frin-network/frind/infrastructure/network/netadapter/router"
)

// HandleNotifyVirtualDaaScoreChanged handles the respectively named RPC command
func HandleNotifyVirtualDaaScoreChanged(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateVirtualDaaScoreChangedNotifications()

	response := appmessage.NewNotifyVirtualDaaScoreChangedResponseMessage()
	return response, nil
}
