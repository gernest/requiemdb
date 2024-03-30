import { Metrics } from "@requiemdb/rq";


const totalRequests =
    (new Metrics())
        .name("http_requests_total")
        .query();