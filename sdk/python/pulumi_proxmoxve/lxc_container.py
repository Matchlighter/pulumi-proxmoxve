# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['LXCContainer']


class LXCContainer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arch: Optional[pulumi.Input[str]] = None,
                 bwlimit: Optional[pulumi.Input[int]] = None,
                 cmode: Optional[pulumi.Input[str]] = None,
                 console: Optional[pulumi.Input[bool]] = None,
                 cores: Optional[pulumi.Input[int]] = None,
                 cpulimit: Optional[pulumi.Input[int]] = None,
                 cpuunits: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 features: Optional[pulumi.Input[pulumi.InputType['LXCContainerFeaturesArgs']]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 hookscript: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ignore_unpack_errors: Optional[pulumi.Input[bool]] = None,
                 lock: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 mountpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LXCContainerMountpointArgs']]]]] = None,
                 nameserver: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LXCContainerNetworkArgs']]]]] = None,
                 onboot: Optional[pulumi.Input[bool]] = None,
                 ostemplate: Optional[pulumi.Input[str]] = None,
                 ostype: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 pool: Optional[pulumi.Input[str]] = None,
                 protection: Optional[pulumi.Input[bool]] = None,
                 restore: Optional[pulumi.Input[bool]] = None,
                 rootfs: Optional[pulumi.Input[pulumi.InputType['LXCContainerRootfsArgs']]] = None,
                 searchdomain: Optional[pulumi.Input[str]] = None,
                 ssh_public_keys: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[bool]] = None,
                 startup: Optional[pulumi.Input[str]] = None,
                 swap: Optional[pulumi.Input[int]] = None,
                 target_node: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[bool]] = None,
                 tty: Optional[pulumi.Input[int]] = None,
                 unique: Optional[pulumi.Input[bool]] = None,
                 unprivileged: Optional[pulumi.Input[bool]] = None,
                 unuseds: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vmid: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a LXCContainer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['arch'] = arch
            __props__['bwlimit'] = bwlimit
            __props__['cmode'] = cmode
            __props__['console'] = console
            __props__['cores'] = cores
            __props__['cpulimit'] = cpulimit
            __props__['cpuunits'] = cpuunits
            __props__['description'] = description
            __props__['features'] = features
            __props__['force'] = force
            __props__['hookscript'] = hookscript
            __props__['hostname'] = hostname
            __props__['ignore_unpack_errors'] = ignore_unpack_errors
            __props__['lock'] = lock
            __props__['memory'] = memory
            __props__['mountpoints'] = mountpoints
            __props__['nameserver'] = nameserver
            __props__['networks'] = networks
            __props__['onboot'] = onboot
            __props__['ostemplate'] = ostemplate
            __props__['ostype'] = ostype
            __props__['password'] = password
            __props__['pool'] = pool
            __props__['protection'] = protection
            __props__['restore'] = restore
            __props__['rootfs'] = rootfs
            __props__['searchdomain'] = searchdomain
            __props__['ssh_public_keys'] = ssh_public_keys
            __props__['start'] = start
            __props__['startup'] = startup
            __props__['swap'] = swap
            if target_node is None:
                raise TypeError("Missing required property 'target_node'")
            __props__['target_node'] = target_node
            __props__['template'] = template
            __props__['tty'] = tty
            __props__['unique'] = unique
            __props__['unprivileged'] = unprivileged
            __props__['unuseds'] = unuseds
            __props__['vmid'] = vmid
        super(LXCContainer, __self__).__init__(
            'proxmoxve:index/lXCContainer:LXCContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arch: Optional[pulumi.Input[str]] = None,
            bwlimit: Optional[pulumi.Input[int]] = None,
            cmode: Optional[pulumi.Input[str]] = None,
            console: Optional[pulumi.Input[bool]] = None,
            cores: Optional[pulumi.Input[int]] = None,
            cpulimit: Optional[pulumi.Input[int]] = None,
            cpuunits: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            features: Optional[pulumi.Input[pulumi.InputType['LXCContainerFeaturesArgs']]] = None,
            force: Optional[pulumi.Input[bool]] = None,
            hookscript: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            ignore_unpack_errors: Optional[pulumi.Input[bool]] = None,
            lock: Optional[pulumi.Input[str]] = None,
            memory: Optional[pulumi.Input[int]] = None,
            mountpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LXCContainerMountpointArgs']]]]] = None,
            nameserver: Optional[pulumi.Input[str]] = None,
            networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LXCContainerNetworkArgs']]]]] = None,
            onboot: Optional[pulumi.Input[bool]] = None,
            ostemplate: Optional[pulumi.Input[str]] = None,
            ostype: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            pool: Optional[pulumi.Input[str]] = None,
            protection: Optional[pulumi.Input[bool]] = None,
            restore: Optional[pulumi.Input[bool]] = None,
            rootfs: Optional[pulumi.Input[pulumi.InputType['LXCContainerRootfsArgs']]] = None,
            searchdomain: Optional[pulumi.Input[str]] = None,
            ssh_public_keys: Optional[pulumi.Input[str]] = None,
            start: Optional[pulumi.Input[bool]] = None,
            startup: Optional[pulumi.Input[str]] = None,
            swap: Optional[pulumi.Input[int]] = None,
            target_node: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[bool]] = None,
            tty: Optional[pulumi.Input[int]] = None,
            unique: Optional[pulumi.Input[bool]] = None,
            unprivileged: Optional[pulumi.Input[bool]] = None,
            unuseds: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vmid: Optional[pulumi.Input[int]] = None) -> 'LXCContainer':
        """
        Get an existing LXCContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arch"] = arch
        __props__["bwlimit"] = bwlimit
        __props__["cmode"] = cmode
        __props__["console"] = console
        __props__["cores"] = cores
        __props__["cpulimit"] = cpulimit
        __props__["cpuunits"] = cpuunits
        __props__["description"] = description
        __props__["features"] = features
        __props__["force"] = force
        __props__["hookscript"] = hookscript
        __props__["hostname"] = hostname
        __props__["ignore_unpack_errors"] = ignore_unpack_errors
        __props__["lock"] = lock
        __props__["memory"] = memory
        __props__["mountpoints"] = mountpoints
        __props__["nameserver"] = nameserver
        __props__["networks"] = networks
        __props__["onboot"] = onboot
        __props__["ostemplate"] = ostemplate
        __props__["ostype"] = ostype
        __props__["password"] = password
        __props__["pool"] = pool
        __props__["protection"] = protection
        __props__["restore"] = restore
        __props__["rootfs"] = rootfs
        __props__["searchdomain"] = searchdomain
        __props__["ssh_public_keys"] = ssh_public_keys
        __props__["start"] = start
        __props__["startup"] = startup
        __props__["swap"] = swap
        __props__["target_node"] = target_node
        __props__["template"] = template
        __props__["tty"] = tty
        __props__["unique"] = unique
        __props__["unprivileged"] = unprivileged
        __props__["unuseds"] = unuseds
        __props__["vmid"] = vmid
        return LXCContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arch(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "arch")

    @property
    @pulumi.getter
    def bwlimit(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "bwlimit")

    @property
    @pulumi.getter
    def cmode(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cmode")

    @property
    @pulumi.getter
    def console(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "console")

    @property
    @pulumi.getter
    def cores(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "cores")

    @property
    @pulumi.getter
    def cpulimit(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "cpulimit")

    @property
    @pulumi.getter
    def cpuunits(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "cpuunits")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def features(self) -> pulumi.Output[Optional['outputs.LXCContainerFeatures']]:
        return pulumi.get(self, "features")

    @property
    @pulumi.getter
    def force(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "force")

    @property
    @pulumi.getter
    def hookscript(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "hookscript")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="ignoreUnpackErrors")
    def ignore_unpack_errors(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "ignore_unpack_errors")

    @property
    @pulumi.getter
    def lock(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "lock")

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def mountpoints(self) -> pulumi.Output[Optional[Sequence['outputs.LXCContainerMountpoint']]]:
        return pulumi.get(self, "mountpoints")

    @property
    @pulumi.getter
    def nameserver(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "nameserver")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Optional[Sequence['outputs.LXCContainerNetwork']]]:
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def onboot(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "onboot")

    @property
    @pulumi.getter
    def ostemplate(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ostemplate")

    @property
    @pulumi.getter
    def ostype(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ostype")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def pool(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "pool")

    @property
    @pulumi.getter
    def protection(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "protection")

    @property
    @pulumi.getter
    def restore(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "restore")

    @property
    @pulumi.getter
    def rootfs(self) -> pulumi.Output[Optional['outputs.LXCContainerRootfs']]:
        return pulumi.get(self, "rootfs")

    @property
    @pulumi.getter
    def searchdomain(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "searchdomain")

    @property
    @pulumi.getter(name="sshPublicKeys")
    def ssh_public_keys(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ssh_public_keys")

    @property
    @pulumi.getter
    def start(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def startup(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "startup")

    @property
    @pulumi.getter
    def swap(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "swap")

    @property
    @pulumi.getter(name="targetNode")
    def target_node(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_node")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def tty(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "tty")

    @property
    @pulumi.getter
    def unique(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "unique")

    @property
    @pulumi.getter
    def unprivileged(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "unprivileged")

    @property
    @pulumi.getter
    def unuseds(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "unuseds")

    @property
    @pulumi.getter
    def vmid(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "vmid")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

