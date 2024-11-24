package jira

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "jira",
		Platform: schema.PlatformInfo{
			Name:     "Atlassian",
			Homepage: sdk.URL("https://github.com/go-jira/jira"),
		},
		Credentials: []schema.CredentialType{
			AuthToken(),
		},
		Executables: []schema.Executable{
			AtlassianCLI(),
		},
	}
}
