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

var prismObjectsDealsGrantUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update grant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "deal-id",
			Required:  true,
			PathParam: "dealId",
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
	Action:          handlePrismObjectsDealsGrantUpdate,
	HideHelpCommand: true,
}

var prismObjectsDealsGrantGet = cli.Command{
	Name:    "get",
	Usage:   "Get grant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "deal-id",
			Required:  true,
			PathParam: "dealId",
		},
	},
	Action:          handlePrismObjectsDealsGrantGet,
	HideHelpCommand: true,
}

func handlePrismObjectsDealsGrantUpdate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("deal-id") && len(unusedArgs) > 0 {
		cmd.Set("deal-id", unusedArgs[0])
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

	params := micro.PrismObjectDealGrantUpdateParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Objects.Deals.Grant.Update(
		ctx,
		cmd.Value("deal-id").(string),
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
		Title:          "prism:objects:deals:grant update",
		Transform:      transform,
	})
}

func handlePrismObjectsDealsGrantGet(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("deal-id") && len(unusedArgs) > 0 {
		cmd.Set("deal-id", unusedArgs[0])
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

	params := micro.PrismObjectDealGrantGetParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Objects.Deals.Grant.Get(
		ctx,
		cmd.Value("deal-id").(string),
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
		Title:          "prism:objects:deals:grant get",
		Transform:      transform,
	})
}
