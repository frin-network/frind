package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/frin-network/frind/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.frindMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.frindMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.frindMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.frindMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.frindMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.frindMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.frindMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.frindMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.frindMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.frindMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.frindMessage_BanRequest{}),
	reflect.TypeOf(protowire.frindMessage_UnbanRequest{}),
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
