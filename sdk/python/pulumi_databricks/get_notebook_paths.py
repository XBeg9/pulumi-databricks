# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetNotebookPathsResult:
    """
    A collection of values returned by getNotebookPaths.
    """
    def __init__(__self__, id=None, notebook_path_lists=None, path=None, recursive=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if notebook_path_lists and not isinstance(notebook_path_lists, list):
            raise TypeError("Expected argument 'notebook_path_lists' to be a list")
        __self__.notebook_path_lists = notebook_path_lists
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        if recursive and not isinstance(recursive, bool):
            raise TypeError("Expected argument 'recursive' to be a bool")
        __self__.recursive = recursive
class AwaitableGetNotebookPathsResult(GetNotebookPathsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNotebookPathsResult(
            id=self.id,
            notebook_path_lists=self.notebook_path_lists,
            path=self.path,
            recursive=self.recursive)

def get_notebook_paths(path=None,recursive=None,opts=None):
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
    __ret__ = pulumi.runtime.invoke('databricks:index/getNotebookPaths:getNotebookPaths', __args__, opts=opts).value

    return AwaitableGetNotebookPathsResult(
        id=__ret__.get('id'),
        notebook_path_lists=__ret__.get('notebookPathLists'),
        path=__ret__.get('path'),
        recursive=__ret__.get('recursive'))
