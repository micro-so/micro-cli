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

var contactsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create Contact",
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
		&requestflag.Flag[any]{
			Name:     "default",
			BodyPath: "default",
		},
		&requestflag.Flag[any]{
			Name:     "extended",
			BodyPath: "extended",
		},
	},
	Action:          handleContactsCreate,
	HideHelpCommand: true,
}

var contactsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update Contact",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "contact-id",
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
		&requestflag.Flag[any]{
			Name:     "default",
			BodyPath: "default",
		},
		&requestflag.Flag[any]{
			Name:     "extended",
			BodyPath: "extended",
		},
	},
	Action:          handleContactsUpdate,
	HideHelpCommand: true,
}

var contactsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Contacts",
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
	Action:          handleContactsList,
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
			Usage:      "Filters as [{ slug: { operator: value } }]",
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

var contactsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Contact",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Required: true,
		},
	},
	Action:          handleContactsDelete,
	HideHelpCommand: true,
}

var contactsImport = requestflag.WithInnerFlags(cli.Command{
	Name:    "import",
	Usage:   "Import Contacts",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "team-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "object",
			Usage:    "Array of objects to import with their property values",
			Required: true,
			BodyPath: "objects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
	},
	Action:          handleContactsImport,
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
		&requestflag.InnerFlag[any]{
			Name:       "object.default",
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
			Usage:      "Property definition ID to deduplicate on",
			InnerField: "dedupe_by",
		},
		&requestflag.InnerFlag[string]{
			Name:       "options.dedupe-type",
			Usage:      "Type of the deduplication property",
			InnerField: "dedupe_type",
		},
	},
})

func handleContactsCreate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.ContactNewParams{
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
	_, err = client.Contacts.New(ctx, params, options...)
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
		Title:          "contacts create",
		Transform:      transform,
	})
}

func handleContactsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("contact-id") && len(unusedArgs) > 0 {
		cmd.Set("contact-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.ContactUpdateParams{
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

	return client.Contacts.Update(
		ctx,
		cmd.Value("contact-id").(string),
		params,
		options...,
	)
}

func handleContactsList(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.ContactListParams{
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
	_, err = client.Contacts.List(ctx, params, options...)
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
		Title:          "contacts list",
		Transform:      transform,
	})
}

func handleContactsDelete(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("contact-id") && len(unusedArgs) > 0 {
		cmd.Set("contact-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.ContactDeleteParams{
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

	return client.Contacts.Delete(
		ctx,
		cmd.Value("contact-id").(string),
		params,
		options...,
	)
}

func handleContactsImport(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.ContactImportParams{
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
	_, err = client.Contacts.Import(ctx, params, options...)
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
		Title:          "contacts import",
		Transform:      transform,
	})
}
