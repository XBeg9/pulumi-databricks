// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks.Aws
{
    public partial class AwsS3Mount : Pulumi.CustomResource
    {
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        [Output("mountName")]
        public Output<string> MountName { get; private set; } = null!;

        [Output("s3BucketName")]
        public Output<string> S3BucketName { get; private set; } = null!;


        /// <summary>
        /// Create a AwsS3Mount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AwsS3Mount(string name, AwsS3MountArgs args, CustomResourceOptions? options = null)
            : base("databricks:aws/awsS3Mount:AwsS3Mount", name, args ?? new AwsS3MountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AwsS3Mount(string name, Input<string> id, AwsS3MountState? state = null, CustomResourceOptions? options = null)
            : base("databricks:aws/awsS3Mount:AwsS3Mount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AwsS3Mount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AwsS3Mount Get(string name, Input<string> id, AwsS3MountState? state = null, CustomResourceOptions? options = null)
        {
            return new AwsS3Mount(name, id, state, options);
        }
    }

    public sealed class AwsS3MountArgs : Pulumi.ResourceArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("mountName", required: true)]
        public Input<string> MountName { get; set; } = null!;

        [Input("s3BucketName", required: true)]
        public Input<string> S3BucketName { get; set; } = null!;

        public AwsS3MountArgs()
        {
        }
    }

    public sealed class AwsS3MountState : Pulumi.ResourceArgs
    {
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("mountName")]
        public Input<string>? MountName { get; set; }

        [Input("s3BucketName")]
        public Input<string>? S3BucketName { get; set; }

        public AwsS3MountState()
        {
        }
    }
}
