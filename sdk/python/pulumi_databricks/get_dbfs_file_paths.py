# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetDbfsFilePathsResult:
    """
    A collection of values returned by getDbfsFilePaths.
    """
    def __init__(__self__, id=None, path=None, path_lists=None, recursive=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        if path_lists and not isinstance(path_lists, list):
            raise TypeError("Expected argument 'path_lists' to be a list")
        __self__.path_lists = path_lists
        if recursive and not isinstance(recursive, bool):
            raise TypeError("Expected argument 'recursive' to be a bool")
        __self__.recursive = recursive
class AwaitableGetDbfsFilePathsResult(GetDbfsFilePathsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDbfsFilePathsResult(
            id=self.id,
            path=self.path,
            path_lists=self.path_lists,
            recursive=self.recursive)

def get_dbfs_file_paths(path=None,recursive=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['path'] = path
    __args__['recursive'] = recursive
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('databricks:index/getDbfsFilePaths:getDbfsFilePaths', __args__, opts=opts).value

    return AwaitableGetDbfsFilePathsResult(
        id=__ret__.get('id'),
        path=__ret__.get('path'),
        path_lists=__ret__.get('pathLists'),
        recursive=__ret__.get('recursive'))
