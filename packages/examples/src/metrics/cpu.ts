import { Metrics } from "@requiemdb/rq";

/**
 * Query instant system.cpu.time
 */
Metrics.render(
    (new Metrics()).
        name("system.cpu.time").
        query()
)