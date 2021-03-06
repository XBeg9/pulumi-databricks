# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class InstancePool(pulumi.CustomResource):
    aws_attributes: pulumi.Output[dict]
    custom_tags: pulumi.Output[dict]
    default_tags: pulumi.Output[dict]
    disk_spec: pulumi.Output[dict]
    enable_elastic_disk: pulumi.Output[bool]
    idle_instance_autotermination_minutes: pulumi.Output[float]
    instance_pool_name: pulumi.Output[str]
    max_capacity: pulumi.Output[float]
    min_idle_instances: pulumi.Output[float]
    node_type_id: pulumi.Output[str]
    preloaded_spark_versions: pulumi.Output[list]
    state: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, aws_attributes=None, custom_tags=None, disk_spec=None, enable_elastic_disk=None, idle_instance_autotermination_minutes=None, instance_pool_name=None, max_capacity=None, min_idle_instances=None, node_type_id=None, preloaded_spark_versions=None, state=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a InstancePool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **aws_attributes** object supports the following:

          * `availability` (`pulumi.Input[str]`)
          * `spotBidPricePercent` (`pulumi.Input[float]`)
          * `zoneId` (`pulumi.Input[str]`)

        The **disk_spec** object supports the following:

          * `azureDiskVolumeType` (`pulumi.Input[str]`)
          * `diskCount` (`pulumi.Input[float]`)
          * `diskSize` (`pulumi.Input[float]`)
          * `ebsVolumeType` (`pulumi.Input[str]`)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['aws_attributes'] = aws_attributes
            __props__['custom_tags'] = custom_tags
            __props__['disk_spec'] = disk_spec
            __props__['enable_elastic_disk'] = enable_elastic_disk
            if idle_instance_autotermination_minutes is None:
                raise TypeError("Missing required property 'idle_instance_autotermination_minutes'")
            __props__['idle_instance_autotermination_minutes'] = idle_instance_autotermination_minutes
            if instance_pool_name is None:
                raise TypeError("Missing required property 'instance_pool_name'")
            __props__['instance_pool_name'] = instance_pool_name
            if max_capacity is None:
                raise TypeError("Missing required property 'max_capacity'")
            __props__['max_capacity'] = max_capacity
            if min_idle_instances is None:
                raise TypeError("Missing required property 'min_idle_instances'")
            __props__['min_idle_instances'] = min_idle_instances
            if node_type_id is None:
                raise TypeError("Missing required property 'node_type_id'")
            __props__['node_type_id'] = node_type_id
            __props__['preloaded_spark_versions'] = preloaded_spark_versions
            __props__['state'] = state
            __props__['default_tags'] = None
        super(InstancePool, __self__).__init__(
            'databricks:index/instancePool:InstancePool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, aws_attributes=None, custom_tags=None, default_tags=None, disk_spec=None, enable_elastic_disk=None, idle_instance_autotermination_minutes=None, instance_pool_name=None, max_capacity=None, min_idle_instances=None, node_type_id=None, preloaded_spark_versions=None, state=None):
        """
        Get an existing InstancePool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **aws_attributes** object supports the following:

          * `availability` (`pulumi.Input[str]`)
          * `spotBidPricePercent` (`pulumi.Input[float]`)
          * `zoneId` (`pulumi.Input[str]`)

        The **disk_spec** object supports the following:

          * `azureDiskVolumeType` (`pulumi.Input[str]`)
          * `diskCount` (`pulumi.Input[float]`)
          * `diskSize` (`pulumi.Input[float]`)
          * `ebsVolumeType` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["aws_attributes"] = aws_attributes
        __props__["custom_tags"] = custom_tags
        __props__["default_tags"] = default_tags
        __props__["disk_spec"] = disk_spec
        __props__["enable_elastic_disk"] = enable_elastic_disk
        __props__["idle_instance_autotermination_minutes"] = idle_instance_autotermination_minutes
        __props__["instance_pool_name"] = instance_pool_name
        __props__["max_capacity"] = max_capacity
        __props__["min_idle_instances"] = min_idle_instances
        __props__["node_type_id"] = node_type_id
        __props__["preloaded_spark_versions"] = preloaded_spark_versions
        __props__["state"] = state
        return InstancePool(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

