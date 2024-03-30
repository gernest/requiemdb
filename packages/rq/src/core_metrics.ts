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

