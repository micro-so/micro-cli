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

var viewsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a view bundle (view + select/filter/sort)",
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
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "view-type",
			Required: true,
			BodyPath: "view_type",
		},
		&requestflag.Flag[string]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "aggregation-prop-def-id",
			BodyPath: "aggregation_prop_def_id",
		},
		&requestflag.Flag[*string]{
			Name:     "aggregation-type",
			BodyPath: "aggregation_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "column-layout",
			BodyPath: "column_layout",
		},
		&requestflag.Flag[string]{
			Name:     "combinator",
			Usage:    `Allowed values: "AND", "OR".`,
			BodyPath: "combinator",
		},
		&requestflag.Flag[string]{
			Name:     "created-at",
			BodyPath: "created_at",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "filter",
			Usage:    "Each entry is { slug: { comparator: value } }",
			BodyPath: "filter",
		},
		&requestflag.Flag[*string]{
			Name:     "group-by",
			Usage:    "Property slug to group by",
			BodyPath: "group_by",
		},
		&requestflag.Flag[any]{
			Name:     "group-hidden-option-ids",
			BodyPath: "group_hidden_option_ids",
		},
		&requestflag.Flag[*bool]{
			Name:     "group-hide-empty",
			BodyPath: "group_hide_empty",
		},
		&requestflag.Flag[*string]{
			Name:     "group-sort",
			BodyPath: "group_sort",
		},
		&requestflag.Flag[*string]{
			Name:     "icon",
			BodyPath: "icon",
		},
		&requestflag.Flag[*string]{
			Name:     "list-id",
			BodyPath: "list_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "select",
			Usage:    "Property slugs (dot-paths permitted for refs)",
			BodyPath: "select",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "sort",
			Usage:    "Each entry is { slug: 'asc' | 'desc' }",
			BodyPath: "sort",
		},
		&requestflag.Flag[*int64]{
			Name:     "sort-order",
			BodyPath: "sort_order",
		},
		&requestflag.Flag[*string]{
			Name:     "team-id",
			BodyPath: "team_id",
		},
		&requestflag.Flag[*string]{
			Name:     "updated-at",
			BodyPath: "updated_at",
		},
		&requestflag.Flag[*string]{
			Name:     "user-id",
			BodyPath: "user_id",
		},
	},
	Action:          handleViewsCreate,
	HideHelpCommand: true,
}

var viewsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a view bundle (select/filter/sort arrays are replaced wholesale when\nsupplied)",
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
		&requestflag.Flag[*string]{
			Name:     "aggregation-prop-def-id",
			BodyPath: "aggregation_prop_def_id",
		},
		&requestflag.Flag[*string]{
			Name:     "aggregation-type",
			BodyPath: "aggregation_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "column-layout",
			BodyPath: "column_layout",
		},
		&requestflag.Flag[string]{
			Name:     "combinator",
			Usage:    `Allowed values: "AND", "OR".`,
			BodyPath: "combinator",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "filter",
			BodyPath: "filter",
		},
		&requestflag.Flag[*string]{
			Name:     "group-by",
			BodyPath: "group_by",
		},
		&requestflag.Flag[any]{
			Name:     "group-hidden-option-ids",
			BodyPath: "group_hidden_option_ids",
		},
		&requestflag.Flag[*bool]{
			Name:     "group-hide-empty",
			BodyPath: "group_hide_empty",
		},
		&requestflag.Flag[*string]{
			Name:     "group-sort",
			BodyPath: "group_sort",
		},
		&requestflag.Flag[*string]{
			Name:     "icon",
			BodyPath: "icon",
		},
		&requestflag.Flag[*string]{
			Name:     "list-id",
			BodyPath: "list_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "select",
			BodyPath: "select",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "sort",
			BodyPath: "sort",
		},
		&requestflag.Flag[*int64]{
			Name:     "sort-order",
			BodyPath: "sort_order",
		},
		&requestflag.Flag[*string]{
			Name:     "team-id",
			BodyPath: "team_id",
		},
		&requestflag.Flag[*string]{
			Name:     "user-id",
			BodyPath: "user_id",
		},
		&requestflag.Flag[string]{
			Name:     "view-type",
			BodyPath: "view_type",
		},
	},
	Action:          handleViewsUpdate,
	HideHelpCommand: true,
}

var viewsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a view bundle",
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
	},
	Action:          handleViewsDelete,
	HideHelpCommand: true,
}

var viewsGet = cli.Command{
	Name:    "get",
	Usage:   "Read a view bundle",
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
	},
	Action:          handleViewsGet,
	HideHelpCommand: true,
}

func handleViewsCreate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-object-type") && len(unusedArgs) > 0 {
		cmd.Set("view-object-type", unusedArgs[0])
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

	params := micro.ViewNewParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.New(
		ctx,
		micro.ViewNewParamsViewObjectType(cmd.Value("view-object-type").(string)),
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
		Title:          "views create",
		Transform:      transform,
	})
}

func handleViewsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := micro.ViewUpdateParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.Update(
		ctx,
		micro.ViewUpdateParamsViewObjectType(cmd.Value("view-object-type").(string)),
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
		Title:          "views update",
		Transform:      transform,
	})
}

func handleViewsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := micro.ViewDeleteParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	return client.Views.Delete(
		ctx,
		micro.ViewDeleteParamsViewObjectType(cmd.Value("view-object-type").(string)),
		cmd.Value("view-id").(string),
		params,
		options...,
	)
}

func handleViewsGet(ctx context.Context, cmd *cli.Command) error {
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

	params := micro.ViewGetParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.Get(
		ctx,
		micro.ViewGetParamsViewObjectType(cmd.Value("view-object-type").(string)),
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
		Title:          "views get",
		Transform:      transform,
	})
}
