// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks
{
    public partial class MwsWorkspaces : Pulumi.CustomResource
    {
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        [Output("awsRegion")]
        public Output<string> AwsRegion { get; private set; } = null!;

        [Output("creationTime")]
        public Output<int> CreationTime { get; private set; } = null!;

        [Output("credentialsId")]
        public Output<string> CredentialsId { get; private set; } = null!;

        [Output("customerManagedKeyId")]
        public Output<string?> CustomerManagedKeyId { get; private set; } = null!;

        [Output("deploymentName")]
        public Output<string> DeploymentName { get; private set; } = null!;

        [Output("isNoPublicIpEnabled")]
        public Output<bool?> IsNoPublicIpEnabled { get; private set; } = null!;

        [Output("networkErrorMessages")]
        public Output<ImmutableArray<Outputs.MwsWorkspacesNetworkErrorMessage>> NetworkErrorMessages { get; private set; } = null!;

        [Output("networkId")]
        public Output<string?> NetworkId { get; private set; } = null!;

        [Output("storageConfigurationId")]
        public Output<string> StorageConfigurationId { get; private set; } = null!;

        [Output("verifyWorkspaceRunnning")]
        public Output<bool> VerifyWorkspaceRunnning { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<int> WorkspaceId { get; private set; } = null!;

        [Output("workspaceName")]
        public Output<string> WorkspaceName { get; private set; } = null!;

        [Output("workspaceStatus")]
        public Output<string> WorkspaceStatus { get; private set; } = null!;

        [Output("workspaceStatusMessage")]
        public Output<string> WorkspaceStatusMessage { get; private set; } = null!;

        [Output("workspaceUrl")]
        public Output<string> WorkspaceUrl { get; private set; } = null!;


        /// <summary>
        /// Create a MwsWorkspaces resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MwsWorkspaces(string name, MwsWorkspacesArgs args, CustomResourceOptions? options = null)
            : base("databricks:index/mwsWorkspaces:MwsWorkspaces", name, args ?? new MwsWorkspacesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MwsWorkspaces(string name, Input<string> id, MwsWorkspacesState? state = null, CustomResourceOptions? options = null)
            : base("databricks:index/mwsWorkspaces:MwsWorkspaces", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MwsWorkspaces resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MwsWorkspaces Get(string name, Input<string> id, MwsWorkspacesState? state = null, CustomResourceOptions? options = null)
        {
            return new MwsWorkspaces(name, id, state, options);
        }
    }

    public sealed class MwsWorkspacesArgs : Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("awsRegion", required: true)]
        public Input<string> AwsRegion { get; set; } = null!;

        [Input("credentialsId", required: true)]
        public Input<string> CredentialsId { get; set; } = null!;

        [Input("customerManagedKeyId")]
        public Input<string>? CustomerManagedKeyId { get; set; }

        [Input("deploymentName", required: true)]
        public Input<string> DeploymentName { get; set; } = null!;

        [Input("isNoPublicIpEnabled")]
        public Input<bool>? IsNoPublicIpEnabled { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("storageConfigurationId", required: true)]
        public Input<string> StorageConfigurationId { get; set; } = null!;

        [Input("verifyWorkspaceRunnning", required: true)]
        public Input<bool> VerifyWorkspaceRunnning { get; set; } = null!;

        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public MwsWorkspacesArgs()
        {
        }
    }

    public sealed class MwsWorkspacesState : Pulumi.ResourceArgs
    {
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        [Input("creationTime")]
        public Input<int>? CreationTime { get; set; }

        [Input("credentialsId")]
        public Input<string>? CredentialsId { get; set; }

        [Input("customerManagedKeyId")]
        public Input<string>? CustomerManagedKeyId { get; set; }

        [Input("deploymentName")]
        public Input<string>? DeploymentName { get; set; }

        [Input("isNoPublicIpEnabled")]
        public Input<bool>? IsNoPublicIpEnabled { get; set; }

        [Input("networkErrorMessages")]
        private InputList<Inputs.MwsWorkspacesNetworkErrorMessageGetArgs>? _networkErrorMessages;
        public InputList<Inputs.MwsWorkspacesNetworkErrorMessageGetArgs> NetworkErrorMessages
        {
            get => _networkErrorMessages ?? (_networkErrorMessages = new InputList<Inputs.MwsWorkspacesNetworkErrorMessageGetArgs>());
            set => _networkErrorMessages = value;
        }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("storageConfigurationId")]
        public Input<string>? StorageConfigurationId { get; set; }

        [Input("verifyWorkspaceRunnning")]
        public Input<bool>? VerifyWorkspaceRunnning { get; set; }

        [Input("workspaceId")]
        public Input<int>? WorkspaceId { get; set; }

        [Input("workspaceName")]
        public Input<string>? WorkspaceName { get; set; }

        [Input("workspaceStatus")]
        public Input<string>? WorkspaceStatus { get; set; }

        [Input("workspaceStatusMessage")]
        public Input<string>? WorkspaceStatusMessage { get; set; }

        [Input("workspaceUrl")]
        public Input<string>? WorkspaceUrl { get; set; }

        public MwsWorkspacesState()
        {
        }
    }
}
