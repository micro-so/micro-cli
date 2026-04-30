// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/micro-cli/internal/apiquery"
	"github.com/stainless-sdks/micro-cli/internal/requestflag"
	"github.com/stainless-sdks/micro-go"
	"github.com/stainless-sdks/micro-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var prismGrantRetrieveGrant = cli.Command{
	Name:    "retrieve-grant",
	Usage:   "Get grant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "object-type",
			Usage:     `Allowed values: "deal", "identity", "ai_chat_thread", "ai_chat_message", "document", "action", "event".`,
			Required:  true,
			PathParam: "objectType",
		},
		&requestflag.Flag[string]{
			Name:      "object-id",
			Required:  true,
			PathParam: "objectId",
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
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "object-type",
			Usage:     `Allowed values: "deal", "identity", "ai_chat_thread", "ai_chat_message", "document", "action", "event".`,
			Required:  true,
			PathParam: "objectType",
		},
		&requestflag.Flag[string]{
			Name:      "object-id",
			Required:  true,
			PathParam: "objectId",
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

	params := micro.PrismGrantGetGrantParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Grant.GetGrant(
		ctx,
		micro.ObjectType(cmd.Value("object-type").(string)),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "prism:grant retrieve-grant",
		Transform:      transform,
	})
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

	params := micro.PrismGrantUpdateGrantParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Grant.UpdateGrant(
		ctx,
		micro.ObjectType(cmd.Value("object-type").(string)),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "prism:grant update-grant",
		Transform:      transform,
	})
}
