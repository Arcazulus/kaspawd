package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Arcazulus/kaspawd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.KaspawdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.KaspawdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.KaspawdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.KaspawdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.KaspawdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.KaspawdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.KaspawdMessage_BanRequest{}),
	reflect.TypeOf(protowire.KaspawdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
