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
    public sealed class UserPersonal
    {
        /// <summary>
        /// .
        /// </summary>
        public readonly string? CompanyName;
        /// <summary>
        /// .
        /// </summary>
        public readonly string? Country;
        /// <summary>
        /// .
        /// </summary>
        public readonly string? FirstName;
        /// <summary>
        /// .
        /// </summary>
        public readonly string? LastName;
        /// <summary>
        /// .
        /// </summary>
        public readonly string? PhoneNumber;

        [OutputConstructor]
        private UserPersonal(
            string? companyName,

            string? country,

            string? firstName,

            string? lastName,

            string? phoneNumber)
        {
            CompanyName = companyName;
            Country = country;
            FirstName = firstName;
            LastName = lastName;
            PhoneNumber = phoneNumber;
        }
    }
}
