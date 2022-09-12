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

    public sealed class AccountBuildGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("nodes")]
        public Input<int>? Nodes { get; set; }

        /// <summary>
        /// How many pipelines can be run in parallel.
        /// ` `node` - (Optional) Number of nodes.
        /// </summary>
        [Input("parallel", required: true)]
        public Input<int> Parallel { get; set; } = null!;

        public AccountBuildGetArgs()
        {
        }
        public static new AccountBuildGetArgs Empty => new AccountBuildGetArgs();
    }
}
