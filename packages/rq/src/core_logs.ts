import { Config } from "./core";
import { Scan_SCOPE } from "./scan";

export class Logs extends Config {
    constructor() {
        super(Scan_SCOPE.LOGS)
    }
}