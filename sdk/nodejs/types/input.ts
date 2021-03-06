// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";

export interface ClusterAutoscale {
    maxWorkers?: pulumi.Input<number>;
    minWorkers?: pulumi.Input<number>;
}

export interface ClusterAwsAttributes {
    availability?: pulumi.Input<string>;
    ebsVolumeCount?: pulumi.Input<number>;
    ebsVolumeSize?: pulumi.Input<number>;
    ebsVolumeType?: pulumi.Input<string>;
    firstOnDemand?: pulumi.Input<number>;
    instanceProfileArn?: pulumi.Input<string>;
    spotBidPricePercent?: pulumi.Input<number>;
    zoneId: pulumi.Input<string>;
}

export interface ClusterClusterLogConf {
    dbfs?: pulumi.Input<inputs.ClusterClusterLogConfDbfs>;
    s3?: pulumi.Input<inputs.ClusterClusterLogConfS3>;
}

export interface ClusterClusterLogConfDbfs {
    destination: pulumi.Input<string>;
}

export interface ClusterClusterLogConfS3 {
    cannedAcl?: pulumi.Input<string>;
    destination: pulumi.Input<string>;
    enableEncryption?: pulumi.Input<boolean>;
    encryptionType?: pulumi.Input<string>;
    endpoint?: pulumi.Input<string>;
    kmsKey?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}

export interface ClusterDockerImage {
    basicAuth?: pulumi.Input<inputs.ClusterDockerImageBasicAuth>;
    url: pulumi.Input<string>;
}

export interface ClusterDockerImageBasicAuth {
    password: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export interface ClusterInitScript {
    dbfs?: pulumi.Input<inputs.ClusterInitScriptDbfs>;
    s3?: pulumi.Input<inputs.ClusterInitScriptS3>;
}

export interface ClusterInitScriptDbfs {
    destination: pulumi.Input<string>;
}

export interface ClusterInitScriptS3 {
    cannedAcl?: pulumi.Input<string>;
    destination: pulumi.Input<string>;
    enableEncryption?: pulumi.Input<boolean>;
    encryptionType?: pulumi.Input<string>;
    endpoint?: pulumi.Input<string>;
    kmsKey?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}

export interface ClusterLibraryCran {
    messages?: pulumi.Input<string>;
    package?: pulumi.Input<string>;
    repo?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

export interface ClusterLibraryEgg {
    messages?: pulumi.Input<pulumi.Input<string>[]>;
    path?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

export interface ClusterLibraryJar {
    messages?: pulumi.Input<pulumi.Input<string>[]>;
    path?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

export interface ClusterLibraryMaven {
    coordinates?: pulumi.Input<string>;
    exclusions?: pulumi.Input<pulumi.Input<string>[]>;
    messages?: pulumi.Input<pulumi.Input<string>[]>;
    repo?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

export interface ClusterLibraryPypi {
    messages?: pulumi.Input<pulumi.Input<string>[]>;
    package: pulumi.Input<string>;
    repo?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

export interface ClusterLibraryWhl {
    messages?: pulumi.Input<pulumi.Input<string>[]>;
    path?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

export interface InstancePoolAwsAttributes {
    availability?: pulumi.Input<string>;
    spotBidPricePercent?: pulumi.Input<number>;
    zoneId: pulumi.Input<string>;
}

export interface InstancePoolDiskSpec {
    azureDiskVolumeType?: pulumi.Input<string>;
    diskCount?: pulumi.Input<number>;
    diskSize?: pulumi.Input<number>;
    ebsVolumeType?: pulumi.Input<string>;
}

export interface JobEmailNotifications {
    noAlertForSkippedRuns?: pulumi.Input<boolean>;
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    onStarts?: pulumi.Input<pulumi.Input<string>[]>;
    onSuccesses?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface JobLibraryCran {
    package: pulumi.Input<string>;
    repo?: pulumi.Input<string>;
}

export interface JobLibraryMaven {
    coordinates: pulumi.Input<string>;
    exclusions?: pulumi.Input<pulumi.Input<string>[]>;
    repo?: pulumi.Input<string>;
}

export interface JobLibraryPypi {
    package: pulumi.Input<string>;
    repo?: pulumi.Input<string>;
}

export interface JobNewCluster {
    autoscale?: pulumi.Input<inputs.JobNewClusterAutoscale>;
    autoterminationMinutes?: pulumi.Input<number>;
    awsAttributes?: pulumi.Input<inputs.JobNewClusterAwsAttributes>;
    clusterLogConf?: pulumi.Input<inputs.JobNewClusterClusterLogConf>;
    clusterName?: pulumi.Input<string>;
    customTags?: pulumi.Input<{[key: string]: any}>;
    dockerImage?: pulumi.Input<inputs.JobNewClusterDockerImage>;
    driverNodeTypeId?: pulumi.Input<string>;
    enableElasticDisk?: pulumi.Input<boolean>;
    initScripts?: pulumi.Input<pulumi.Input<inputs.JobNewClusterInitScript>[]>;
    instancePoolId?: pulumi.Input<string>;
    nodeTypeId?: pulumi.Input<string>;
    numWorkers?: pulumi.Input<number>;
    sparkConf?: pulumi.Input<{[key: string]: any}>;
    sparkEnvVars?: pulumi.Input<{[key: string]: any}>;
    sparkVersion?: pulumi.Input<string>;
    sshPublicKeys?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface JobNewClusterAutoscale {
    maxWorkers?: pulumi.Input<number>;
    minWorkers?: pulumi.Input<number>;
}

export interface JobNewClusterAwsAttributes {
    availability?: pulumi.Input<string>;
    ebsVolumeCount?: pulumi.Input<number>;
    ebsVolumeSize?: pulumi.Input<number>;
    ebsVolumeType?: pulumi.Input<string>;
    firstOnDemand?: pulumi.Input<number>;
    instanceProfileArn?: pulumi.Input<string>;
    spotBidPricePercent?: pulumi.Input<number>;
    zoneId: pulumi.Input<string>;
}

export interface JobNewClusterClusterLogConf {
    dbfs?: pulumi.Input<inputs.JobNewClusterClusterLogConfDbfs>;
    s3?: pulumi.Input<inputs.JobNewClusterClusterLogConfS3>;
}

export interface JobNewClusterClusterLogConfDbfs {
    destination: pulumi.Input<string>;
}

export interface JobNewClusterClusterLogConfS3 {
    cannedAcl?: pulumi.Input<string>;
    destination: pulumi.Input<string>;
    enableEncryption?: pulumi.Input<boolean>;
    encryptionType?: pulumi.Input<string>;
    endpoint?: pulumi.Input<string>;
    kmsKey?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}

export interface JobNewClusterDockerImage {
    basicAuth?: pulumi.Input<inputs.JobNewClusterDockerImageBasicAuth>;
    url: pulumi.Input<string>;
}

export interface JobNewClusterDockerImageBasicAuth {
    password: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export interface JobNewClusterInitScript {
    dbfs?: pulumi.Input<inputs.JobNewClusterInitScriptDbfs>;
    s3?: pulumi.Input<inputs.JobNewClusterInitScriptS3>;
}

export interface JobNewClusterInitScriptDbfs {
    destination: pulumi.Input<string>;
}

export interface JobNewClusterInitScriptS3 {
    cannedAcl?: pulumi.Input<string>;
    destination: pulumi.Input<string>;
    enableEncryption?: pulumi.Input<boolean>;
    encryptionType?: pulumi.Input<string>;
    endpoint?: pulumi.Input<string>;
    kmsKey?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}

export interface JobSchedule {
    quartzCronExpression: pulumi.Input<string>;
    timezoneId: pulumi.Input<string>;
}

export interface ProviderAzureAuth {
    azureRegion: pulumi.Input<string>;
    clientId: pulumi.Input<string>;
    clientSecret: pulumi.Input<string>;
    managedResourceGroup: pulumi.Input<string>;
    resourceGroup: pulumi.Input<string>;
    subscriptionId: pulumi.Input<string>;
    tenantId: pulumi.Input<string>;
    workspaceName: pulumi.Input<string>;
}

export interface ProviderBasicAuth {
    password: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export namespace aws {
    export interface MwsNetworksErrorMessage {
        errorMessage?: pulumi.Input<string>;
        errorType?: pulumi.Input<string>;
    }

    export interface MwsWorkspacesNetworkErrorMessage {
        errorMessage?: pulumi.Input<string>;
        errorType?: pulumi.Input<string>;
    }
}
