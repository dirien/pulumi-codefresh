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

    public sealed class UserShortProfileGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The new user name.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public UserShortProfileGetArgs()
        {
        }
        public static new UserShortProfileGetArgs Empty => new UserShortProfileGetArgs();
    }
}
