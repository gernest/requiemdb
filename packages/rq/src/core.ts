import {
    Scan, Scan_SCOPE, Scan_TimeRange,
    Scan_Filter, Scan_BaseFilter, Scan_BaseProp,
    Scan_AttributeProp, Scan_AttrFilter,
} from "./scan";
import { Timestamp } from "./timestamp";
import { Visitor } from "./visit";

export class Config {
    base: Scan
    range?: any
    constructor(scope: Scan_SCOPE) {
        this.base = Scan.create();
        this.base.scope = scope
        this.latest()
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

    createVisitor(): Visitor {
        const vs = new Visitor();
        return vs.timeRange(
            this.range.FromUnixNano(),
            this.range.ToUnixNano(),
        )
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
        this.range = ts
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
    /**
     * 
     * @param visitor 
     * @returns a new ScanData with samples matching visitor filter
     */
    visit(visitor: Visitor) {
        //@ts-ignore
        return new ScanData(RQ.Visit(this.ptr, visitor.ptr))
    }
}

export type BaseValue = number | string | boolean | ScanData | ScanData[];

export type Value = BaseValue | Record<string, BaseValue>;

export const render = (value: Value) => {
    //@ts-ignore
    RQ.Render(value)
}
