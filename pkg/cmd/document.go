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

var documentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create Document",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[any]{
			Name:     "crm",
			BodyPath: "crm",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default",
			Usage:    "Properties keyed by property slug. Values can be strings, numbers, booleans, arrays, or null.",
			BodyPath: "default",
		},
		&requestflag.Flag[any]{
			Name:     "extended",
			BodyPath: "extended",
		},
	},
	Action:          handleDocumentsCreate,
	HideHelpCommand: true,
}

var documentsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update Document",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "document-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[any]{
			Name:     "crm",
			BodyPath: "crm",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default",
			Usage:    "Properties keyed by property slug. Values can be strings, numbers, booleans, arrays, or null.",
			BodyPath: "default",
		},
		&requestflag.Flag[any]{
			Name:     "extended",
			BodyPath: "extended",
		},
	},
	Action:          handleDocumentsUpdate,
	HideHelpCommand: true,
}

var documentsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Documents",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "query",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[any]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "box",
			BodyPath: "boxes",
		},
		&requestflag.Flag[bool]{
			Name:     "deleted",
			BodyPath: "deleted",
		},
		&requestflag.Flag[[]string]{
			Name:     "source",
			BodyPath: "sources",
		},
	},
	Action:          handleDocumentsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"query": {
		&requestflag.InnerFlag[[]string]{
			Name:       "query.select",
			Usage:      "Property slugs to select. Use dot notation for relationships (e.g. attendee.contact.first_name)",
			InnerField: "select",
		},
		&requestflag.InnerFlag[string]{
			Name:       "query.combinator",
			Usage:      "Logical operator for combining filters",
			InnerField: "combinator",
		},
		&requestflag.InnerFlag[string]{
			Name:       "query.crm-id",
			InnerField: "crm_id",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "query.filter",
			Usage:      "Filters as [{ slug: { operator: value } }]. For select/multiselect properties, values must be option slugs",
			InnerField: "filter",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "query.limit",
			InnerField: "limit",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "query.page",
			InnerField: "page",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "query.sort",
			Usage:      "Sort order as [{ slug: direction }]. Array order determines sort priority",
			InnerField: "sort",
		},
	},
})

var documentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Document",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "document-id",
			Required: true,
		},
	},
	Action:          handleDocumentsDelete,
	HideHelpCommand: true,
}

func handleDocumentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.DocumentNewParams{
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Documents.New(ctx, params, options...)
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
		Title:          "documents create",
		Transform:      transform,
	})
}

func handleDocumentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.DocumentUpdateParams{
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

	return client.Documents.Update(
		ctx,
		cmd.Value("document-id").(string),
		params,
		options...,
	)
}

func handleDocumentsList(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.DocumentListParams{
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Documents.List(ctx, params, options...)
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
		Title:          "documents list",
		Transform:      transform,
	})
}

func handleDocumentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.DocumentDeleteParams{
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

	return client.Documents.Delete(
		ctx,
		cmd.Value("document-id").(string),
		params,
		options...,
	)
}
