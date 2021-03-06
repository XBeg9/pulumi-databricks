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
    public sealed class InstancePoolDiskSpec
    {
        public readonly string? AzureDiskVolumeType;
        public readonly int? DiskCount;
        public readonly int? DiskSize;
        public readonly string? EbsVolumeType;

        [OutputConstructor]
        private InstancePoolDiskSpec(
            string? azureDiskVolumeType,

            int? diskCount,

            int? diskSize,

            string? ebsVolumeType)
        {
            AzureDiskVolumeType = azureDiskVolumeType;
            DiskCount = diskCount;
            DiskSize = diskSize;
            EbsVolumeType = ebsVolumeType;
        }
    }
}
