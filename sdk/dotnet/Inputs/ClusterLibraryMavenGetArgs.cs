// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Databricks.Inputs
{

    public sealed class ClusterLibraryMavenGetArgs : Pulumi.ResourceArgs
    {
        [Input("coordinates")]
        public Input<string>? Coordinates { get; set; }

        [Input("exclusions")]
        private InputList<string>? _exclusions;
        public InputList<string> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<string>());
            set => _exclusions = value;
        }

        [Input("messages")]
        private InputList<string>? _messages;
        public InputList<string> Messages
        {
            get => _messages ?? (_messages = new InputList<string>());
            set => _messages = value;
        }

        [Input("repo")]
        public Input<string>? Repo { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public ClusterLibraryMavenGetArgs()
        {
        }
    }
}
