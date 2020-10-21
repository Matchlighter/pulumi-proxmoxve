// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmoxve
{
    public partial class Lxc_container : Pulumi.CustomResource
    {
        [Output("arch")]
        public Output<string?> Arch { get; private set; } = null!;

        [Output("bwlimit")]
        public Output<int?> Bwlimit { get; private set; } = null!;

        [Output("cmode")]
        public Output<string?> Cmode { get; private set; } = null!;

        [Output("console")]
        public Output<bool?> Console { get; private set; } = null!;

        [Output("cores")]
        public Output<int?> Cores { get; private set; } = null!;

        [Output("cpulimit")]
        public Output<int?> Cpulimit { get; private set; } = null!;

        [Output("cpuunits")]
        public Output<int?> Cpuunits { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("features")]
        public Output<ImmutableArray<Outputs.Lxc_containerFeature>> Features { get; private set; } = null!;

        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        [Output("hookscript")]
        public Output<string?> Hookscript { get; private set; } = null!;

        [Output("hostname")]
        public Output<string?> Hostname { get; private set; } = null!;

        [Output("ignoreUnpackErrors")]
        public Output<bool?> IgnoreUnpackErrors { get; private set; } = null!;

        [Output("lock")]
        public Output<string?> Lock { get; private set; } = null!;

        [Output("memory")]
        public Output<int?> Memory { get; private set; } = null!;

        [Output("mountpoints")]
        public Output<ImmutableArray<Outputs.Lxc_containerMountpoint>> Mountpoints { get; private set; } = null!;

        [Output("nameserver")]
        public Output<string?> Nameserver { get; private set; } = null!;

        [Output("networks")]
        public Output<ImmutableArray<Outputs.Lxc_containerNetwork>> Networks { get; private set; } = null!;

        [Output("onboot")]
        public Output<bool?> Onboot { get; private set; } = null!;

        [Output("ostemplate")]
        public Output<string?> Ostemplate { get; private set; } = null!;

        [Output("ostype")]
        public Output<string?> Ostype { get; private set; } = null!;

        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("pool")]
        public Output<string?> Pool { get; private set; } = null!;

        [Output("protection")]
        public Output<bool?> Protection { get; private set; } = null!;

        [Output("restore")]
        public Output<bool?> Restore { get; private set; } = null!;

        [Output("rootfs")]
        public Output<string?> Rootfs { get; private set; } = null!;

        [Output("searchdomain")]
        public Output<string?> Searchdomain { get; private set; } = null!;

        [Output("sshPublicKeys")]
        public Output<string?> SshPublicKeys { get; private set; } = null!;

        [Output("start")]
        public Output<bool?> Start { get; private set; } = null!;

        [Output("startup")]
        public Output<string?> Startup { get; private set; } = null!;

        [Output("storage")]
        public Output<string?> Storage { get; private set; } = null!;

        [Output("swap")]
        public Output<int?> Swap { get; private set; } = null!;

        [Output("targetNode")]
        public Output<string> TargetNode { get; private set; } = null!;

        [Output("template")]
        public Output<bool?> Template { get; private set; } = null!;

        [Output("tty")]
        public Output<int?> Tty { get; private set; } = null!;

        [Output("unique")]
        public Output<bool?> Unique { get; private set; } = null!;

        [Output("unprivileged")]
        public Output<bool?> Unprivileged { get; private set; } = null!;

        [Output("unuseds")]
        public Output<ImmutableArray<string>> Unuseds { get; private set; } = null!;

        [Output("vmid")]
        public Output<int?> Vmid { get; private set; } = null!;


        /// <summary>
        /// Create a Lxc_container resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Lxc_container(string name, Lxc_containerArgs args, CustomResourceOptions? options = null)
            : base("proxmoxve:index/lxc_container:lxc_container", name, args ?? new Lxc_containerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Lxc_container(string name, Input<string> id, Lxc_containerState? state = null, CustomResourceOptions? options = null)
            : base("proxmoxve:index/lxc_container:lxc_container", name, state, MakeResourceOptions(options, id))
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
        /// <summary>
        /// Get an existing Lxc_container resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Lxc_container Get(string name, Input<string> id, Lxc_containerState? state = null, CustomResourceOptions? options = null)
        {
            return new Lxc_container(name, id, state, options);
        }
    }

    public sealed class Lxc_containerArgs : Pulumi.ResourceArgs
    {
        [Input("arch")]
        public Input<string>? Arch { get; set; }

        [Input("bwlimit")]
        public Input<int>? Bwlimit { get; set; }

        [Input("cmode")]
        public Input<string>? Cmode { get; set; }

        [Input("console")]
        public Input<bool>? Console { get; set; }

        [Input("cores")]
        public Input<int>? Cores { get; set; }

        [Input("cpulimit")]
        public Input<int>? Cpulimit { get; set; }

        [Input("cpuunits")]
        public Input<int>? Cpuunits { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("features")]
        private InputList<Inputs.Lxc_containerFeatureArgs>? _features;
        public InputList<Inputs.Lxc_containerFeatureArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.Lxc_containerFeatureArgs>());
            set => _features = value;
        }

        [Input("force")]
        public Input<bool>? Force { get; set; }

        [Input("hookscript")]
        public Input<string>? Hookscript { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("ignoreUnpackErrors")]
        public Input<bool>? IgnoreUnpackErrors { get; set; }

        [Input("lock")]
        public Input<string>? Lock { get; set; }

        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("mountpoints")]
        private InputList<Inputs.Lxc_containerMountpointArgs>? _mountpoints;
        public InputList<Inputs.Lxc_containerMountpointArgs> Mountpoints
        {
            get => _mountpoints ?? (_mountpoints = new InputList<Inputs.Lxc_containerMountpointArgs>());
            set => _mountpoints = value;
        }

        [Input("nameserver")]
        public Input<string>? Nameserver { get; set; }

        [Input("networks")]
        private InputList<Inputs.Lxc_containerNetworkArgs>? _networks;
        public InputList<Inputs.Lxc_containerNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.Lxc_containerNetworkArgs>());
            set => _networks = value;
        }

        [Input("onboot")]
        public Input<bool>? Onboot { get; set; }

        [Input("ostemplate")]
        public Input<string>? Ostemplate { get; set; }

        [Input("ostype")]
        public Input<string>? Ostype { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("pool")]
        public Input<string>? Pool { get; set; }

        [Input("protection")]
        public Input<bool>? Protection { get; set; }

        [Input("restore")]
        public Input<bool>? Restore { get; set; }

        [Input("rootfs")]
        public Input<string>? Rootfs { get; set; }

        [Input("searchdomain")]
        public Input<string>? Searchdomain { get; set; }

        [Input("sshPublicKeys")]
        public Input<string>? SshPublicKeys { get; set; }

        [Input("start")]
        public Input<bool>? Start { get; set; }

        [Input("startup")]
        public Input<string>? Startup { get; set; }

        [Input("storage")]
        public Input<string>? Storage { get; set; }

        [Input("swap")]
        public Input<int>? Swap { get; set; }

        [Input("targetNode", required: true)]
        public Input<string> TargetNode { get; set; } = null!;

        [Input("template")]
        public Input<bool>? Template { get; set; }

        [Input("tty")]
        public Input<int>? Tty { get; set; }

        [Input("unique")]
        public Input<bool>? Unique { get; set; }

        [Input("unprivileged")]
        public Input<bool>? Unprivileged { get; set; }

        [Input("unuseds")]
        private InputList<string>? _unuseds;
        public InputList<string> Unuseds
        {
            get => _unuseds ?? (_unuseds = new InputList<string>());
            set => _unuseds = value;
        }

        [Input("vmid")]
        public Input<int>? Vmid { get; set; }

        public Lxc_containerArgs()
        {
        }
    }

    public sealed class Lxc_containerState : Pulumi.ResourceArgs
    {
        [Input("arch")]
        public Input<string>? Arch { get; set; }

        [Input("bwlimit")]
        public Input<int>? Bwlimit { get; set; }

        [Input("cmode")]
        public Input<string>? Cmode { get; set; }

        [Input("console")]
        public Input<bool>? Console { get; set; }

        [Input("cores")]
        public Input<int>? Cores { get; set; }

        [Input("cpulimit")]
        public Input<int>? Cpulimit { get; set; }

        [Input("cpuunits")]
        public Input<int>? Cpuunits { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("features")]
        private InputList<Inputs.Lxc_containerFeatureGetArgs>? _features;
        public InputList<Inputs.Lxc_containerFeatureGetArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.Lxc_containerFeatureGetArgs>());
            set => _features = value;
        }

        [Input("force")]
        public Input<bool>? Force { get; set; }

        [Input("hookscript")]
        public Input<string>? Hookscript { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("ignoreUnpackErrors")]
        public Input<bool>? IgnoreUnpackErrors { get; set; }

        [Input("lock")]
        public Input<string>? Lock { get; set; }

        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("mountpoints")]
        private InputList<Inputs.Lxc_containerMountpointGetArgs>? _mountpoints;
        public InputList<Inputs.Lxc_containerMountpointGetArgs> Mountpoints
        {
            get => _mountpoints ?? (_mountpoints = new InputList<Inputs.Lxc_containerMountpointGetArgs>());
            set => _mountpoints = value;
        }

        [Input("nameserver")]
        public Input<string>? Nameserver { get; set; }

        [Input("networks")]
        private InputList<Inputs.Lxc_containerNetworkGetArgs>? _networks;
        public InputList<Inputs.Lxc_containerNetworkGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.Lxc_containerNetworkGetArgs>());
            set => _networks = value;
        }

        [Input("onboot")]
        public Input<bool>? Onboot { get; set; }

        [Input("ostemplate")]
        public Input<string>? Ostemplate { get; set; }

        [Input("ostype")]
        public Input<string>? Ostype { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("pool")]
        public Input<string>? Pool { get; set; }

        [Input("protection")]
        public Input<bool>? Protection { get; set; }

        [Input("restore")]
        public Input<bool>? Restore { get; set; }

        [Input("rootfs")]
        public Input<string>? Rootfs { get; set; }

        [Input("searchdomain")]
        public Input<string>? Searchdomain { get; set; }

        [Input("sshPublicKeys")]
        public Input<string>? SshPublicKeys { get; set; }

        [Input("start")]
        public Input<bool>? Start { get; set; }

        [Input("startup")]
        public Input<string>? Startup { get; set; }

        [Input("storage")]
        public Input<string>? Storage { get; set; }

        [Input("swap")]
        public Input<int>? Swap { get; set; }

        [Input("targetNode")]
        public Input<string>? TargetNode { get; set; }

        [Input("template")]
        public Input<bool>? Template { get; set; }

        [Input("tty")]
        public Input<int>? Tty { get; set; }

        [Input("unique")]
        public Input<bool>? Unique { get; set; }

        [Input("unprivileged")]
        public Input<bool>? Unprivileged { get; set; }

        [Input("unuseds")]
        private InputList<string>? _unuseds;
        public InputList<string> Unuseds
        {
            get => _unuseds ?? (_unuseds = new InputList<string>());
            set => _unuseds = value;
        }

        [Input("vmid")]
        public Input<int>? Vmid { get; set; }

        public Lxc_containerState()
        {
        }
    }
}