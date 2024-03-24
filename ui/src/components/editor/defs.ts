export const rqDefinitions=`export declare class Attribute {
    private ptr;
    constructor(ptr: any);
    key(): string;
    value(): string;
}
export declare class Attributes {
    private ptr;
    private size;
    protected pos: number;
    constructor(ptr: any, size: number);
    next(): {
        done: boolean;
        value?: Attribute;
    };
}
`