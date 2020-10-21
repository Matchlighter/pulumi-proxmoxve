// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the proxmoxve package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'proxmoxve';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            if (!args || args.pmApiUrl === undefined) {
                throw new Error("Missing required property 'pmApiUrl'");
            }
            if (!args || args.pmPassword === undefined) {
                throw new Error("Missing required property 'pmPassword'");
            }
            if (!args || args.pmUser === undefined) {
                throw new Error("Missing required property 'pmUser'");
            }
            inputs["pmApiUrl"] = args ? args.pmApiUrl : undefined;
            inputs["pmLogEnable"] = pulumi.output(args ? args.pmLogEnable : undefined).apply(JSON.stringify);
            inputs["pmLogFile"] = args ? args.pmLogFile : undefined;
            inputs["pmLogLevels"] = pulumi.output(args ? args.pmLogLevels : undefined).apply(JSON.stringify);
            inputs["pmOtp"] = args ? args.pmOtp : undefined;
            inputs["pmParallel"] = pulumi.output(args ? args.pmParallel : undefined).apply(JSON.stringify);
            inputs["pmPassword"] = args ? args.pmPassword : undefined;
            inputs["pmTimeout"] = pulumi.output(args ? args.pmTimeout : undefined).apply(JSON.stringify);
            inputs["pmTlsInsecure"] = pulumi.output(args ? args.pmTlsInsecure : undefined).apply(JSON.stringify);
            inputs["pmUser"] = args ? args.pmUser : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * https://host.fqdn:8006/api2/json
     */
    readonly pmApiUrl: pulumi.Input<string>;
    readonly pmLogEnable?: pulumi.Input<boolean>;
    readonly pmLogFile?: pulumi.Input<string>;
    readonly pmLogLevels?: pulumi.Input<{[key: string]: any}>;
    /**
     * OTP 2FA code (if required)
     */
    readonly pmOtp?: pulumi.Input<string>;
    readonly pmParallel?: pulumi.Input<number>;
    /**
     * secret
     */
    readonly pmPassword: pulumi.Input<string>;
    readonly pmTimeout?: pulumi.Input<number>;
    readonly pmTlsInsecure?: pulumi.Input<boolean>;
    /**
     * username, maywith with @pam
     */
    readonly pmUser: pulumi.Input<string>;
}