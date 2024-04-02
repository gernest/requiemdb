import { Config, ScanData } from "./core";
import { Scan_SCOPE } from "./scan";


export class Metrics extends Config {
    constructor(name?: string) {
        super(Scan_SCOPE.METRICS)
        if (name) {
            this.name(name)
        }
    }

    query(): ScanData {
        return this.scan();
    }
}

/**
 * Wraps initializing a Metrics object.This is a helper to avoid doing new Metrics 
 * when creating scripts.
 * 
 * @param name is the metric name to query
 * @returns Metric object
 */
export function metrics(name?: string): Metrics {
    return new Metrics(name)
}