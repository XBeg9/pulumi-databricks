// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AwsS3Mount extends pulumi.CustomResource {
    /**
     * Get an existing AwsS3Mount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsS3MountState, opts?: pulumi.CustomResourceOptions): AwsS3Mount {
        return new AwsS3Mount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'databricks:aws/awsS3Mount:AwsS3Mount';

    /**
     * Returns true if the given object is an instance of AwsS3Mount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AwsS3Mount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AwsS3Mount.__pulumiType;
    }

    public readonly clusterId!: pulumi.Output<string>;
    public readonly mountName!: pulumi.Output<string>;
    public readonly s3BucketName!: pulumi.Output<string>;

    /**
     * Create a AwsS3Mount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AwsS3MountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AwsS3MountArgs | AwsS3MountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AwsS3MountState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["mountName"] = state ? state.mountName : undefined;
            inputs["s3BucketName"] = state ? state.s3BucketName : undefined;
        } else {
            const args = argsOrState as AwsS3MountArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.mountName === undefined) {
                throw new Error("Missing required property 'mountName'");
            }
            if (!args || args.s3BucketName === undefined) {
                throw new Error("Missing required property 's3BucketName'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["mountName"] = args ? args.mountName : undefined;
            inputs["s3BucketName"] = args ? args.s3BucketName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AwsS3Mount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AwsS3Mount resources.
 */
export interface AwsS3MountState {
    readonly clusterId?: pulumi.Input<string>;
    readonly mountName?: pulumi.Input<string>;
    readonly s3BucketName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AwsS3Mount resource.
 */
export interface AwsS3MountArgs {
    readonly clusterId: pulumi.Input<string>;
    readonly mountName: pulumi.Input<string>;
    readonly s3BucketName: pulumi.Input<string>;
}
