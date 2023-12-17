package rpchandlers

import (
	"github.com/frin-network/frind/app/appmessage"
	"github.com/frin-network/frind/app/rpc/rpccontext"
	"github.com/frin-network/frind/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
