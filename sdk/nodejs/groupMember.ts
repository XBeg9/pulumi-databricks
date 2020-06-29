// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class GroupMember extends pulumi.CustomResource {
    /**
     * Get an existing GroupMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMemberState, opts?: pulumi.CustomResourceOptions): GroupMember {
        return new GroupMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'databricks:index/groupMember:GroupMember';

    /**
     * Returns true if the given object is an instance of GroupMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMember.__pulumiType;
    }

    public readonly groupId!: pulumi.Output<string>;
    public readonly memberId!: pulumi.Output<string>;

    /**
     * Create a GroupMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMemberArgs | GroupMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupMemberState | undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["memberId"] = state ? state.memberId : undefined;
        } else {
            const args = argsOrState as GroupMemberArgs | undefined;
            if (!args || args.groupId === undefined) {
                throw new Error("Missing required property 'groupId'");
            }
            if (!args || args.memberId === undefined) {
                throw new Error("Missing required property 'memberId'");
            }
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["memberId"] = args ? args.memberId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMember resources.
 */
export interface GroupMemberState {
    readonly groupId?: pulumi.Input<string>;
    readonly memberId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMember resource.
 */
export interface GroupMemberArgs {
    readonly groupId: pulumi.Input<string>;
    readonly memberId: pulumi.Input<string>;
}