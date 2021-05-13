// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmoxve.Inputs
{

    public sealed class LXCContainerMountpointGetArgs : Pulumi.ResourceArgs
    {
        [Input("acl")]
        public Input<bool>? Acl { get; set; }

        [Input("backup")]
        public Input<bool>? Backup { get; set; }

        [Input("file")]
        public Input<string>? File { get; set; }

        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("mp", required: true)]
        public Input<string> Mp { get; set; } = null!;

        [Input("quota")]
        public Input<bool>? Quota { get; set; }

        [Input("replicate")]
        public Input<bool>? Replicate { get; set; }

        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        [Input("slot", required: true)]
        public Input<int> Slot { get; set; } = null!;

        [Input("storage", required: true)]
        public Input<string> Storage { get; set; } = null!;

        [Input("volume")]
        public Input<string>? Volume { get; set; }

        public LXCContainerMountpointGetArgs()
        {
        }
    }
}
