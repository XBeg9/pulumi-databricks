// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks.Inputs
{

    public sealed class InstancePoolAwsAttributesGetArgs : Pulumi.ResourceArgs
    {
        [Input("availability")]
        public Input<string>? Availability { get; set; }

        [Input("spotBidPricePercent")]
        public Input<int>? SpotBidPricePercent { get; set; }

        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public InstancePoolAwsAttributesGetArgs()
        {
        }
    }
}
