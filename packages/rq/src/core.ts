import {
    Scan, Scan_SCOPE, Scan_TimeRange,
    Scan_Filter, Scan_BaseFilter, Scan_BaseProp,
    Scan_AttributeProp, Scan_AttrFilter,
    Data,
} from "./scan";
import { Timestamp } from "./timestamp";

export class Config {
    base: Scan
    constructor(scope: Scan_SCOPE) {
        this.base = Scan.create();
        this.base.scope = scope
    }

    public scan(): ScanResult {
        const encode = Scan.toBinary(this.base)
        //@ts-ignore
        const result = RQ.Scan(encode)
        return new ScanResult(result)
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
        const b = Scan_BaseFilter.create();
        b.prop = prop;
        b.value = value;
        const f = Scan_Filter.create()
        f.value = {
            oneofKind: "base",
            base: b,
        }
        return this.filter(f)
    }

    private attrFilter(prop: Scan_AttributeProp, key: string, value: string) {
        const b = Scan_AttrFilter.create();
        b.prop = prop;
        b.key = key;
        b.value = value;
        const f = Scan_Filter.create()
        f.value = {
            oneofKind: "attr",
            attr: b,
        }
        return this.filter(f)
    }

    public filter(f: Scan_Filter) {
        this.base.filters.push(f)
        return this
    }

    public today() {
        //@ts-ignore
        const ts = TimeRange.Today();
        this.base.timeRange = this.createTimeRange(
            ts.FromSeconds(),
            ts.ToSeconds(),
        )
        return this
    }



    private createTimeRange(fromSecs: number, toSecs: number): Scan_TimeRange {
        const from = Timestamp.create();
        from.seconds = fromSecs;
        const to = Timestamp.create();
        to.seconds = toSecs;
        const range = Scan_TimeRange.create()
        range.start = from;
        range.end = to;
        return range
    }
}


export class ScanResult {
    constructor(private ptr: any) { }

    [Symbol.iterator]() {
        return {
            next: () => {
                if (this.ptr.Next()) {
                    return { done: false, value: new ScanData(this.ptr.Current()) }
                }
                return { done: true }
            }
        }
    };
}

export class ScanData {
    constructor(private ptr: any) { }
}

