package rpchandlers

import (
	"github.com/waglayla/waglaylad/app/appmessage"
	"github.com/waglayla/waglaylad/app/rpc/rpccontext"
	"github.com/waglayla/waglaylad/domain/consensus/model/externalapi"
	"github.com/waglayla/waglaylad/infrastructure/network/netadapter/router"
	"github.com/pkg/errors"
)

// HandleGetBlock handles the respectively named RPC command
func HandleGetBlock(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	getBlockRequest := request.(*appmessage.GetBlockRequestMessage)

	// Load the raw block bytes from the database.
	hash, err := externalapi.NewDomainHashFromString(getBlockRequest.Hash)
	if err != nil {
		errorMessage := &appmessage.GetBlockResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Hash could not be parsed: %s", err)
		return errorMessage, nil
	}

	block, err := context.Domain.Consensus().GetBlockEvenIfHeaderOnly(hash)
	if err != nil {
		errorMessage := &appmessage.GetBlockResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Block %s not found", hash)
		return errorMessage, nil
	}

	response := appmessage.NewGetBlockResponseMessage()

	if getBlockRequest.IncludeTransactions {
		response.Block = appmessage.DomainBlockToRPCBlock(block)
	} else {
		response.Block = appmessage.DomainBlockToRPCBlock(&externalapi.DomainBlock{Header: block.Header})
	}

	err = context.PopulateBlockWithVerboseData(response.Block, block.Header, block, getBlockRequest.IncludeTransactions)
	if err != nil {
		if errors.Is(err, rpccontext.ErrBuildBlockVerboseDataInvalidBlock) {
			errorMessage := &appmessage.GetBlockResponseMessage{}
			errorMessage.Error = appmessage.RPCErrorf("Block %s is invalid", hash)
			return errorMessage, nil
		}
		return nil, err
	}

	return response, nil
}
