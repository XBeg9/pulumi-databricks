// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func GetAzureAuth(ctx *pulumi.Context) string {
	return config.Get(ctx, "databricks:azureAuth")
}
func GetBasicAuth(ctx *pulumi.Context) string {
	return config.Get(ctx, "databricks:basicAuth")
}

// Location of the Databricks CLI credentials file, that is created by `databricks configure --token` command. By default,
// it is located in ~/.databrickscfg. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for
// docs. Config file credetials will only be used when host/token are not provided.
func GetConfigFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "databricks:configFile")
}
func GetHost(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "databricks:host")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "DATABRICKS_HOST").(string)
}

// Connection profile specified within ~/.databrickscfg. Please check
// https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.
func GetProfile(ctx *pulumi.Context) string {
	return config.Get(ctx, "databricks:profile")
}
func GetToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "databricks:token")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "DATABRICKS_TOKEN").(string)
}
