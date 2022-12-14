// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Project Resource
 *
 * The top-level concept in Codefresh. You can create projects to group pipelines that are related. In most cases a single project will be a single application (that itself contains many micro-services). You are free to use projects as you see fit. For example, you could create a project for a specific Kubernetes cluster or a specific team/department.
 * More about pipeline concepts see in the [official documentation](https://codefresh.io/docs/docs/configure-ci-cd-pipeline/pipelines/#pipeline-concepts).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as codefresh from "@pulumi/codefresh";
 *
 * const test = new codefresh.Project("test", {
 *     tags: [
 *         "production",
 *         "docker",
 *     ],
 *     variables: {
 *         go_version: "1.13",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import codefresh:index/project:Project test xxxxxxxxxxxxxxxxxxx
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'codefresh:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * The display name for the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of tags to mark a project for easy management and access control.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * project variables.
     */
    public readonly variables!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The display name for the project.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of tags to mark a project for easy management and access control.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * project variables.
     */
    variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The display name for the project.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of tags to mark a project for easy management and access control.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * project variables.
     */
    variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
