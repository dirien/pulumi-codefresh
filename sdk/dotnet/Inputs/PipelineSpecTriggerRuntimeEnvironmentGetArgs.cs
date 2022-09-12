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

    public sealed class PipelineSpecTriggerRuntimeEnvironmentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A required amount of CPU.
        /// </summary>
        [Input("cpu")]
        public Input<string>? Cpu { get; set; }

        /// <summary>
        /// A pipeline shared storage.
        /// </summary>
        [Input("dindStorage")]
        public Input<string>? DindStorage { get; set; }

        /// <summary>
        /// A required amount of memory.
        /// </summary>
        [Input("memory")]
        public Input<string>? Memory { get; set; }

        /// <summary>
        /// A name of runtime.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PipelineSpecTriggerRuntimeEnvironmentGetArgs()
        {
        }
        public static new PipelineSpecTriggerRuntimeEnvironmentGetArgs Empty => new PipelineSpecTriggerRuntimeEnvironmentGetArgs();
    }
}
