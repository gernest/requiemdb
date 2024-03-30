import { Metrics, render } from "@requiemdb/rq";

/**
 *  Instant Vectors
 */
render(
    (new Metrics("http_requests_total"))
        .query(),
)

// With attributes filter
render(
    (new Metrics("http_requests_total"))
        .attr("job", "rq")
        .attr("group", "canary")
        .query(),
)

/**
 * Range Vectors
 */
render(
    (new Metrics("http_requests_total"))
        .attr("job", "rq")
        .ago("5m")
        .query(),
)

