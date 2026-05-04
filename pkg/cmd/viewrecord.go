// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/micro-so/micro-sdk-go"
	"github.com/micro-so/micro-sdk-go/option"
	"github.com/stainless-sdks/micro-cli/internal/apiquery"
	"github.com/stainless-sdks/micro-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var viewsRecordsList = cli.Command{
	Name:    "list",
	Usage:   "List records selected by a view (filters and sorts applied; pinned record_order\noverlaid first)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "view-object-type",
			Usage:     `Allowed values: "action", "deal", "document", "event", "identity", "organization".`,
			Required:  true,
			PathParam: "viewObjectType",
		},
		&requestflag.Flag[string]{
			Name:      "view-id",
			Required:  true,
			PathParam: "viewId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			QueryPath: "page",
		},
	},
	Action:          handleViewsRecordsList,
	HideHelpCommand: true,
}

var viewsRecordsPin = cli.Command{
	Name:    "pin",
	Usage:   "Pin a record to the view (append to record_order)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "view-object-type",
			Usage:     `Allowed values: "action", "deal", "document", "event", "identity", "organization".`,
			Required:  true,
			PathParam: "viewObjectType",
		},
		&requestflag.Flag[string]{
			Name:      "view-id",
			Required:  true,
			PathParam: "viewId",
		},
		&requestflag.Flag[string]{
			Name:      "object-id",
			Required:  true,
			PathParam: "objectId",
		},
	},
	Action:          handleViewsRecordsPin,
	HideHelpCommand: true,
}

var viewsRecordsReorder = cli.Command{
	Name:    "reorder",
	Usage:   "Bulk reorder pinned records",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "view-object-type",
			Usage:     `Allowed values: "action", "deal", "document", "event", "identity", "organization".`,
			Required:  true,
			PathParam: "viewObjectType",
		},
		&requestflag.Flag[string]{
			Name:      "view-id",
			Required:  true,
			PathParam: "viewId",
		},
		&requestflag.Flag[[]string]{
			Name:     "object-id",
			Required: true,
			BodyPath: "object_ids",
		},
	},
	Action:          handleViewsRecordsReorder,
	HideHelpCommand: true,
}

var viewsRecordsUnpin = cli.Command{
	Name:    "unpin",
	Usage:   "Unpin a record from the view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "view-object-type",
			Usage:     `Allowed values: "action", "deal", "document", "event", "identity", "organization".`,
			Required:  true,
			PathParam: "viewObjectType",
		},
		&requestflag.Flag[string]{
			Name:      "view-id",
			Required:  true,
			PathParam: "viewId",
		},
		&requestflag.Flag[string]{
			Name:      "object-id",
			Required:  true,
			PathParam: "objectId",
		},
	},
	Action:          handleViewsRecordsUnpin,
	HideHelpCommand: true,
}

func handleViewsRecordsList(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-object-type") && len(unusedArgs) > 0 {
		cmd.Set("view-object-type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("view-id") && len(unusedArgs) > 0 {
		cmd.Set("view-id", unusedArgs[0])
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

	params := micro.ViewRecordListParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.Records.List(
		ctx,
		micro.ViewRecordListParamsViewObjectType(cmd.Value("view-object-type").(string)),
		cmd.Value("view-id").(string),
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
		Title:          "views:records list",
		Transform:      transform,
	})
}

func handleViewsRecordsPin(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-object-type") && len(unusedArgs) > 0 {
		cmd.Set("view-object-type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("view-id") && len(unusedArgs) > 0 {
		cmd.Set("view-id", unusedArgs[0])
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

	params := micro.ViewRecordPinParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	return client.Views.Records.Pin(
		ctx,
		micro.ViewRecordPinParamsViewObjectType(cmd.Value("view-object-type").(string)),
		cmd.Value("view-id").(string),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
}

func handleViewsRecordsReorder(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-object-type") && len(unusedArgs) > 0 {
		cmd.Set("view-object-type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("view-id") && len(unusedArgs) > 0 {
		cmd.Set("view-id", unusedArgs[0])
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

	params := micro.ViewRecordReorderParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	return client.Views.Records.Reorder(
		ctx,
		micro.ViewRecordReorderParamsViewObjectType(cmd.Value("view-object-type").(string)),
		cmd.Value("view-id").(string),
		params,
		options...,
	)
}

func handleViewsRecordsUnpin(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-object-type") && len(unusedArgs) > 0 {
		cmd.Set("view-object-type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("view-id") && len(unusedArgs) > 0 {
		cmd.Set("view-id", unusedArgs[0])
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

	params := micro.ViewRecordUnpinParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	return client.Views.Records.Unpin(
		ctx,
		micro.ViewRecordUnpinParamsViewObjectType(cmd.Value("view-object-type").(string)),
		cmd.Value("view-id").(string),
		cmd.Value("object-id").(string),
		params,
		options...,
	)
}
