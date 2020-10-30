// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface LXCContainerFeatures {
    fuse?: pulumi.Input<boolean>;
    keyctl?: pulumi.Input<boolean>;
    mount?: pulumi.Input<string>;
    nesting?: pulumi.Input<boolean>;
}

export interface LXCContainerMountpoint {
    acl?: pulumi.Input<boolean>;
    backup?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    mp: pulumi.Input<string>;
    quota?: pulumi.Input<boolean>;
    replicate?: pulumi.Input<boolean>;
    shared?: pulumi.Input<boolean>;
    size: pulumi.Input<string>;
    slot: pulumi.Input<number>;
    storage: pulumi.Input<string>;
    volume?: pulumi.Input<string>;
}

export interface LXCContainerNetwork {
    bridge?: pulumi.Input<string>;
    firewall?: pulumi.Input<boolean>;
    gw?: pulumi.Input<string>;
    gw6?: pulumi.Input<string>;
    hwaddr?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    ip6?: pulumi.Input<string>;
    mtu?: pulumi.Input<string>;
    name: pulumi.Input<string>;
    rate?: pulumi.Input<number>;
    tag?: pulumi.Input<number>;
    trunks?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}

export interface LXCContainerRootfs {
    size: pulumi.Input<string>;
    storage: pulumi.Input<string>;
    volume?: pulumi.Input<string>;
}

export interface LXCDiskMountoptions {
    noatime?: pulumi.Input<boolean>;
    nodev?: pulumi.Input<boolean>;
    noexec?: pulumi.Input<string>;
    nosuid?: pulumi.Input<boolean>;
}

export interface QemuVMDisk {
    backup?: pulumi.Input<boolean>;
    cache?: pulumi.Input<string>;
    discard?: pulumi.Input<string>;
    file?: pulumi.Input<string>;
    format?: pulumi.Input<string>;
    iothread?: pulumi.Input<boolean>;
    mbps?: pulumi.Input<number>;
    mbpsRd?: pulumi.Input<number>;
    mbpsRdMax?: pulumi.Input<number>;
    mbpsWr?: pulumi.Input<number>;
    mbpsWrMax?: pulumi.Input<number>;
    media?: pulumi.Input<string>;
    replicate?: pulumi.Input<boolean>;
    size: pulumi.Input<string>;
    ssd?: pulumi.Input<boolean>;
    storage: pulumi.Input<string>;
    type: pulumi.Input<string>;
    volume?: pulumi.Input<string>;
}

export interface QemuVMNetwork {
    bridge?: pulumi.Input<string>;
    firewall?: pulumi.Input<boolean>;
    linkDown?: pulumi.Input<boolean>;
    macaddr?: pulumi.Input<string>;
    model: pulumi.Input<string>;
    queues?: pulumi.Input<number>;
    rate?: pulumi.Input<number>;
    tag?: pulumi.Input<number>;
}

export interface QemuVMSerial {
    id: pulumi.Input<number>;
    type: pulumi.Input<string>;
}

export interface QemuVMVga {
    memory?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
}
