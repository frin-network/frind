package rpchandlers

import (
	"github.com/frin-network/frind/app/appmessage"
	"github.com/frin-network/frind/app/rpc/rpccontext"
	"github.com/frin-network/frind/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
