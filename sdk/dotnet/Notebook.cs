// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks
{
    public partial class Notebook : Pulumi.CustomResource
    {
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        [Output("language")]
        public Output<string?> Language { get; private set; } = null!;

        [Output("mkdirs")]
        public Output<bool?> Mkdirs { get; private set; } = null!;

        [Output("objectId")]
        public Output<int> ObjectId { get; private set; } = null!;

        [Output("objectType")]
        public Output<string> ObjectType { get; private set; } = null!;

        [Output("overwrite")]
        public Output<bool?> Overwrite { get; private set; } = null!;

        [Output("path")]
        public Output<string> Path { get; private set; } = null!;


        /// <summary>
        /// Create a Notebook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notebook(string name, NotebookArgs args, CustomResourceOptions? options = null)
            : base("databricks:index/notebook:Notebook", name, args ?? new NotebookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Notebook(string name, Input<string> id, NotebookState? state = null, CustomResourceOptions? options = null)
            : base("databricks:index/notebook:Notebook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Notebook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notebook Get(string name, Input<string> id, NotebookState? state = null, CustomResourceOptions? options = null)
        {
            return new Notebook(name, id, state, options);
        }
    }

    public sealed class NotebookArgs : Pulumi.ResourceArgs
    {
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("language")]
        public Input<string>? Language { get; set; }

        [Input("mkdirs")]
        public Input<bool>? Mkdirs { get; set; }

        [Input("overwrite")]
        public Input<bool>? Overwrite { get; set; }

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public NotebookArgs()
        {
        }
    }

    public sealed class NotebookState : Pulumi.ResourceArgs
    {
        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("language")]
        public Input<string>? Language { get; set; }

        [Input("mkdirs")]
        public Input<bool>? Mkdirs { get; set; }

        [Input("objectId")]
        public Input<int>? ObjectId { get; set; }

        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

        [Input("overwrite")]
        public Input<bool>? Overwrite { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        public NotebookState()
        {
        }
    }
}