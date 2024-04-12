import { Config, RESOURCE } from "./core";
import { LogsData } from "./otel";

export class Logs extends Config {
    constructor() {
        super(RESOURCE.LOGS)
    }

    query(): LogsData {
        return this.scan().GetLogs()
    }
}
