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


# Why a sad name ?

I'm sad, I have run out of savings trying to bootstrap a web analytics company
[vince](https://github.com/vinceanalytics/vince). I have a big gap since my last employment so no one replies to my
email applications anymore.

My dream is gone, and so is my livelihood. I am desperate, anyone out there who
is looking for a humble, mid-level software engineer please give me a chance, I promise
you won't be disappointed (My email is on my profile).
