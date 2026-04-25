// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/micro-cli/internal/apiquery"
	"github.com/stainless-sdks/micro-cli/internal/requestflag"
	"github.com/stainless-sdks/micro-go"
	"github.com/urfave/cli/v3"
)

var prismGrantRetrieveGrant = cli.Command{
	Name:    "retrieve-grant",
	Usage:   "Get grant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-type",
			Usage:    `Allowed values: "deal", "identity", "ai_chat_thread", "ai_chat_message", "document", "action", "event".`,
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
	},
	Action:          handlePrismGrantRetrieveGrant,
	HideHelpCommand: true,
}

var prismGrantUpdateGrant = cli.Command{
	Name:    "update-grant",
	Usage:   "Update grant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-type",
			Usage:    `Allowed values: "deal", "identity", "ai_chat_thread", "ai_chat_message", "document", "action", "event".`,
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "team-group-id",
			BodyPath: "team_group_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "team-id",
			BodyPath: "team_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "user-id",
			BodyPath: "user_id",
		},
	},
	Action:          handlePrismGrantUpdateGrant,
	HideHelpCommand: true,
}

func handlePrismGrantRetrieveGrant(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("object-type") && len(unusedArgs) > 0 {
		cmd.Set("object-type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.PrismGrantGetGrantParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Prism.Grant.GetGrant(
		ctx,
		micro.ObjectType(cmd.Value("object-type").(string)),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
}

func handlePrismGrantUpdateGrant(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("object-type") && len(unusedArgs) > 0 {
		cmd.Set("object-type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("object-id") && len(unusedArgs) > 0 {
		cmd.Set("object-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.PrismGrantUpdateGrantParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Prism.Grant.UpdateGrant(
		ctx,
		micro.ObjectType(cmd.Value("object-type").(string)),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
}
