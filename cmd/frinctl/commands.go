package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/frin-network/frind/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.FrindMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.FrindMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.FrindMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.FrindMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.FrindMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.FrindMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.FrindMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.FrindMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.FrindMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.FrindMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.FrindMessage_BanRequest{}),
	reflect.TypeOf(protowire.FrindMessage_UnbanRequest{}),
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
