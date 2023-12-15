package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"xchange/x/isomessages/types"
)

func CmdListSecuritiesTradeConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-securities-trade-confirmation",
		Short: "list all SecuritiesTradeConfirmation",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSecuritiesTradeConfirmationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SecuritiesTradeConfirmationAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSecuritiesTradeConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-securities-trade-confirmation [id]",
		Short: "shows a SecuritiesTradeConfirmation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetSecuritiesTradeConfirmationRequest{
				Id: id,
			}

			res, err := queryClient.SecuritiesTradeConfirmation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
