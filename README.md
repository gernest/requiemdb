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

Let's check system cpu time, create a file `cpu.ts`
```ts
// cpu.ts
import { Metrics } from "@requiemdb/rq";

/**
 * Query instant system.cpu.time
 */
Metrics.render(
    (new Metrics()).
        name("system.cpu.time").
        query()
)
```

```bash
rq query cpu.ts
```
```
system.cpu.time
TIMESTAMP            VALUE         ATTRIBUTES            
2024-04-03 14:37:40  193h2m11.21s  { state = "idle" }    
2024-04-03 14:37:40  13h36m20.38s  { state = "user" }    
2024-04-03 14:37:40  8h20m8.09s    { state = "system" }  
2024-04-03 14:37:40  0s            { state = "other" }   
```

