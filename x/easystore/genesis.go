package functt

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/xuyp1991/cosaccount/x/easystore/types"
	// "github.com/cosmos/cosmos-sdk/x/simulation"
	// "math/rand"
	// "fmt"
)

// GenesisState - all auth state that must be provided at genesis
type GenesisState struct {
	Nullstring string `json:"null" yaml:"null"`
}

// NewGenesisState - Create a new genesis state
func NewGenesisState(tempstring string) GenesisState {
	return GenesisState{
		Nullstring: tempstring,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState("nothing")
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// startingProposalID, _ := k.GetProposalID(ctx)
	// depositParams := k.GetDepositParams(ctx)
	// votingParams := k.GetVotingParams(ctx)
	// tallyParams := k.GetTallyParams(ctx)
	// proposals := k.GetProposals(ctx)

	// var proposalsDeposits Deposits
	// var proposalsVotes Votes
	// for _, proposal := range proposals {
	// 	deposits := k.GetDeposits(ctx, proposal.ProposalID)
	// 	proposalsDeposits = append(proposalsDeposits, deposits...)

	// 	votes := k.GetVotes(ctx, proposal.ProposalID)
	// 	proposalsVotes = append(proposalsVotes, votes...)
	// }

	return GenesisState{
		Nullstring:"",
	}
}

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	// k.SetProposalID(ctx, data.StartingProposalID)
	// k.SetDepositParams(ctx, data.DepositParams)
	// k.SetVotingParams(ctx, data.VotingParams)
	// k.SetTallyParams(ctx, data.TallyParams)

	// // check if the deposits pool account exists
	// moduleAcc := k.GetGovernanceAccount(ctx)
	// if moduleAcc == nil {
	// 	panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	// }

	// var totalDeposits sdk.Coins
	// for _, deposit := range data.Deposits {
	// 	k.SetDeposit(ctx, deposit)
	// 	totalDeposits = totalDeposits.Add(deposit.Amount...)
	// }

	// for _, vote := range data.Votes {
	// 	k.SetVote(ctx, vote)
	// }

	// for _, proposal := range data.Proposals {
	// 	switch proposal.Status {
	// 	case StatusDepositPeriod:
	// 		k.InsertInactiveProposalQueue(ctx, proposal.ProposalID, proposal.DepositEndTime)
	// 	case StatusVotingPeriod:
	// 		k.InsertActiveProposalQueue(ctx, proposal.ProposalID, proposal.VotingEndTime)
	// 	}
	// 	k.SetProposal(ctx, proposal)
	// }

	// // add coins if not provided on genesis
	// if bk.GetAllBalances(ctx, moduleAcc.GetAddress()).IsZero() {
	// 	if err := bk.SetBalances(ctx, moduleAcc.GetAddress(), totalDeposits); err != nil {
	// 		panic(err)
	// 	}
	// 	supplyKeeper.SetModuleAccount(ctx, moduleAcc)
	// }
}

const (
	keyCommunityTax        = "communitytax"
	keyBaseProposerReward  = "baseproposerreward"
	keyBonusProposerReward = "bonusproposerreward"
)

// func ParamChanges(r *rand.Rand) []simulation.ParamChange {
// 	return []simulation.ParamChange{
// 		simulation.NewSimParamChange(types.ModuleName, keyCommunityTax,
// 			func(r *rand.Rand) string {
// 				return fmt.Sprintf("\"%s\"", "just test")
// 			},
// 		),
// 	}
// }

