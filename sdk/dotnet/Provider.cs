// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks
{
    /// <summary>
    /// The provider type for the databricks package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("databricks", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        [Input("azureAuth", json: true)]
        public Input<Inputs.ProviderAzureAuthArgs>? AzureAuth { get; set; }

        [Input("basicAuth", json: true)]
        public Input<Inputs.ProviderBasicAuthArgs>? BasicAuth { get; set; }

        /// <summary>
        /// Location of the Databricks CLI credentials file, that is created by `databricks configure --token` command. By default,
        /// it is located in ~/.databrickscfg. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for
        /// docs. Config file credetials will only be used when host/token are not provided.
        /// </summary>
        [Input("configFile")]
        public Input<string>? ConfigFile { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Connection profile specified within ~/.databrickscfg. Please check
        /// https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        [Input("token")]
        public Input<string>? Token { get; set; }

        public ProviderArgs()
        {
            Host = Utilities.GetEnv("DATABRICKS_HOST");
            Token = Utilities.GetEnv("DATABRICKS_TOKEN");
        }
    }
}
