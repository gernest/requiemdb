"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Attributes = exports.Attribute = void 0;
class Attribute {
    constructor(ptr) {
        this.ptr = ptr;
    }
    key() {
        return this.ptr.GetKey();
    }
    value() {
        const v = this.ptr.GetValue();
        if (v) {
            return v.GetStringValue();
        }
        return "";
    }
}
exports.Attribute = Attribute;
class Attributes {
    constructor(ptr, size) {
        this.ptr = ptr;
        this.size = size;
        this.pos = 0;
    }
    next() {
        if (this.pos < this.size) {
            const v = this.ptr[this.pos];
            this.pos++;
            return { done: false, value: new Attribute(v) };
        }
        return { done: true, value: new Attribute(undefined) };
    }
}
exports.Attributes = Attributes;
