import { Config, ScanData, RESOURCE } from "./core";


export class Metrics extends Config {
    constructor(name?: string) {
        super(RESOURCE.METRICS)
        if (name) {
            this.name(name)
        }
    }

    query(): ScanData {
        return this.scan();
    }
}

