// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # User resource
 *
 * Use this resource to create a new user.
 *
 * ## Example usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as codefresh from "@pulumiverse/codefresh";
 *
 * const test = new codefresh.Account("test", {
 *     limits: [{
 *         collaborators: 25,
 *         dataRetentionWeeks: 5,
 *     }],
 *     builds: [{
 *         parallel: 2,
 *     }],
 * });
 * const _new = new codefresh.User("new", {
 *     email: "<EMAIL>",
 *     userName: "<USER>",
 *     activate: true,
 *     roles: [
 *         "Admin",
 *         "User",
 *     ],
 *     logins: [
 *         {
 *             idpId: data.codefresh_idps.idp_azure.id,
 *             sso: true,
 *         },
 *         {
 *             idpId: data.codefresh_idps.local.id,
 *         },
 *     ],
 *     personal: {
 *         firstName: "John",
 *         lastName: "Smith",
 *     },
 *     accounts: [
 *         test.id,
 *         "59009117c102763beda7ce71",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import codefresh:index/user:User new xxxxxxxxxxxxxxxxxxx
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'codefresh:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * A list of user roles. Possible values - `Admin`, `User`.
     */
    public readonly accounts!: pulumi.Output<string[]>;
    /**
     * Boolean. Activate the new use or not. If a new user is not activate, it'll be left pending.
     */
    public readonly activate!: pulumi.Output<boolean | undefined>;
    /**
     * A new user email.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * A collection of `login` blocks as documented below.
     */
    public readonly logins!: pulumi.Output<outputs.UserLogin[] | undefined>;
    /**
     * A collection of `personal` blocks as documented below.
     */
    public readonly personal!: pulumi.Output<outputs.UserPersonal | undefined>;
    public readonly roles!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly shortProfiles!: pulumi.Output<outputs.UserShortProfile[]>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The new user name.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["accounts"] = state ? state.accounts : undefined;
            resourceInputs["activate"] = state ? state.activate : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["logins"] = state ? state.logins : undefined;
            resourceInputs["personal"] = state ? state.personal : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["shortProfiles"] = state ? state.shortProfiles : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.accounts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accounts'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["accounts"] = args ? args.accounts : undefined;
            resourceInputs["activate"] = args ? args.activate : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["logins"] = args ? args.logins : undefined;
            resourceInputs["personal"] = args ? args.personal : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["shortProfiles"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * A list of user roles. Possible values - `Admin`, `User`.
     */
    accounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean. Activate the new use or not. If a new user is not activate, it'll be left pending.
     */
    activate?: pulumi.Input<boolean>;
    /**
     * A new user email.
     */
    email?: pulumi.Input<string>;
    /**
     * A collection of `login` blocks as documented below.
     */
    logins?: pulumi.Input<pulumi.Input<inputs.UserLogin>[]>;
    /**
     * A collection of `personal` blocks as documented below.
     */
    personal?: pulumi.Input<inputs.UserPersonal>;
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    shortProfiles?: pulumi.Input<pulumi.Input<inputs.UserShortProfile>[]>;
    status?: pulumi.Input<string>;
    /**
     * The new user name.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * A list of user roles. Possible values - `Admin`, `User`.
     */
    accounts: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean. Activate the new use or not. If a new user is not activate, it'll be left pending.
     */
    activate?: pulumi.Input<boolean>;
    /**
     * A new user email.
     */
    email: pulumi.Input<string>;
    /**
     * A collection of `login` blocks as documented below.
     */
    logins?: pulumi.Input<pulumi.Input<inputs.UserLogin>[]>;
    /**
     * A collection of `personal` blocks as documented below.
     */
    personal?: pulumi.Input<inputs.UserPersonal>;
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The new user name.
     */
    userName: pulumi.Input<string>;
}