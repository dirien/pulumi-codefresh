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
    public sealed class RegistrySpecGar
    {
        /// <summary>
        /// _(Required)_ String representing service account json file contents
        /// </summary>
        public readonly string Keyfile;
        /// <summary>
        /// _(Required)_ String representing one of the Google's gar locations
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// _(Optional)_ String. See the [docs](https://codefresh.io/docs/docs/integrations/docker-registries/#using-an-optional-repository-prefix).
        /// </summary>
        public readonly string? RepositoryPrefix;

        [OutputConstructor]
        private RegistrySpecGar(
            string keyfile,

            string location,

            string? repositoryPrefix)
        {
            Keyfile = keyfile;
            Location = location;
            RepositoryPrefix = repositoryPrefix;
        }
    }
}
