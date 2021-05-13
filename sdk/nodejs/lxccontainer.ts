// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LXCContainer extends pulumi.CustomResource {
    /**
     * Get an existing LXCContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LXCContainerState, opts?: pulumi.CustomResourceOptions): LXCContainer {
        return new LXCContainer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'proxmoxve:index/lXCContainer:LXCContainer';

    /**
     * Returns true if the given object is an instance of LXCContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LXCContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LXCContainer.__pulumiType;
    }

    public readonly arch!: pulumi.Output<string | undefined>;
    public readonly bwlimit!: pulumi.Output<number | undefined>;
    public readonly cmode!: pulumi.Output<string | undefined>;
    public readonly console!: pulumi.Output<boolean | undefined>;
    public readonly cores!: pulumi.Output<number | undefined>;
    public readonly cpulimit!: pulumi.Output<number | undefined>;
    public readonly cpuunits!: pulumi.Output<number | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly features!: pulumi.Output<outputs.LXCContainerFeatures | undefined>;
    public readonly force!: pulumi.Output<boolean | undefined>;
    public readonly hookscript!: pulumi.Output<string | undefined>;
    public readonly hostname!: pulumi.Output<string | undefined>;
    public readonly ignoreUnpackErrors!: pulumi.Output<boolean | undefined>;
    public readonly lock!: pulumi.Output<string | undefined>;
    public readonly memory!: pulumi.Output<number | undefined>;
    public readonly mountpoints!: pulumi.Output<outputs.LXCContainerMountpoint[] | undefined>;
    public readonly nameserver!: pulumi.Output<string | undefined>;
    public readonly networks!: pulumi.Output<outputs.LXCContainerNetwork[] | undefined>;
    public readonly onboot!: pulumi.Output<boolean | undefined>;
    public readonly ostemplate!: pulumi.Output<string | undefined>;
    public readonly ostype!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly pool!: pulumi.Output<string | undefined>;
    public readonly protection!: pulumi.Output<boolean | undefined>;
    public readonly restore!: pulumi.Output<boolean | undefined>;
    public readonly rootfs!: pulumi.Output<outputs.LXCContainerRootfs | undefined>;
    public readonly searchdomain!: pulumi.Output<string | undefined>;
    public readonly sshPublicKeys!: pulumi.Output<string | undefined>;
    public readonly start!: pulumi.Output<boolean | undefined>;
    public readonly startup!: pulumi.Output<string | undefined>;
    public readonly swap!: pulumi.Output<number | undefined>;
    public readonly tags!: pulumi.Output<string | undefined>;
    public readonly targetNode!: pulumi.Output<string>;
    public readonly template!: pulumi.Output<boolean | undefined>;
    public readonly tty!: pulumi.Output<number | undefined>;
    public readonly unique!: pulumi.Output<boolean | undefined>;
    public readonly unprivileged!: pulumi.Output<boolean | undefined>;
    public readonly unuseds!: pulumi.Output<string[] | undefined>;
    public readonly vmid!: pulumi.Output<number | undefined>;

    /**
     * Create a LXCContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LXCContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LXCContainerArgs | LXCContainerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LXCContainerState | undefined;
            inputs["arch"] = state ? state.arch : undefined;
            inputs["bwlimit"] = state ? state.bwlimit : undefined;
            inputs["cmode"] = state ? state.cmode : undefined;
            inputs["console"] = state ? state.console : undefined;
            inputs["cores"] = state ? state.cores : undefined;
            inputs["cpulimit"] = state ? state.cpulimit : undefined;
            inputs["cpuunits"] = state ? state.cpuunits : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["features"] = state ? state.features : undefined;
            inputs["force"] = state ? state.force : undefined;
            inputs["hookscript"] = state ? state.hookscript : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["ignoreUnpackErrors"] = state ? state.ignoreUnpackErrors : undefined;
            inputs["lock"] = state ? state.lock : undefined;
            inputs["memory"] = state ? state.memory : undefined;
            inputs["mountpoints"] = state ? state.mountpoints : undefined;
            inputs["nameserver"] = state ? state.nameserver : undefined;
            inputs["networks"] = state ? state.networks : undefined;
            inputs["onboot"] = state ? state.onboot : undefined;
            inputs["ostemplate"] = state ? state.ostemplate : undefined;
            inputs["ostype"] = state ? state.ostype : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["pool"] = state ? state.pool : undefined;
            inputs["protection"] = state ? state.protection : undefined;
            inputs["restore"] = state ? state.restore : undefined;
            inputs["rootfs"] = state ? state.rootfs : undefined;
            inputs["searchdomain"] = state ? state.searchdomain : undefined;
            inputs["sshPublicKeys"] = state ? state.sshPublicKeys : undefined;
            inputs["start"] = state ? state.start : undefined;
            inputs["startup"] = state ? state.startup : undefined;
            inputs["swap"] = state ? state.swap : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetNode"] = state ? state.targetNode : undefined;
            inputs["template"] = state ? state.template : undefined;
            inputs["tty"] = state ? state.tty : undefined;
            inputs["unique"] = state ? state.unique : undefined;
            inputs["unprivileged"] = state ? state.unprivileged : undefined;
            inputs["unuseds"] = state ? state.unuseds : undefined;
            inputs["vmid"] = state ? state.vmid : undefined;
        } else {
            const args = argsOrState as LXCContainerArgs | undefined;
            if (!args || args.targetNode === undefined) {
                throw new Error("Missing required property 'targetNode'");
            }
            inputs["arch"] = args ? args.arch : undefined;
            inputs["bwlimit"] = args ? args.bwlimit : undefined;
            inputs["cmode"] = args ? args.cmode : undefined;
            inputs["console"] = args ? args.console : undefined;
            inputs["cores"] = args ? args.cores : undefined;
            inputs["cpulimit"] = args ? args.cpulimit : undefined;
            inputs["cpuunits"] = args ? args.cpuunits : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["features"] = args ? args.features : undefined;
            inputs["force"] = args ? args.force : undefined;
            inputs["hookscript"] = args ? args.hookscript : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["ignoreUnpackErrors"] = args ? args.ignoreUnpackErrors : undefined;
            inputs["lock"] = args ? args.lock : undefined;
            inputs["memory"] = args ? args.memory : undefined;
            inputs["mountpoints"] = args ? args.mountpoints : undefined;
            inputs["nameserver"] = args ? args.nameserver : undefined;
            inputs["networks"] = args ? args.networks : undefined;
            inputs["onboot"] = args ? args.onboot : undefined;
            inputs["ostemplate"] = args ? args.ostemplate : undefined;
            inputs["ostype"] = args ? args.ostype : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["pool"] = args ? args.pool : undefined;
            inputs["protection"] = args ? args.protection : undefined;
            inputs["restore"] = args ? args.restore : undefined;
            inputs["rootfs"] = args ? args.rootfs : undefined;
            inputs["searchdomain"] = args ? args.searchdomain : undefined;
            inputs["sshPublicKeys"] = args ? args.sshPublicKeys : undefined;
            inputs["start"] = args ? args.start : undefined;
            inputs["startup"] = args ? args.startup : undefined;
            inputs["swap"] = args ? args.swap : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetNode"] = args ? args.targetNode : undefined;
            inputs["template"] = args ? args.template : undefined;
            inputs["tty"] = args ? args.tty : undefined;
            inputs["unique"] = args ? args.unique : undefined;
            inputs["unprivileged"] = args ? args.unprivileged : undefined;
            inputs["unuseds"] = args ? args.unuseds : undefined;
            inputs["vmid"] = args ? args.vmid : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LXCContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LXCContainer resources.
 */
export interface LXCContainerState {
    readonly arch?: pulumi.Input<string>;
    readonly bwlimit?: pulumi.Input<number>;
    readonly cmode?: pulumi.Input<string>;
    readonly console?: pulumi.Input<boolean>;
    readonly cores?: pulumi.Input<number>;
    readonly cpulimit?: pulumi.Input<number>;
    readonly cpuunits?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly features?: pulumi.Input<inputs.LXCContainerFeatures>;
    readonly force?: pulumi.Input<boolean>;
    readonly hookscript?: pulumi.Input<string>;
    readonly hostname?: pulumi.Input<string>;
    readonly ignoreUnpackErrors?: pulumi.Input<boolean>;
    readonly lock?: pulumi.Input<string>;
    readonly memory?: pulumi.Input<number>;
    readonly mountpoints?: pulumi.Input<pulumi.Input<inputs.LXCContainerMountpoint>[]>;
    readonly nameserver?: pulumi.Input<string>;
    readonly networks?: pulumi.Input<pulumi.Input<inputs.LXCContainerNetwork>[]>;
    readonly onboot?: pulumi.Input<boolean>;
    readonly ostemplate?: pulumi.Input<string>;
    readonly ostype?: pulumi.Input<string>;
    readonly password?: pulumi.Input<string>;
    readonly pool?: pulumi.Input<string>;
    readonly protection?: pulumi.Input<boolean>;
    readonly restore?: pulumi.Input<boolean>;
    readonly rootfs?: pulumi.Input<inputs.LXCContainerRootfs>;
    readonly searchdomain?: pulumi.Input<string>;
    readonly sshPublicKeys?: pulumi.Input<string>;
    readonly start?: pulumi.Input<boolean>;
    readonly startup?: pulumi.Input<string>;
    readonly swap?: pulumi.Input<number>;
    readonly tags?: pulumi.Input<string>;
    readonly targetNode?: pulumi.Input<string>;
    readonly template?: pulumi.Input<boolean>;
    readonly tty?: pulumi.Input<number>;
    readonly unique?: pulumi.Input<boolean>;
    readonly unprivileged?: pulumi.Input<boolean>;
    readonly unuseds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly vmid?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a LXCContainer resource.
 */
export interface LXCContainerArgs {
    readonly arch?: pulumi.Input<string>;
    readonly bwlimit?: pulumi.Input<number>;
    readonly cmode?: pulumi.Input<string>;
    readonly console?: pulumi.Input<boolean>;
    readonly cores?: pulumi.Input<number>;
    readonly cpulimit?: pulumi.Input<number>;
    readonly cpuunits?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly features?: pulumi.Input<inputs.LXCContainerFeatures>;
    readonly force?: pulumi.Input<boolean>;
    readonly hookscript?: pulumi.Input<string>;
    readonly hostname?: pulumi.Input<string>;
    readonly ignoreUnpackErrors?: pulumi.Input<boolean>;
    readonly lock?: pulumi.Input<string>;
    readonly memory?: pulumi.Input<number>;
    readonly mountpoints?: pulumi.Input<pulumi.Input<inputs.LXCContainerMountpoint>[]>;
    readonly nameserver?: pulumi.Input<string>;
    readonly networks?: pulumi.Input<pulumi.Input<inputs.LXCContainerNetwork>[]>;
    readonly onboot?: pulumi.Input<boolean>;
    readonly ostemplate?: pulumi.Input<string>;
    readonly ostype?: pulumi.Input<string>;
    readonly password?: pulumi.Input<string>;
    readonly pool?: pulumi.Input<string>;
    readonly protection?: pulumi.Input<boolean>;
    readonly restore?: pulumi.Input<boolean>;
    readonly rootfs?: pulumi.Input<inputs.LXCContainerRootfs>;
    readonly searchdomain?: pulumi.Input<string>;
    readonly sshPublicKeys?: pulumi.Input<string>;
    readonly start?: pulumi.Input<boolean>;
    readonly startup?: pulumi.Input<string>;
    readonly swap?: pulumi.Input<number>;
    readonly tags?: pulumi.Input<string>;
    readonly targetNode: pulumi.Input<string>;
    readonly template?: pulumi.Input<boolean>;
    readonly tty?: pulumi.Input<number>;
    readonly unique?: pulumi.Input<boolean>;
    readonly unprivileged?: pulumi.Input<boolean>;
    readonly unuseds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly vmid?: pulumi.Input<number>;
}
