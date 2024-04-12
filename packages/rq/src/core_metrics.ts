import { Config, RESOURCE } from "./core";
import { MetricsData } from "./otel";

export class Metrics extends Config {
    constructor(name?: string) {
        super(RESOURCE.METRICS)
        if (name) {
            this.name(name)
        }
    }

    query(): MetricsData {
        return this.scan().GetMetrics();
    }
}

