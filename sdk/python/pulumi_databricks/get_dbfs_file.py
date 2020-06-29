# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetDbfsFileResult:
    """
    A collection of values returned by getDbfsFile.
    """
    def __init__(__self__, content=None, file_size=None, id=None, limit_file_size=None, path=None):
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        __self__.content = content
        if file_size and not isinstance(file_size, float):
            raise TypeError("Expected argument 'file_size' to be a float")
        __self__.file_size = file_size
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if limit_file_size and not isinstance(limit_file_size, bool):
            raise TypeError("Expected argument 'limit_file_size' to be a bool")
        __self__.limit_file_size = limit_file_size
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
class AwaitableGetDbfsFileResult(GetDbfsFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDbfsFileResult(
            content=self.content,
            file_size=self.file_size,
            id=self.id,
            limit_file_size=self.limit_file_size,
            path=self.path)

def get_dbfs_file(limit_file_size=None,path=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['limitFileSize'] = limit_file_size
    __args__['path'] = path
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('databricks:index/getDbfsFile:getDbfsFile', __args__, opts=opts).value

    return AwaitableGetDbfsFileResult(
        content=__ret__.get('content'),
        file_size=__ret__.get('fileSize'),
        id=__ret__.get('id'),
        limit_file_size=__ret__.get('limitFileSize'),
        path=__ret__.get('path'))