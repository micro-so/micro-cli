// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/micro-so/micro-cli/internal/apiquery"
	"github.com/micro-so/micro-cli/internal/requestflag"
	"github.com/micro-so/micro-sdk-go"
	"github.com/micro-so/micro-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var prismObjectsEventsGrantUpdate = cli.Command{
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
			Name:      "event-id",
			Required:  true,
			PathParam: "eventId",
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
	Action:          handlePrismObjectsEventsGrantUpdate,
	HideHelpCommand: true,
}

var prismObjectsEventsGrantGet = cli.Command{
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
			Name:      "event-id",
			Required:  true,
			PathParam: "eventId",
		},
	},
	Action:          handlePrismObjectsEventsGrantGet,
	HideHelpCommand: true,
}

func handlePrismObjectsEventsGrantUpdate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
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

	params := micro.PrismObjectEventGrantUpdateParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Objects.Events.Grant.Update(
		ctx,
		cmd.Value("event-id").(string),
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
		Title:          "prism:objects:events:grant update",
		Transform:      transform,
	})
}

func handlePrismObjectsEventsGrantGet(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
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

	params := micro.PrismObjectEventGrantGetParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Objects.Events.Grant.Get(
		ctx,
		cmd.Value("event-id").(string),
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
		Title:          "prism:objects:events:grant get",
		Transform:      transform,
	})
}
