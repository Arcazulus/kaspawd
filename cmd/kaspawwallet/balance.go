package main

import (
	"context"
	"fmt"

	"github.com/Arcazulus/kaspawd/cmd/kaspawwallet/daemon/client"
	"github.com/Arcazulus/kaspawd/cmd/kaspawwallet/daemon/pb"
	"github.com/Arcazulus/kaspawd/cmd/kaspawwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatKap(addressBalance.Available), utils.FormatKap(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, KAP %s %s%s\n", utils.FormatKap(response.Available), utils.FormatKap(response.Pending), pendingSuffix)

	return nil
}
