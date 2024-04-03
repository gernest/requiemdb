import {
    Scan, Scan_SCOPE, Scan_TimeRange,
    Scan_Filter, Scan_BaseProp,
    Scan_AttributeProp, Data,
} from "./scan";
import { Struct } from "./struct";

export class Config {
    base: Scan
    constructor(scope: Scan_SCOPE) {
        this.base = Scan.create();
        this.base.scope = scope
    }

    public scan(): ScanData {
        const encode = Scan.toBinary(this.base)
        //@ts-ignore
        return new ScanData(RQ.Scan(encode))
    }


    public limit(num_samples: number) {
        this.base.limit = num_samples
        return this
    }

    public reverse() {
        this.base.reverse = true
        return this
    }

    public resourceSchema(schema: string) {
        return this.baseFilter(
            Scan_BaseProp.RESOURCE_SCHEMA,
            schema,
        )
    }

    public scopeSchema(schema: string) {
        return this.baseFilter(
            Scan_BaseProp.SCOPE_SCHEMA,
            schema,
        )
    }

    public scopeName(name: string) {
        return this.baseFilter(
            Scan_BaseProp.SCOPE_NAME,
            name,
        )
    }

    public scopeVersion(version: string) {
        return this.baseFilter(
            Scan_BaseProp.SCOPE_NAME,
            version,
        )
    }

    public name(value: string) {
        return this.baseFilter(
            Scan_BaseProp.NAME,
            value,
        )
    }

    public traceId(value: string) {
        return this.baseFilter(
            Scan_BaseProp.TRACE_ID,
            value,
        )
    }

    public spanId(value: string) {
        return this.baseFilter(
            Scan_BaseProp.SPAN_ID,
            value,
        )
    }

    public parentSpanId(value: string) {
        return this.baseFilter(
            Scan_BaseProp.SPAN_ID,
            value,
        )
    }

    public logLevel(value: string) {
        return this.baseFilter(
            Scan_BaseProp.LOGS_LEVEL,
            value,
        )
    }

    public resourceAttr(key: string, value: string) {
        return this.attrFilter(
            Scan_AttributeProp.RESOURCE_ATTRIBUTES,
            key, value,
        )
    }

    public scopeAttr(key: string, value: string) {
        return this.attrFilter(
            Scan_AttributeProp.SCOPE_ATTRIBUTES,
            key, value,
        )
    }

    public attr(key: string, value: string) {
        return this.attrFilter(
            Scan_AttributeProp.ATTRIBUTES,
            key, value,
        )
    }

    private baseFilter(prop: Scan_BaseProp, value: string) {
        return this.filter({
            value: {
                oneofKind: "base",
                base: {
                    prop, value
                }
            },
        })
    }

    private attrFilter(prop: Scan_AttributeProp, key: string, value: string) {
        return this.filter({
            value: {
                oneofKind: "attr",
                attr: {
                    prop, key, value,
                }
            }
        })
    }

    public filter(f: Scan_Filter) {
        this.base.filters.push(f)
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
        //@ts-ignore
        return this.setRange(TimeRange.Today())
    }

    public thisWeek() {
        //@ts-ignore
        return this.setRange(TimeRange.ThisWeek())
    }

    public thisYear() {
        //@ts-ignore
        return this.setRange(TimeRange.ThisYear())
    }

    /**
     * 
     * @param duration is ISO 8601 duration string
     * @returns 
     */
    public ago(duration: string) {
        //@ts-ignore
        return this.setRange(TimeRange.ago(duration))
    }

    public thisMonth() {
        //@ts-ignore
        return this.setRange(TimeRange.ThisWeek())
    }


    protected setRange(ts: any) {
        this.base.timeRange = this.createTimeRange(
            ts.FromUnix(),
            ts.ToUnix(),
        )
        return this
    }



    private createTimeRange(fromSecs: number, toSecs: number): Scan_TimeRange {
        return {
            start: { seconds: fromSecs, nanos: 0 },
            end: { seconds: toSecs, nanos: 0 },
        }
    }
}


export class ScanData {
    constructor(private ptr: any) { }

    toData(): Data {
        //@ts-ignore
        return Data.fromBinary(new Uint8Array(RQ.Marshal(this.ptr)))
    }

    formData(data: Data): ScanData {
        const binary = Data.toBinary(data)
        //@ts-ignore
        return new ScanData(RQ.Unmarshal(binary))
    }

    static is(value: any): boolean {
        return (value as ScanData).ptr != undefined;
    }
}

/**
 * Serialize value and exit the script. This must only be called once, subsequent calls have
 * no effect.
 * @param value 
 */
export const render = (value: Struct | Data | ScanData) => {
    if (Data.is(value)) {
        //@ts-ignore
        RQ.RenderData(Data.toBinary(value))
    }
    if (Struct.is(value)) {
        //@ts-ignore
        RQ.RenderStruct(Struct.toBinary(value))
    }
    if (ScanData.is(value)) {
        //@ts-ignore
        RQ.RenderNative(value.ptr)
    }
}


