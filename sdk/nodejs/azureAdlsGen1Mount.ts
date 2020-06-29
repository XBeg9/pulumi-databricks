// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AzureAdlsGen1Mount extends pulumi.CustomResource {
    /**
     * Get an existing AzureAdlsGen1Mount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureAdlsGen1MountState, opts?: pulumi.CustomResourceOptions): AzureAdlsGen1Mount {
        return new AzureAdlsGen1Mount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'databricks:index/azureAdlsGen1Mount:AzureAdlsGen1Mount';

    /**
     * Returns true if the given object is an instance of AzureAdlsGen1Mount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureAdlsGen1Mount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureAdlsGen1Mount.__pulumiType;
    }

    public readonly clientId!: pulumi.Output<string>;
    public readonly clientSecretKey!: pulumi.Output<string>;
    public readonly clientSecretScope!: pulumi.Output<string>;
    public readonly clusterId!: pulumi.Output<string>;
    public readonly directory!: pulumi.Output<string>;
    public readonly mountName!: pulumi.Output<string>;
    public readonly sparkConfPrefix!: pulumi.Output<string | undefined>;
    public readonly storageResourceName!: pulumi.Output<string>;
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a AzureAdlsGen1Mount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureAdlsGen1MountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureAdlsGen1MountArgs | AzureAdlsGen1MountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AzureAdlsGen1MountState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecretKey"] = state ? state.clientSecretKey : undefined;
            inputs["clientSecretScope"] = state ? state.clientSecretScope : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["directory"] = state ? state.directory : undefined;
            inputs["mountName"] = state ? state.mountName : undefined;
            inputs["sparkConfPrefix"] = state ? state.sparkConfPrefix : undefined;
            inputs["storageResourceName"] = state ? state.storageResourceName : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as AzureAdlsGen1MountArgs | undefined;
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.clientSecretKey === undefined) {
                throw new Error("Missing required property 'clientSecretKey'");
            }
            if (!args || args.clientSecretScope === undefined) {
                throw new Error("Missing required property 'clientSecretScope'");
            }
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.mountName === undefined) {
                throw new Error("Missing required property 'mountName'");
            }
            if (!args || args.storageResourceName === undefined) {
                throw new Error("Missing required property 'storageResourceName'");
            }
            if (!args || args.tenantId === undefined) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecretKey"] = args ? args.clientSecretKey : undefined;
            inputs["clientSecretScope"] = args ? args.clientSecretScope : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["directory"] = args ? args.directory : undefined;
            inputs["mountName"] = args ? args.mountName : undefined;
            inputs["sparkConfPrefix"] = args ? args.sparkConfPrefix : undefined;
            inputs["storageResourceName"] = args ? args.storageResourceName : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AzureAdlsGen1Mount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureAdlsGen1Mount resources.
 */
export interface AzureAdlsGen1MountState {
    readonly clientId?: pulumi.Input<string>;
    readonly clientSecretKey?: pulumi.Input<string>;
    readonly clientSecretScope?: pulumi.Input<string>;
    readonly clusterId?: pulumi.Input<string>;
    readonly directory?: pulumi.Input<string>;
    readonly mountName?: pulumi.Input<string>;
    readonly sparkConfPrefix?: pulumi.Input<string>;
    readonly storageResourceName?: pulumi.Input<string>;
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureAdlsGen1Mount resource.
 */
export interface AzureAdlsGen1MountArgs {
    readonly clientId: pulumi.Input<string>;
    readonly clientSecretKey: pulumi.Input<string>;
    readonly clientSecretScope: pulumi.Input<string>;
    readonly clusterId: pulumi.Input<string>;
    readonly directory?: pulumi.Input<string>;
    readonly mountName: pulumi.Input<string>;
    readonly sparkConfPrefix?: pulumi.Input<string>;
    readonly storageResourceName: pulumi.Input<string>;
    readonly tenantId: pulumi.Input<string>;
}