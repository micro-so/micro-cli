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

var prismMetadataList = cli.Command{
	Name:    "list",
	Usage:   "Get metadata properties by object type",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "team-id",
			Required:  true,
			PathParam: "teamId",
		},
		&requestflag.Flag[string]{
			Name:      "object-type",
			Usage:     `Allowed values: "deal", "identity", "ai_chat_thread", "ai_chat_message", "document", "action", "event", "organization", "contact".`,
			Required:  true,
			PathParam: "objectType",
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
	Action:          handlePrismMetadataList,
	HideHelpCommand: true,
}

func handlePrismMetadataList(ctx context.Context, cmd *cli.Command) error {
	client := micro.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("object-type") && len(unusedArgs) > 0 {
		cmd.Set("object-type", unusedArgs[0])
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

	params := micro.PrismMetadataListParams{
		TeamID: micro.F(cmd.Value("team-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Prism.Metadata.List(
		ctx,
		micro.PrismMetadataListParamsObjectType(cmd.Value("object-type").(string)),
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
		Title:          "prism:metadata list",
		Transform:      transform,
	})
}
