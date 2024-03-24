

export class Attribute {
    constructor(private ptr: any) { }
    key(): string {
        return this.ptr.GetKey()
    }
    value(): string {
        const v = this.ptr.GetValue();
        if (v) {
            return v.GetStringValue()
        }
        return ""
    }
}

export class Attributes {
    protected pos: number = 0
    constructor(private ptr: any, private size: number) { }

    next(): { done: boolean, value?: Attribute } {
        if (this.pos < this.size) {
            const v = this.ptr[this.pos];
            this.pos++
            return { done: false, value: new Attribute(v) }
        }
        return { done: true, value: new Attribute(undefined) }
    }
}

