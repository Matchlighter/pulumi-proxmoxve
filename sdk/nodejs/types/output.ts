// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface LXCContainerFeatures {
    fuse?: boolean;
    keyctl?: boolean;
    mount?: string;
    nesting?: boolean;
}

export interface LXCContainerMountpoint {
    acl?: boolean;
    backup?: boolean;
    key: string;
    mp: string;
    quota?: boolean;
    replicate?: boolean;
    shared?: boolean;
    size: string;
    slot: number;
    storage: string;
    volume: string;
}

export interface LXCContainerNetwork {
    bridge?: string;
    firewall?: boolean;
    gw?: string;
    gw6?: string;
    hwaddr: string;
    ip?: string;
    ip6?: string;
    mtu?: string;
    name: string;
    rate?: number;
    tag: number;
    trunks: string;
    type: string;
}

export interface LXCContainerRootfs {
    size: string;
    storage: string;
    volume: string;
}

export interface LXCDiskMountoptions {
    noatime?: boolean;
    nodev?: boolean;
    noexec?: string;
    nosuid?: boolean;
}

export interface QemuVMDisk {
    backup?: boolean;
    cache?: string;
    discard?: string;
    file: string;
    format?: string;
    iothread?: boolean;
    mbps?: number;
    mbpsRd?: number;
    mbpsRdMax?: number;
    mbpsWr?: number;
    mbpsWrMax?: number;
    media: string;
    replicate?: boolean;
    size: string;
    ssd?: boolean;
    storage: string;
    type: string;
    volume: string;
}

export interface QemuVMNetwork {
    bridge?: string;
    firewall?: boolean;
    linkDown?: boolean;
    macaddr: string;
    model: string;
    queues: number;
    rate: number;
    tag?: number;
}

export interface QemuVMSerial {
    id: number;
    type: string;
}

export interface QemuVMVga {
    memory?: number;
    type?: string;
}
