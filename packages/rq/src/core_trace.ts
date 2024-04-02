import { Config } from "./core";
import { Scan_SCOPE } from "./scan";

export class Trace extends Config {
    constructor() {
        super(Scan_SCOPE.TRACES)
    }
}


/**
 * Wraps initializing a Trace object.This is a helper to avoid doing new Trace 
 * when creating scripts.
 * 
 * @returns Trace object
 */
export function trace(): Trace {
    return new Trace()
}