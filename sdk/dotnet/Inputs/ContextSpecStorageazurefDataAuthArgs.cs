// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Codefresh.Inputs
{

    public sealed class ContextSpecStorageazurefDataAuthArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountKey", required: true)]
        public Input<string> AccountKey { get; set; } = null!;

        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ContextSpecStorageazurefDataAuthArgs()
        {
        }
        public static new ContextSpecStorageazurefDataAuthArgs Empty => new ContextSpecStorageazurefDataAuthArgs();
    }
}
