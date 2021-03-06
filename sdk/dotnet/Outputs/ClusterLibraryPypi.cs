// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks.Outputs
{

    [OutputType]
    public sealed class ClusterLibraryPypi
    {
        public readonly ImmutableArray<string> Messages;
        public readonly string Package;
        public readonly string? Repo;
        public readonly string? Status;

        [OutputConstructor]
        private ClusterLibraryPypi(
            ImmutableArray<string> messages,

            string package,

            string? repo,

            string? status)
        {
            Messages = messages;
            Package = package;
            Repo = repo;
            Status = status;
        }
    }
}
