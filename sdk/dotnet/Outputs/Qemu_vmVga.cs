// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmoxve.Outputs
{

    [OutputType]
    public sealed class Qemu_vmVga
    {
        public readonly int? Memory;
        public readonly string? Type;

        [OutputConstructor]
        private Qemu_vmVga(
            int? memory,

            string? type)
        {
            Memory = memory;
            Type = type;
        }
    }
}
