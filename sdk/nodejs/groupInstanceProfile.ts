// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class GroupInstanceProfile extends pulumi.CustomResource {
    /**
     * Get an existing GroupInstanceProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupInstanceProfileState, opts?: pulumi.CustomResourceOptions): GroupInstanceProfile {
        return new GroupInstanceProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'databricks:index/groupInstanceProfile:GroupInstanceProfile';

    /**
     * Returns true if the given object is an instance of GroupInstanceProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupInstanceProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupInstanceProfile.__pulumiType;
    }

    public readonly groupId!: pulumi.Output<string>;
    public readonly instanceProfileId!: pulumi.Output<string>;

    /**
     * Create a GroupInstanceProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupInstanceProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupInstanceProfileArgs | GroupInstanceProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupInstanceProfileState | undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["instanceProfileId"] = state ? state.instanceProfileId : undefined;
        } else {
            const args = argsOrState as GroupInstanceProfileArgs | undefined;
            if (!args || args.groupId === undefined) {
                throw new Error("Missing required property 'groupId'");
            }
            if (!args || args.instanceProfileId === undefined) {
                throw new Error("Missing required property 'instanceProfileId'");
            }
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["instanceProfileId"] = args ? args.instanceProfileId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupInstanceProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupInstanceProfile resources.
 */
export interface GroupInstanceProfileState {
    readonly groupId?: pulumi.Input<string>;
    readonly instanceProfileId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupInstanceProfile resource.
 */
export interface GroupInstanceProfileArgs {
    readonly groupId: pulumi.Input<string>;
    readonly instanceProfileId: pulumi.Input<string>;
}