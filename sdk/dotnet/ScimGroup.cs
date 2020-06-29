// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks
{
    public partial class ScimGroup : Pulumi.CustomResource
    {
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        [Output("entitlements")]
        public Output<ImmutableArray<string>> Entitlements { get; private set; } = null!;

        [Output("inheritedRoles")]
        public Output<ImmutableArray<string>> InheritedRoles { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;


        /// <summary>
        /// Create a ScimGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScimGroup(string name, ScimGroupArgs args, CustomResourceOptions? options = null)
            : base("databricks:index/scimGroup:ScimGroup", name, args ?? new ScimGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScimGroup(string name, Input<string> id, ScimGroupState? state = null, CustomResourceOptions? options = null)
            : base("databricks:index/scimGroup:ScimGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScimGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScimGroup Get(string name, Input<string> id, ScimGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ScimGroup(name, id, state, options);
        }
    }

    public sealed class ScimGroupArgs : Pulumi.ResourceArgs
    {
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("entitlements")]
        private InputList<string>? _entitlements;
        public InputList<string> Entitlements
        {
            get => _entitlements ?? (_entitlements = new InputList<string>());
            set => _entitlements = value;
        }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public ScimGroupArgs()
        {
        }
    }

    public sealed class ScimGroupState : Pulumi.ResourceArgs
    {
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("entitlements")]
        private InputList<string>? _entitlements;
        public InputList<string> Entitlements
        {
            get => _entitlements ?? (_entitlements = new InputList<string>());
            set => _entitlements = value;
        }

        [Input("inheritedRoles")]
        private InputList<string>? _inheritedRoles;
        public InputList<string> InheritedRoles
        {
            get => _inheritedRoles ?? (_inheritedRoles = new InputList<string>());
            set => _inheritedRoles = value;
        }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public ScimGroupState()
        {
        }
    }
}