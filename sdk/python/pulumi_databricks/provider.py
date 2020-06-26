# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, azure_auth=None, basic_auth=None, config_file=None, host=None, profile=None, token=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the databricks package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_file: Location of the Databricks CLI credentials file, that is created by `databricks configure --token` command. By default,
               it is located in ~/.databrickscfg. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for
               docs. Config file credetials will only be used when host/token are not provided.
        :param pulumi.Input[str] profile: Connection profile specified within ~/.databrickscfg. Please check
               https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.

        The **azure_auth** object supports the following:

          * `azureRegion` (`pulumi.Input[str]`)
          * `clientId` (`pulumi.Input[str]`)
          * `clientSecret` (`pulumi.Input[str]`)
          * `managedResourceGroup` (`pulumi.Input[str]`)
          * `resourceGroup` (`pulumi.Input[str]`)
          * `subscriptionId` (`pulumi.Input[str]`)
          * `tenantId` (`pulumi.Input[str]`)
          * `workspaceName` (`pulumi.Input[str]`)

        The **basic_auth** object supports the following:

          * `password` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)
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

            __props__['azure_auth'] = pulumi.Output.from_input(azure_auth).apply(json.dumps) if azure_auth is not None else None
            __props__['basic_auth'] = pulumi.Output.from_input(basic_auth).apply(json.dumps) if basic_auth is not None else None
            __props__['config_file'] = config_file
            __props__['host'] = host
            __props__['profile'] = profile
            __props__['token'] = token
        super(Provider, __self__).__init__(
            'databricks',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
