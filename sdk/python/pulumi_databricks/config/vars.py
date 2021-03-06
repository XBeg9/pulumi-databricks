# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('databricks')

azure_auth = __config__.get('azureAuth')

basic_auth = __config__.get('basicAuth')

config_file = __config__.get('configFile')
"""
Location of the Databricks CLI credentials file, that is created by `databricks configure --token` command. By default,
it is located in ~/.databrickscfg. Check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for
docs. Config file credetials will only be used when host/token are not provided.
"""

host = __config__.get('host') or utilities.get_env('DATABRICKS_HOST')

profile = __config__.get('profile')
"""
Connection profile specified within ~/.databrickscfg. Please check
https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.
"""

token = __config__.get('token') or utilities.get_env('DATABRICKS_TOKEN')

