// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Component extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'example::Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    public readonly a!: pulumi.Output<boolean>;
    public readonly b!: pulumi.Output<boolean | undefined>;
    public readonly bar!: pulumi.Output<outputs.Foo | undefined>;
    public readonly baz!: pulumi.Output<outputs.Foo[] | undefined>;
    public readonly c!: pulumi.Output<number>;
    public readonly d!: pulumi.Output<number | undefined>;
    public readonly e!: pulumi.Output<string>;
    public readonly f!: pulumi.Output<string | undefined>;
    public readonly foo!: pulumi.Output<outputs.Foo | undefined>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.a === undefined) && !opts.urn) {
                throw new Error("Missing required property 'a'");
            }
            if ((!args || args.c === undefined) && !opts.urn) {
                throw new Error("Missing required property 'c'");
            }
            if ((!args || args.e === undefined) && !opts.urn) {
                throw new Error("Missing required property 'e'");
            }
            inputs["a"] = args ? args.a : undefined;
            inputs["b"] = args ? args.b : undefined;
            inputs["bar"] = args ? args.bar : undefined;
            inputs["baz"] = args ? args.baz : undefined;
            inputs["bazMap"] = args ? args.bazMap : undefined;
            inputs["c"] = args ? args.c : undefined;
            inputs["d"] = args ? args.d : undefined;
            inputs["e"] = args ? args.e : undefined;
            inputs["f"] = args ? args.f : undefined;
            inputs["foo"] = args ? args.foo : undefined;
        } else {
            inputs["a"] = undefined /*out*/;
            inputs["b"] = undefined /*out*/;
            inputs["bar"] = undefined /*out*/;
            inputs["baz"] = undefined /*out*/;
            inputs["c"] = undefined /*out*/;
            inputs["d"] = undefined /*out*/;
            inputs["e"] = undefined /*out*/;
            inputs["f"] = undefined /*out*/;
            inputs["foo"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Component.__pulumiType, name, inputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    a: boolean;
    b?: boolean;
    bar?: inputs.Foo;
    baz?: inputs.Foo[];
    bazMap?: {[key: string]: inputs.Foo};
    c: number;
    d?: number;
    e: string;
    f?: string;
    foo?: pulumi.Input<inputs.FooArgs>;
}