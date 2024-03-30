
export class Visitor {
    ptr: any
    constructor() {
        //@ts-ignore
        this.ptr = RQ.CreateVisitor();
    }


    timeRange(start_nano: number, end_nano: number) {
        this.ptr.SetTimeRange(start_nano, end_nano)
        return this;
    }

    resourceSchema(schema: string) {
        this.ptr.SetResourceSchema(schema)
        return this;
    }

    resourceAttr(key: string, value: string) {
        this.ptr.SetResourceAttr(key, value)
        return this;
    }

    scopeSchema(schema: string) {
        this.ptr.SetScopeSchema(schema)
        return this;
    }

    scopeName(name: string) {
        this.ptr.SetScopeName(name)
        return this;
    }

    scopeVersion(version: string) {
        this.ptr.SetScopeVersion(version)
        return this;
    }

    scopeAttr(key: string, value: string) {
        this.ptr.SetScopeAttr(key, value)
        return this;
    }

    attr(key: string, value: string) {
        this.ptr.SetAttr(key, value)
        return this;
    }

    name(name: string) {
        this.ptr.SetName(name)
        return this;
    }

    traceId(id: string) {
        this.ptr.SetTraceID(id)
        return this;
    }

    spanId(id: string) {
        this.ptr.SetSpanID(id)
        return this;
    }

    parentSpanId(id: string) {
        this.ptr.SetParentSpanID(id)
        return this;
    }

    logLevel(lvl: string) {
        this.ptr.SetLogLevel(lvl)
        return this;
    }
}