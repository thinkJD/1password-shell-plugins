package jira

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func AtlassianCLI() schema.Executable {
	return schema.Executable{
		Name:    "Atlassian Jira CLI",
		Runs:    []string{"jira"},
		DocsURL: sdk.URL("https://github.com/ankitpokhrel/jira-cli"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.AuthToken,
			},
		},
	}
}
