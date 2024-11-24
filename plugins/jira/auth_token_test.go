package jira

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestAuthTokenProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, AuthToken().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Token: "t%/{pj41\w/PBuy2{]$&#43;B1KGi)/f5~-ECOye/Gq.;I#J1AIX$pSdMuT`&lt;&#34;&gt;PkgfjZ&amp;Pys9p&amp;`rY}](&amp;QflaG`:{fJ52&lt;2&gt;iC5:yc/gA#~&#34;Ve2UHm=LGeE?4yt$rRdgZ4G:-DyboP[1LRW&amp;Aapa~X2t&#43;s5jd$JWlnx96zuZRXZ1Iq}jQ9hD.#fv$u*EXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"JIRA_API_TOKEN": "t%/{pj41\w/PBuy2{]$&#43;B1KGi)/f5~-ECOye/Gq.;I#J1AIX$pSdMuT`&lt;&#34;&gt;PkgfjZ&amp;Pys9p&amp;`rY}](&amp;QflaG`:{fJ52&lt;2&gt;iC5:yc/gA#~&#34;Ve2UHm=LGeE?4yt$rRdgZ4G:-DyboP[1LRW&amp;Aapa~X2t&#43;s5jd$JWlnx96zuZRXZ1Iq}jQ9hD.#fv$u*EXAMPLE",
				},
			},
		},
	})
}

func TestAuthTokenImporter(t *testing.T) {
	plugintest.TestImporter(t, AuthToken().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"JIRA_API_TOKEN": "t%/{pj41\w/PBuy2{]$&#43;B1KGi)/f5~-ECOye/Gq.;I#J1AIX$pSdMuT`&lt;&#34;&gt;PkgfjZ&amp;Pys9p&amp;`rY}](&amp;QflaG`:{fJ52&lt;2&gt;iC5:yc/gA#~&#34;Ve2UHm=LGeE?4yt$rRdgZ4G:-DyboP[1LRW&amp;Aapa~X2t&#43;s5jd$JWlnx96zuZRXZ1Iq}jQ9hD.#fv$u*EXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Token: "t%/{pj41\w/PBuy2{]$&#43;B1KGi)/f5~-ECOye/Gq.;I#J1AIX$pSdMuT`&lt;&#34;&gt;PkgfjZ&amp;Pys9p&amp;`rY}](&amp;QflaG`:{fJ52&lt;2&gt;iC5:yc/gA#~&#34;Ve2UHm=LGeE?4yt$rRdgZ4G:-DyboP[1LRW&amp;Aapa~X2t&#43;s5jd$JWlnx96zuZRXZ1Iq}jQ9hD.#fv$u*EXAMPLE",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in jira-cli/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "t%/{pj41\w/PBuy2{]$&#43;B1KGi)/f5~-ECOye/Gq.;I#J1AIX$pSdMuT`&lt;&#34;&gt;PkgfjZ&amp;Pys9p&amp;`rY}](&amp;QflaG`:{fJ52&lt;2&gt;iC5:yc/gA#~&#34;Ve2UHm=LGeE?4yt$rRdgZ4G:-DyboP[1LRW&amp;Aapa~X2t&#43;s5jd$JWlnx96zuZRXZ1Iq}jQ9hD.#fv$u*EXAMPLE",
			// 		},
			// 	},
			},
		},
	})
}
