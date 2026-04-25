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

var organizationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create Organization",
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
	Action:          handleOrganizationsCreate,
	HideHelpCommand: true,
}

var organizationsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update Organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "organization-id",
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
	Action:          handleOrganizationsUpdate,
	HideHelpCommand: true,
}

var organizationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Organizations",
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
	Action:          handleOrganizationsList,
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

var organizationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "organization-id",
			Required: true,
		},
	},
	Action:          handleOrganizationsDelete,
	HideHelpCommand: true,
}

var organizationsImport = requestflag.WithInnerFlags(cli.Command{
	Name:    "import",
	Usage:   "Import Organizations",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
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
	Action:          handleOrganizationsImport,
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
			Usage:      "Properties keyed by property slug. Values can be strings, numbers, booleans, arrays, or null.",
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

func handleOrganizationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.OrganizationNewParams{
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
	_, err = client.Organizations.New(ctx, params, options...)
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
		Title:          "organizations create",
		Transform:      transform,
	})
}

func handleOrganizationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.OrganizationUpdateParams{
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

	return client.Organizations.Update(
		ctx,
		cmd.Value("organization-id").(string),
		params,
		options...,
	)
}

func handleOrganizationsList(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.OrganizationListParams{
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
	_, err = client.Organizations.List(ctx, params, options...)
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
		Title:          "organizations list",
		Transform:      transform,
	})
}

func handleOrganizationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.OrganizationDeleteParams{
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

	return client.Organizations.Delete(
		ctx,
		cmd.Value("organization-id").(string),
		params,
		options...,
	)
}

func handleOrganizationsImport(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.OrganizationImportParams{
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
	_, err = client.Organizations.Import(ctx, params, options...)
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
		Title:          "organizations import",
		Transform:      transform,
	})
}
