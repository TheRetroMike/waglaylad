package main

import (
	"context"
	"fmt"
	"github.com/waglayla/waglaylad/cmd/waglaylawallet/utils"
	"os"

	"github.com/waglayla/waglaylad/cmd/waglaylawallet/daemon/client"
	"github.com/waglayla/waglaylad/cmd/waglaylawallet/daemon/pb"
)

func createUnsignedTransaction(conf *createUnsignedTransactionConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()

	sendAmountLeor, err := utils.WalaToLeor(conf.SendAmount)

	if err != nil {
		return err
	}

	response, err := daemonClient.CreateUnsignedTransactions(ctx, &pb.CreateUnsignedTransactionsRequest{
		From:                     conf.FromAddresses,
		Address:                  conf.ToAddress,
		Amount:                   sendAmountLeor,
		IsSendAll:                conf.IsSendAll,
		UseExistingChangeAddress: conf.UseExistingChangeAddress,
	})
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Created unsigned transaction")
	fmt.Println(encodeTransactionsToHex(response.UnsignedTransactions))

	return nil
}
