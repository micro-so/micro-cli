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

var prismObjectsOrganizationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create object",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
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
			Usage:    "Properties keyed by property slug. Values can be strings, numbers, booleans, arrays, or null. For select/multiselect properties, values may be option slugs or option UUIDs on write; option slugs are returned on read.",
			BodyPath: "default",
		},
		&requestflag.Flag[any]{
			Name:     "extended",
			BodyPath: "extended",
		},
	},
	Action:          handlePrismObjectsOrganizationsCreate,
	HideHelpCommand: true,
}

var prismObjectsOrganizationsBulkCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "bulk-create",
	Usage:   "Import multiple objects in batch. Properties are keyed by slug. Automatically\nroutes based on size: <100 records sync (immediate response), >=100 records\nasync (S3/Lambda with WebSocket progress)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "object",
			Usage:    "Array of objects to import with property values keyed by slug",
			Required: true,
			BodyPath: "objects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
	},
	Action:          handlePrismObjectsOrganizationsBulkCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"object": {
		&requestflag.InnerFlag[string]{
			Name:       "object.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "object.crm",
			InnerField: "crm",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "object.default",
			Usage:      "Properties keyed by property slug. Values can be strings, numbers, booleans, arrays, or null. For select/multiselect properties, values may be option slugs or option UUIDs on write; option slugs are returned on read.",
			InnerField: "default",
		},
		&requestflag.InnerFlag[any]{
			Name:       "object.extended",
			InnerField: "extended",
		},
	},
	"options": {
		&requestflag.InnerFlag[bool]{
			Name:       "options.case-insensitive",
			Usage:      "Whether deduplication should be case insensitive",
			InnerField: "caseInsensitive",
		},
		&requestflag.InnerFlag[string]{
			Name:       "options.crm-id",
			Usage:      "App/CRM ID for context (optional)",
			InnerField: "crm_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "options.dedupe-by",
			Usage:      "Property slug to deduplicate on",
			InnerField: "dedupe_by",
		},
	},
})

var prismObjectsOrganizationsQuery = requestflag.WithInnerFlags(cli.Command{
	Name:    "query",
	Usage:   "Query v2",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
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
	Action:          handlePrismObjectsOrganizationsQuery,
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
			Usage:      "Filters as [{ slug: { operator: value } }]. For select/multiselect properties, values may be option slugs or option UUIDs.",
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

func handlePrismObjectsOrganizationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := micro.PrismObjectOrganizationNewParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Objects.Organizations.New(ctx, params, options...)
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
		Title:          "prism:objects:organizations create",
		Transform:      transform,
	})
}

func handlePrismObjectsOrganizationsBulkCreate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := micro.PrismObjectOrganizationBulkNewParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Objects.Organizations.BulkNew(ctx, params, options...)
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
		Title:          "prism:objects:organizations bulk-create",
		Transform:      transform,
	})
}

func handlePrismObjectsOrganizationsQuery(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := micro.PrismObjectOrganizationQueryParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Objects.Organizations.Query(ctx, params, options...)
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
		Title:          "prism:objects:organizations query",
		Transform:      transform,
	})
}
