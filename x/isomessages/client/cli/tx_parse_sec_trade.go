package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"xchange/x/isomessages/types"
)

var _ = strconv.Itoa(0)

func CmdParseSecTrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse-sec-trade [trad-conf-id] [trad-conf-dt-tm] [rltd-ordr-id] [rltd-csh-mvmnt-id]",
		Short: "Broadcast message parseSecTrade",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradConfId := args[0]
			argTradConfDtTm := args[1]
			argRltdOrdrId := args[2]
			argRltdCshMvmntId := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgParseSecTrade(
				clientCtx.GetFromAddress().String(),
				argTradConfId,
				argTradConfDtTm,
				argRltdOrdrId,
				argRltdCshMvmntId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
