import { MetricsData, TracesData, LogsData } from "./otel";

export interface Data {
    GetMetrics(): MetricsData
    GetLogs(): LogsData
    GetTraces(): TracesData
}

export enum RESOURCE {
    METRICS = "metrics",
    LOGS = "logs",
    TRACES = "traces",
}

export class Config {
    ptr: any
    constructor(resource: RESOURCE) {
        //@ts-ignore
        this.ptr = SCAN.Create(resource);
    }

    public scan(): Data {
        //@ts-ignore
        return RQ.Scan(this.ptr) as Data
    }


    /**
     * Limits the number of matched samples to process during scanning.
     * 
     * @param num_samples maximum number of matched samples to process
     * @returns 
     */
    public limit(num_samples: number) {
        this.ptr.Limit(num_samples)
        return this
    }

    /**
     * Duration relative to current scanning time to start evaluation.
     *   endTime = current_scan_time - offset 
     * 
     * @param duration is a Go time.Duration string
     * @returns 
     */
    public offset(duration: string) {
        this.ptr.Offset(duration)
        return this
    }

    /**
     * Sets scan evaluation time. All time computations during scanning will
     * be relative to this time. Useful for reproducible analysis
     * 
     * @param date 
     * @returns 
     */
    public now(date: Date) {
        this.ptr.Now(date.getTime())
        return this
    }

    public reverse() {
        this.ptr.Reverse()
        return this
    }

    public resourceSchema(schema: string) {
        this.ptr.ResourceSchema(schema)
        return this
    }

    public scopeSchema(schema: string) {
        this.ptr.ScopeSchema(schema)
        return this
    }

    public scopeName(name: string) {
        this.ptr.ScopeName(name)
        return this
    }

    public scopeVersion(version: string) {
        this.ptr.ScopeVersion(version)
        return this
    }

    public name(value: string) {
        this.ptr.Name(value)
        return this
    }

    public traceId(value: string) {
        this.ptr.TraceID(value)
        return this
    }

    public spanId(value: string) {
        this.ptr.SpanID(value)
        return this
    }

    public parentSpanId(value: string) {
        this.ptr.ParentSpanID(value)
        return this
    }

    public logLevel(value: string) {
        this.ptr.LogLevel(value)
        return this
    }

    public resourceAttr(key: string, value: string) {
        this.ptr.ResourceAttr(key, value)
        return this
    }

    public scopeAttr(key: string, value: string) {
        this.ptr.ScopeAttr(key, value)
        return this
    }

    public attr(key: string, value: string) {
        this.ptr.Attr(key, value)
        return this
    }

    /**
     * 
     * @returns samples for the last 15 minutes
     */
    public latest() {
        return this.ago("15m")
    }

    public today() {
        this.ptr.Today()
        return this
    }

    public thisWeek() {
        this.ptr.ThisWeek()
        return this
    }

    public thisMonth() {
        this.ptr.ThisMonth()
        return this
    }
    public thisYear() {
        this.ptr.ThisYear()
        return this
    }

    /**
     * 
     * @param duration is go time.Duration string
     * @returns 
     */
    public ago(duration: string) {
        this.ptr.Ago(duration)
        return this
    }
}

export interface JSONOptions {
    pretty: boolean
}

export interface TextOptions {
    /**
     * Includes resource information in the output.Includes resource schema 
     * and resource attributes
     */
    resource: boolean
    /**
     * Includes scope information in the output. Includes scope name, version, 
     * schema and attributes
     */
    scope: boolean
}

/**
 * writes args to script stdout
 * @param args 
 */
export function print(...args: any[]) {
    //@ts-ignore
    RQ.Print(...args)
}

/**
 * Like print but adds a new line
 * @param args 
 */
export function println(...args: any[]) {
    //@ts-ignore
    RQ.PrintLn(...args)
}