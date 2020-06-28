package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return client.GetCommands(
		queryPlanCmd(cdc),
		queryPlansCmd(cdc),
		querySubscriptionCmd(cdc),
		querySubscriptionsCmd(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Subscription module sub-commands",
	}

	cmd.AddCommand(client.PostCommands(
		txAddPlanCmd(cdc),
		txSetPlanStatusCmd(cdc),
		txAddNodeForPlanCmd(cdc),
		txRemoveNodeForPlanCmd(cdc),
		client.LineBreak,
		txStartPlanSubscription(cdc),
		txStartNodeSubscription(cdc),
		txAddAddressForSubscription(cdc),
		txRemoveAddressForSubscription(cdc),
		txEndSubscription(cdc),
	)...)

	return []*cobra.Command{cmd}
}
