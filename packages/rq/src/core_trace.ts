import { Config, RESOURCE } from "./core";
import { TracesData } from "./otel";

export class Trace extends Config {
    constructor() {
        super(RESOURCE.TRACES)
    }

    query(): TracesData {
        return this.scan().GetTraces()
    }
}
