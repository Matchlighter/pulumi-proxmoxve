// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Proxmoxve
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("proxmoxve");
        /// <summary>
        /// https://host.fqdn:8006/api2/json
        /// </summary>
        public static string? PmApiUrl { get; set; } = __config.Get("pmApiUrl");

        /// <summary>
        /// By default this provider will exit if an unknown attribute is found. This is to prevent the accidential destruction of
        /// VMs or Data when something in the proxmox API has changed/updated and is not confirmed to work with this provider. Set
        /// this to true at your own risk. It may allow you to proceed in cases when the provider refuses to work, but be aware of
        /// the danger in doing so.
        /// </summary>
        public static bool? PmDangerouslyIgnoreUnknownAttributes { get; set; } = __config.GetBoolean("pmDangerouslyIgnoreUnknownAttributes");

        public static bool? PmLogEnable { get; set; } = __config.GetBoolean("pmLogEnable");

        public static string? PmLogFile { get; set; } = __config.Get("pmLogFile");

        public static ImmutableDictionary<string, object>? PmLogLevels { get; set; } = __config.GetObject<ImmutableDictionary<string, object>>("pmLogLevels");

        /// <summary>
        /// OTP 2FA code (if required)
        /// </summary>
        public static string? PmOtp { get; set; } = __config.Get("pmOtp");

        public static int? PmParallel { get; set; } = __config.GetInt32("pmParallel");

        /// <summary>
        /// secret
        /// </summary>
        public static string? PmPassword { get; set; } = __config.Get("pmPassword");

        public static int? PmTimeout { get; set; } = __config.GetInt32("pmTimeout");

        public static bool? PmTlsInsecure { get; set; } = __config.GetBoolean("pmTlsInsecure");

        /// <summary>
        /// username, maywith with @pam
        /// </summary>
        public static string? PmUser { get; set; } = __config.Get("pmUser");

    }
}
