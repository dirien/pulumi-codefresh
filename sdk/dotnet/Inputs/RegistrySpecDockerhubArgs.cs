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

    public sealed class RegistrySpecDockerhubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// _(Required, Sensitive)_ String.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// _(Required)_ String.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public RegistrySpecDockerhubArgs()
        {
        }
        public static new RegistrySpecDockerhubArgs Empty => new RegistrySpecDockerhubArgs();
    }
}
