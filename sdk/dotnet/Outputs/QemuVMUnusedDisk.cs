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
    public sealed class QemuVMUnusedDisk
    {
        public readonly string? File;
        public readonly int? Slot;
        public readonly string? Storage;

        [OutputConstructor]
        private QemuVMUnusedDisk(
            string? file,

            int? slot,

            string? storage)
        {
            File = file;
            Slot = slot;
            Storage = storage;
        }
    }
}