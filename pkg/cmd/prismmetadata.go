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

var prismMetadataProperties = cli.Command{
	Name:    "properties",
	Usage:   "Get metadata properties by object type",
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
		&requestflag.Flag[bool]{
			Name:      "autofill",
			QueryPath: "autofill",
		},
		&requestflag.Flag[string]{
			Name:      "crm-id",
			QueryPath: "crmId",
		},
		&requestflag.Flag[string]{
			Name:      "term",
			QueryPath: "term",
		},
	},
	Action:          handlePrismMetadataProperties,
	HideHelpCommand: true,
}

func handlePrismMetadataProperties(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("object-type") && len(unusedArgs) > 0 {
		cmd.Set("object-type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := micro.PrismMetadataPropertiesParams{
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

	return client.Prism.Metadata.Properties(
		ctx,
		micro.ObjectType(cmd.Value("object-type").(string)),
		params,
		options...,
	)
}
