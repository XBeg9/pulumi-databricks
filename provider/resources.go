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

package databricks

import (
	"unicode"

	"github.com/databrickslabs/databricks-terraform/databricks"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "databricks"
	// modules:
	mainMod  = "index"
	awsMod   = "aws"
	azureMod = "azure"
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := databricks.Provider("1.0.0").(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "databricks",
		Description: "A Pulumi package for creating and managing databricks cloud resources.",
		Keywords:    []string{"pulumi", "databricks"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/XBeg9/pulumi-databricks",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			"host": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DATABRICKS_HOST"},
				},
			},
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DATABRICKS_TOKEN"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			"databricks_token":         {Tok: makeResource(mainMod, "Token")},
			"databricks_secret_scope":  {Tok: makeResource(mainMod, "SecretScope")},
			"databricks_secret":        {Tok: makeResource(mainMod, "Secret")},
			"databricks_secret_acl":    {Tok: makeResource(mainMod, "SecretAcl")},
			"databricks_permissions":   {Tok: makeResource(mainMod, "Permissions")},
			"databricks_instance_pool": {Tok: makeResource(mainMod, "InstancePool")},
			"databricks_scim_user":     {Tok: makeResource(mainMod, "ScimUser")},
			"databricks_scim_group":    {Tok: makeResource(mainMod, "ScimGroup")},

			// Scim Group is split into multiple components for flexibility to pick and choose
			"databricks_group":                  {Tok: makeResource(mainMod, "Group")},
			"databricks_group_instance_profile": {Tok: makeResource(mainMod, "GroupInstanceProfile")},
			"databricks_group_member":           {Tok: makeResource(mainMod, "GroupMember")},
			"databricks_notebook":               {Tok: makeResource(mainMod, "Notebook")},
			"databricks_cluster":                {Tok: makeResource(mainMod, "Cluster")},
			"databricks_cluster_policy":         {Tok: makeResource(mainMod, "ClusterPolicy")},
			"databricks_job":                    {Tok: makeResource(mainMod, "Job")},
			"databricks_dbfs_file":              {Tok: makeResource(mainMod, "DbfsFile")},
			"databricks_dbfs_file_sync":         {Tok: makeResource(mainMod, "DbfsFileSync")},
			"databricks_instance_profile":       {Tok: makeResource(mainMod, "InstanceProfile")},
			"databricks_aws_s3_mount":           {Tok: makeResource(awsMod, "AwsS3Mount")},
			"databricks_azure_blob_mount":       {Tok: makeResource(azureMod, "AzureBlobMount")},
			"databricks_azure_adls_gen1_mount":  {Tok: makeResource(azureMod, "AzureAdlsGen1Mount")},
			"databricks_azure_adls_gen2_mount":  {Tok: makeResource(azureMod, "AzureAdlsGen2Mount")},

			//	MWS (multiple workspaces) resources are only limited to AWS as azure already has a built in concept of MWS
			"databricks_mws_credentials":            {Tok: makeResource(awsMod, "MwsCredentials")},
			"databricks_mws_storage_configurations": {Tok: makeResource(awsMod, "MwsStorageConfigurations")},
			"databricks_mws_networks":               {Tok: makeResource(awsMod, "MwsNetworks")},
			"databricks_mws_workspaces":             {Tok: makeResource(awsMod, "MwsWorkspaces")},
		},
		//
		// "aws_acm_certificate": {
		// 	Tok: makeResource(mainMod, "Certificate"),
		// 	Fields: map[string]*tfbridge.SchemaInfo{
		// 		"tags": {Type: makeType(mainPkg, "Tags")},
		// 	},
		// },
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			"databricks_default_user_roles": {Tok: makeDataSource(mainMod, "getDefaultUserRoles")},
			"databricks_notebook":           {Tok: makeDataSource(mainMod, "getNotebook")},
			"databricks_notebook_paths":     {Tok: makeDataSource(mainMod, "getNotebookPaths")},
			"databricks_dbfs_file":          {Tok: makeDataSource(mainMod, "getDbfsFile")},
			"databricks_dbfs_file_paths":    {Tok: makeDataSource(mainMod, "getDbfsFilePaths")},
			"databricks_zones":              {Tok: makeDataSource(mainMod, "getZones")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// AsyncDataSources: true,
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
