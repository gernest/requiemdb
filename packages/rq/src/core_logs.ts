import { Config } from "./core";
import { Scan_SCOPE } from "./scan";

export class Logs extends Config {
    constructor() {
        super(Scan_SCOPE.LOGS)
    }
}

/**
 * Wraps initializing a Logs object.This is a helper to avoid doing new Logs 
 * when creating scripts.
 * 
 * @returns Logs object
 */
export function logs(): Logs {
    return new Logs()
}