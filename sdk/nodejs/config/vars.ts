// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("databricks");

export let azureAuth: { azureRegion: string, clientId: string, clientSecret: string, managedResourceGroup: string, resourceGroup: string, subscriptionId: string, tenantId: string, workspaceName: string } | undefined = __config.getObject<{ azureRegion: string, clientId: string, clientSecret: string, managedResourceGroup: string, resourceGroup: string, subscriptionId: string, tenantId: string, workspaceName: string }>("azureAuth");
export let basicAuth: { password: string, username: string } | undefined = __config.getObject<{ password: string, username: string }>("basicAuth");
/**
 * Location of the Databricks CLI credentials file, that is created by `databricks configure --token` command. By default,
 * it is located in ~/.databrickscfg. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for
 * docs. Config file credetials will only be used when host/token are not provided.
 */
export let configFile: string | undefined = __config.get("configFile");
export let host: string | undefined = __config.get("host");
/**
 * Connection profile specified within ~/.databrickscfg. Please check
 * https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.
 */
export let profile: string | undefined = __config.get("profile");
export let token: string | undefined = __config.get("token");