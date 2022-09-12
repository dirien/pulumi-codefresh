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

    public sealed class PipelineSpecOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean for the Settings under pending approval: `When build enters "Pending Approval" state, volume should`:
        /// * Default (attribute not specified): "Use Setting accounts"
        /// * true: "Remain (build remains active)"
        /// * false: "Be removed"
        /// </summary>
        [Input("keepPvcsForPendingApproval")]
        public Input<bool>? KeepPvcsForPendingApproval { get; set; }

        /// <summary>
        /// Boolean for the Settings under pending approval: `Pipeline concurrency policy: Builds on "Pending Approval" state should be`:
        /// * Default (attribute not specified): "Use Setting accounts"
        /// * true: "Included in concurrency"
        /// * false: "Not included in concurrency"
        /// </summary>
        [Input("pendingApprovalConcurrencyApplied")]
        public Input<bool>? PendingApprovalConcurrencyApplied { get; set; }

        public PipelineSpecOptionsGetArgs()
        {
        }
        public static new PipelineSpecOptionsGetArgs Empty => new PipelineSpecOptionsGetArgs();
    }
}
