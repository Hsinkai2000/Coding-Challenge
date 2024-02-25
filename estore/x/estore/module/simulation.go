package estore

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"estore/testutil/sample"
	estoresimulation "estore/x/estore/simulation"
	"estore/x/estore/types"
)

// avoid unused import issue
var (
	_ = estoresimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateProduct = "op_weight_msg_product"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProduct int = 100

	opWeightMsgUpdateProduct = "op_weight_msg_product"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProduct int = 100

	opWeightMsgDeleteProduct = "op_weight_msg_product"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProduct int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	estoreGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ProductList: []types.Product{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ProductCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&estoreGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProduct int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProduct, &weightMsgCreateProduct, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProduct = defaultWeightMsgCreateProduct
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProduct,
		estoresimulation.SimulateMsgCreateProduct(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProduct int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProduct, &weightMsgUpdateProduct, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProduct = defaultWeightMsgUpdateProduct
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProduct,
		estoresimulation.SimulateMsgUpdateProduct(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProduct int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteProduct, &weightMsgDeleteProduct, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProduct = defaultWeightMsgDeleteProduct
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProduct,
		estoresimulation.SimulateMsgDeleteProduct(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProduct,
			defaultWeightMsgCreateProduct,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				estoresimulation.SimulateMsgCreateProduct(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateProduct,
			defaultWeightMsgUpdateProduct,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				estoresimulation.SimulateMsgUpdateProduct(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteProduct,
			defaultWeightMsgDeleteProduct,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				estoresimulation.SimulateMsgDeleteProduct(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
