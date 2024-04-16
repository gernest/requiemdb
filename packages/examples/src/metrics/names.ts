import { Metrics, println } from "@requiemdb/rq";

/**
 * Prints names of observed metrics
 */
const data = (new Metrics()).query()
data.resource_metrics.forEach((rm) => {
    rm.scope_metrics.forEach((sm) => {
        sm.metrics.forEach((m) => {
            println(m.name)
        })
    })
})