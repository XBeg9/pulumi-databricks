// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks
{
    public static class GetDbfsFilePaths
    {
        public static Task<GetDbfsFilePathsResult> InvokeAsync(GetDbfsFilePathsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDbfsFilePathsResult>("databricks:index/getDbfsFilePaths:getDbfsFilePaths", args ?? new GetDbfsFilePathsArgs(), options.WithVersion());
    }


    public sealed class GetDbfsFilePathsArgs : Pulumi.InvokeArgs
    {
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        [Input("recursive", required: true)]
        public bool Recursive { get; set; }

        public GetDbfsFilePathsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDbfsFilePathsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Path;
        public readonly ImmutableArray<Outputs.GetDbfsFilePathsPathListResult> PathLists;
        public readonly bool Recursive;

        [OutputConstructor]
        private GetDbfsFilePathsResult(
            string id,

            string path,

            ImmutableArray<Outputs.GetDbfsFilePathsPathListResult> pathLists,

            bool recursive)
        {
            Id = id;
            Path = path;
            PathLists = pathLists;
            Recursive = recursive;
        }
    }
}
