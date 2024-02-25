package estore

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "estore/api/estore/estore"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ProductAll",
					Use:       "list-product",
					Short:     "List all product",
				},
				{
					RpcMethod:      "Product",
					Use:            "show-product [id]",
					Short:          "Shows a product by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateProduct",
					Use:            "create-product [title] [description] [price] [quantity]",
					Short:          "Create product",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "price"}, {ProtoField: "quantity"}},
				},
				{
					RpcMethod:      "UpdateProduct",
					Use:            "update-product [id] [title] [description] [price] [quantity]",
					Short:          "Update product",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "price"}, {ProtoField: "quantity"}},
				},
				{
					RpcMethod:      "DeleteProduct",
					Use:            "delete-product [id]",
					Short:          "Delete product",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
