// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codefresh

import (
	"fmt"
	"path/filepath"

	"github.com/codefresh-io/terraform-provider-codefresh/codefresh"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumiverse/pulumi-codefresh/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "codefresh"
	// modules:
	mainMod = "index" // the codefresh module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(codefresh.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "codefresh",
		DisplayName:       "Codefresh",
		Publisher:         "Pulumiverse",
		LogoURL:           "https://avatars3.githubusercontent.com/codefresh-io",
		PluginDownloadURL: "github://api.github.com/dirien/pulumi-codefresh",
		Description:       "A Pulumi package for creating and managing codefresh cloud resources.",
		Keywords:          []string{"pulumi", "codefresh", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://www.pulumi.com",
		Repository:        "https://github.com/pulumiverse/pulumi-qovery",
		GitHubOrg:         "codefresh-io",
		Config:            map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"codefresh_account":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Account")},
			"codefresh_account-admins": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AccountAdmins")},
			"codefresh_api-key":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApiKey")},
			"codefresh_context":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Context")},
			"codefresh_idp-accounts":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IdpAccounts")},
			"codefresh_permissions":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Permissions")},
			"codefresh_pipeline":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Pipeline")},
			"codefresh_project":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project")},
			"codefresh_registry":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Registry")},
			"codefresh_step-types":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StepTypes")},
			"codefresh_team":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team")},
			"codefresh_user":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/codefresh",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_codefresh",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumiverse",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
