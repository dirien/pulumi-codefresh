// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Codefresh.Outputs
{

    [OutputType]
    public sealed class PipelineSpecTerminationPolicy
    {
        /// <summary>
        /// A `on_create_branch` block as documented below.
        /// </summary>
        public readonly Outputs.PipelineSpecTerminationPolicyOnCreateBranch? OnCreateBranch;
        /// <summary>
        /// Boolean. Enables the policy `Once a build is terminated, terminate all child builds initiated from it`. Default false.
        /// </summary>
        public readonly bool? OnTerminateAnnotation;

        [OutputConstructor]
        private PipelineSpecTerminationPolicy(
            Outputs.PipelineSpecTerminationPolicyOnCreateBranch? onCreateBranch,

            bool? onTerminateAnnotation)
        {
            OnCreateBranch = onCreateBranch;
            OnTerminateAnnotation = onTerminateAnnotation;
        }
    }
}
