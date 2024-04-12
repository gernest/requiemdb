# requiemdb

Permanent Storage for Open Telemetry Data.

# Features

- **OTLP gRPC server**: You can send metrics, traces and logs directly
 from your apps using otel gRPC exporter. Works with all supported language SDK.
- **Query as code**: version, reuse , run or experiment with scripts.
- **Embedded javascript engine**: Use modern javascript or typescript for querying.
- **Standard compliant data**: Work with data as defined in Open Telemetry Standards

 You don't need bespoke query language to understand and get insight from your
 applications. Work with samples as defined in [Open Telemetry Protocol](https://github.com/open-telemetry/opentelemetry-proto) using modern javascript or typescript.


## Example querying instant metrics

```ts
import { Metrics, render } from "@requiemdb/rq";

/**
 *  Instant Vectors
 */
render(
    (new Metrics())
        .name("http_requests_total")
        .query(),
)
```


# Getting started

## Installation

You can see release page for downloads if you have go installed do this

```bash
go install github.com/gernest/requiemdb/cmd/rq@latest
```

This will have `rq` binary installed

## Start the server

```
rq
```

Wait for 2 minutes to have `rq`  collect the process stats so we can query them.

Let's check the number of active goroutines our process has, create a file `goroutines.ts`
and paste this content.

```ts
// goroutines.ts
import { Metrics, render } from "@requiemdb/rq";

/**
 *  Instant Vectors
 */
render(
    (new Metrics())
        .name("process.runtime.go.goroutines")
        .query(),
)
```

```bash
rq query goroutines.ts
```


