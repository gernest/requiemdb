import { Metrics, render } from "@requiemdb/rq";

render(
    (new Metrics("http_requests_total"))
        .query(),
)

