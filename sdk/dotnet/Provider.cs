// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmoxve
{
    /// <summary>
    /// The provider type for the proxmoxve package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("proxmoxve", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API TokenID e.g. root@pam!mytesttoken
        /// </summary>
        [Input("pmApiTokenId")]
        public Input<string>? PmApiTokenId { get; set; }

        /// <summary>
        /// The secret uuid corresponding to a TokenID
        /// </summary>
        [Input("pmApiTokenSecret")]
        public Input<string>? PmApiTokenSecret { get; set; }

        /// <summary>
        /// https://host.fqdn:8006/api2/json
        /// </summary>
        [Input("pmApiUrl", required: true)]
        public Input<string> PmApiUrl { get; set; } = null!;

        /// <summary>
        /// By default this provider will exit if an unknown attribute is found. This is to prevent the accidential destruction of
        /// VMs or Data when something in the proxmox API has changed/updated and is not confirmed to work with this provider. Set
        /// this to true at your own risk. It may allow you to proceed in cases when the provider refuses to work, but be aware of
        /// the danger in doing so.
        /// </summary>
        [Input("pmDangerouslyIgnoreUnknownAttributes", json: true)]
        public Input<bool>? PmDangerouslyIgnoreUnknownAttributes { get; set; }

        /// <summary>
        /// Enable provider logging to get proxmox API logs
        /// </summary>
        [Input("pmLogEnable", json: true)]
        public Input<bool>? PmLogEnable { get; set; }

        /// <summary>
        /// Write logs to this specific file
        /// </summary>
        [Input("pmLogFile")]
        public Input<string>? PmLogFile { get; set; }

        [Input("pmLogLevels", json: true)]
        private InputMap<object>? _pmLogLevels;

        /// <summary>
        /// Configure the logging level to display; trace, debug, info, warn, etc
        /// </summary>
        public InputMap<object> PmLogLevels
        {
            get => _pmLogLevels ?? (_pmLogLevels = new InputMap<object>());
            set => _pmLogLevels = value;
        }

        /// <summary>
        /// OTP 2FA code (if required)
        /// </summary>
        [Input("pmOtp")]
        public Input<string>? PmOtp { get; set; }

        [Input("pmParallel", json: true)]
        public Input<int>? PmParallel { get; set; }

        /// <summary>
        /// Password to authenticate into proxmox
        /// </summary>
        [Input("pmPassword")]
        public Input<string>? PmPassword { get; set; }

        [Input("pmTimeout", json: true)]
        public Input<int>? PmTimeout { get; set; }

        /// <summary>
        /// By default, every TLS connection is verified to be secure. This option allows terraform to proceed and operate on
        /// servers considered insecure. For example if you're connecting to a remote host and you do not have the CA cert that
        /// issued the proxmox api url's certificate.
        /// </summary>
        [Input("pmTlsInsecure", json: true)]
        public Input<bool>? PmTlsInsecure { get; set; }

        /// <summary>
        /// Username e.g. myuser or myuser@pam
        /// </summary>
        [Input("pmUser")]
        public Input<string>? PmUser { get; set; }

        public ProviderArgs()
        {
        }
    }
}
