import { Config } from "./core";
import { Scan_SCOPE } from "./scan";

export class Trace extends Config {
    constructor() {
        super(Scan_SCOPE.TRACES)
    }
}