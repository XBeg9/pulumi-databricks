// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class MwsNetworks extends pulumi.CustomResource {
    /**
     * Get an existing MwsNetworks resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MwsNetworksState, opts?: pulumi.CustomResourceOptions): MwsNetworks {
        return new MwsNetworks(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'databricks:index/mwsNetworks:MwsNetworks';

    /**
     * Returns true if the given object is an instance of MwsNetworks.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MwsNetworks {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MwsNetworks.__pulumiType;
    }

    public readonly accountId!: pulumi.Output<string>;
    public /*out*/ readonly creationTime!: pulumi.Output<number>;
    public readonly errorMessages!: pulumi.Output<outputs.MwsNetworksErrorMessage[]>;
    public /*out*/ readonly networkId!: pulumi.Output<string>;
    public readonly networkName!: pulumi.Output<string>;
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    public readonly subnetIds!: pulumi.Output<string[]>;
    public readonly vpcId!: pulumi.Output<string>;
    public /*out*/ readonly vpcStatus!: pulumi.Output<string>;
    public /*out*/ readonly workspaceId!: pulumi.Output<number>;

    /**
     * Create a MwsNetworks resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MwsNetworksArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MwsNetworksArgs | MwsNetworksState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MwsNetworksState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["errorMessages"] = state ? state.errorMessages : undefined;
            inputs["networkId"] = state ? state.networkId : undefined;
            inputs["networkName"] = state ? state.networkName : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["vpcStatus"] = state ? state.vpcStatus : undefined;
            inputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as MwsNetworksArgs | undefined;
            if (!args || args.accountId === undefined) {
                throw new Error("Missing required property 'accountId'");
            }
            if (!args || args.networkName === undefined) {
                throw new Error("Missing required property 'networkName'");
            }
            if (!args || args.securityGroupIds === undefined) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            if (!args || args.subnetIds === undefined) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["errorMessages"] = args ? args.errorMessages : undefined;
            inputs["networkName"] = args ? args.networkName : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["networkId"] = undefined /*out*/;
            inputs["vpcStatus"] = undefined /*out*/;
            inputs["workspaceId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MwsNetworks.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MwsNetworks resources.
 */
export interface MwsNetworksState {
    readonly accountId?: pulumi.Input<string>;
    readonly creationTime?: pulumi.Input<number>;
    readonly errorMessages?: pulumi.Input<pulumi.Input<inputs.MwsNetworksErrorMessage>[]>;
    readonly networkId?: pulumi.Input<string>;
    readonly networkName?: pulumi.Input<string>;
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly vpcId?: pulumi.Input<string>;
    readonly vpcStatus?: pulumi.Input<string>;
    readonly workspaceId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a MwsNetworks resource.
 */
export interface MwsNetworksArgs {
    readonly accountId: pulumi.Input<string>;
    readonly errorMessages?: pulumi.Input<pulumi.Input<inputs.MwsNetworksErrorMessage>[]>;
    readonly networkName: pulumi.Input<string>;
    readonly securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    readonly subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    readonly vpcId: pulumi.Input<string>;
}
