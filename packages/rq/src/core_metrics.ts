import { Config, ScanData, ScanResult } from "./core";
import { Scan_SCOPE } from "./scan";


export interface MetricsResultMiddleware {
    applyResult(data: ScanResult): ScanData[]
}

export interface MetricsMiddleware {
    applyData(data: ScanData): ScanData | null | undefined
}

export class Metrics extends Config {
    constructor() {
        super(Scan_SCOPE.METRICS)
    }

    query(middleware?: MetricsMiddleware): ScanData[] {
        const result = this.scan();
        const collect: ScanData[] = [];
        for (let data of result) {
            if (data) {
                if (middleware) {
                    const applied = middleware.applyData(data)
                    if (applied) {
                        collect.push(data)
                    }
                }
            }
        }
        return collect;
    }
    /**
     * Like query but applies MetricsResultMiddleware
     * @param middleware 
     */
    queryFull(middleware?: MetricsResultMiddleware) {
        const result = this.scan();
        if (middleware) {
            return middleware.applyResult(result);
        }
        const collect: ScanData[] = [];
        // same as plain query
        for (let data of result) {
            if (data) {
                collect.push(data)
            }
        }
        return;
    }
}

