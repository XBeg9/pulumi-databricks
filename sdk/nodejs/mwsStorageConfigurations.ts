// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class MwsStorageConfigurations extends pulumi.CustomResource {
    /**
     * Get an existing MwsStorageConfigurations resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MwsStorageConfigurationsState, opts?: pulumi.CustomResourceOptions): MwsStorageConfigurations {
        return new MwsStorageConfigurations(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'databricks:index/mwsStorageConfigurations:MwsStorageConfigurations';

    /**
     * Returns true if the given object is an instance of MwsStorageConfigurations.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MwsStorageConfigurations {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MwsStorageConfigurations.__pulumiType;
    }

    public readonly accountId!: pulumi.Output<string>;
    public readonly bucketName!: pulumi.Output<string>;
    public /*out*/ readonly creationTime!: pulumi.Output<number>;
    public /*out*/ readonly storageConfigurationId!: pulumi.Output<string>;
    public readonly storageConfigurationName!: pulumi.Output<string>;

    /**
     * Create a MwsStorageConfigurations resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MwsStorageConfigurationsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MwsStorageConfigurationsArgs | MwsStorageConfigurationsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MwsStorageConfigurationsState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["bucketName"] = state ? state.bucketName : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["storageConfigurationId"] = state ? state.storageConfigurationId : undefined;
            inputs["storageConfigurationName"] = state ? state.storageConfigurationName : undefined;
        } else {
            const args = argsOrState as MwsStorageConfigurationsArgs | undefined;
            if (!args || args.accountId === undefined) {
                throw new Error("Missing required property 'accountId'");
            }
            if (!args || args.bucketName === undefined) {
                throw new Error("Missing required property 'bucketName'");
            }
            if (!args || args.storageConfigurationName === undefined) {
                throw new Error("Missing required property 'storageConfigurationName'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["bucketName"] = args ? args.bucketName : undefined;
            inputs["storageConfigurationName"] = args ? args.storageConfigurationName : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["storageConfigurationId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MwsStorageConfigurations.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MwsStorageConfigurations resources.
 */
export interface MwsStorageConfigurationsState {
    readonly accountId?: pulumi.Input<string>;
    readonly bucketName?: pulumi.Input<string>;
    readonly creationTime?: pulumi.Input<number>;
    readonly storageConfigurationId?: pulumi.Input<string>;
    readonly storageConfigurationName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MwsStorageConfigurations resource.
 */
export interface MwsStorageConfigurationsArgs {
    readonly accountId: pulumi.Input<string>;
    readonly bucketName: pulumi.Input<string>;
    readonly storageConfigurationName: pulumi.Input<string>;
}
