import { Config, RESOURCE, TextOptions, JSONOptions } from "./core";
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


    static renderJSON(data: MetricsData, options?: JSONOptions) {
        //@ts-ignore
        RQ.RenderMetricsDataJSON(data, options)
    }
    /**
     * Renders data in a native text format.
     * @param data 
     */
    static render(data: MetricsData, options?: TextOptions & {
        /**
         * Includes metrics information in output. Include metrics description and unit
         */
        metrics: boolean
    }) {
        //@ts-ignore
        RQ.RenderMetricsData(data, options)
    }
}

