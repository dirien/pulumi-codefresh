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

    public sealed class ContextSpecStorages3DataGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("auth", required: true)]
        public Input<Inputs.ContextSpecStorages3DataAuthGetArgs> Auth { get; set; } = null!;

        public ContextSpecStorages3DataGetArgs()
        {
        }
        public static new ContextSpecStorages3DataGetArgs Empty => new ContextSpecStorages3DataGetArgs();
    }
}
