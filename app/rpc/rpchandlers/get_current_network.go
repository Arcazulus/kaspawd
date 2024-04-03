package rpchandlers

import (
	"github.com/Arcazulus/kaspawd/app/appmessage"
	"github.com/Arcazulus/kaspawd/app/rpc/rpccontext"
	"github.com/Arcazulus/kaspawd/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
