// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks
{
    public partial class SecretScope : Pulumi.CustomResource
    {
        [Output("backendType")]
        public Output<string> BackendType { get; private set; } = null!;

        [Output("initialManagePrincipal")]
        public Output<string?> InitialManagePrincipal { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a SecretScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretScope(string name, SecretScopeArgs? args = null, CustomResourceOptions? options = null)
            : base("databricks:index/secretScope:SecretScope", name, args ?? new SecretScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretScope(string name, Input<string> id, SecretScopeState? state = null, CustomResourceOptions? options = null)
            : base("databricks:index/secretScope:SecretScope", name, state, MakeResourceOptions(options, id))
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
        /// <summary>
        /// Get an existing SecretScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretScope Get(string name, Input<string> id, SecretScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretScope(name, id, state, options);
        }
    }

    public sealed class SecretScopeArgs : Pulumi.ResourceArgs
    {
        [Input("initialManagePrincipal")]
        public Input<string>? InitialManagePrincipal { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public SecretScopeArgs()
        {
        }
    }

    public sealed class SecretScopeState : Pulumi.ResourceArgs
    {
        [Input("backendType")]
        public Input<string>? BackendType { get; set; }

        [Input("initialManagePrincipal")]
        public Input<string>? InitialManagePrincipal { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public SecretScopeState()
        {
        }
    }
}
