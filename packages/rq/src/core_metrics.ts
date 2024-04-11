import { Config, ScanData, Resource } from "./core";


export class Metrics extends Config {
    constructor(name?: string) {
        super(Resource.METRICS)
        if (name) {
            this.name(name)
        }
    }

    query(): ScanData {
        return this.scan();
    }
}

