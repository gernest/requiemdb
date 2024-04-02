import { metrics, render } from "@requiemdb/rq";

/**
 *  Instant Vectors
 */
render(
    metrics()
        .name("http_requests_total")
        .query(),
)

// With attributes filter
render(
    metrics()
        .name("http_requests_total")
        .attr("job", "rq")
        .attr("group", "canary")
        .query(),
)

/**
 * Range Vectors
 */
render(
    metrics()
        .name("http_requests_total")
        .attr("job", "rq")
        .ago("5m")
        .query(),
)

