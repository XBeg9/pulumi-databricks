// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SecretScope extends pulumi.CustomResource {
    /**
     * Get an existing SecretScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretScopeState, opts?: pulumi.CustomResourceOptions): SecretScope {
        return new SecretScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'databricks:index/secretScope:SecretScope';

    /**
     * Returns true if the given object is an instance of SecretScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretScope.__pulumiType;
    }

    public /*out*/ readonly backendType!: pulumi.Output<string>;
    public readonly initialManagePrincipal!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a SecretScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretScopeArgs | SecretScopeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretScopeState | undefined;
            inputs["backendType"] = state ? state.backendType : undefined;
            inputs["initialManagePrincipal"] = state ? state.initialManagePrincipal : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SecretScopeArgs | undefined;
            inputs["initialManagePrincipal"] = args ? args.initialManagePrincipal : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["backendType"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretScope resources.
 */
export interface SecretScopeState {
    readonly backendType?: pulumi.Input<string>;
    readonly initialManagePrincipal?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretScope resource.
 */
export interface SecretScopeArgs {
    readonly initialManagePrincipal?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
}