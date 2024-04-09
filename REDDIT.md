# Permanent Storage For Open Telemetry Data

 `requiemdb` is a pure `go` database for open telemetry that uses modern
javascript or typescript as query language.


# Technology used

- [Open Telemetry](https://github.com/open-telemetry) standard for Metrics, Traces and Logs collection.
- [badger](https://github.com/dgraph-io/badger) the underlying key/value store
- [Roaring Bitmaps](https://github.com/RoaringBitmap/roaring) for labels indexing
- [Apache Arrow](https://github.com/apache/arrow/tree/main/go) for global sample metadata indexing
- [goja](https://github.com/dop251/goja) Embedded JS engine to power query api

# Why ?

All existing solutions are not purpose built for open telemetry.A lot of
information is lost during ingestion.

I wanted to work with samples as they were observed, and there is no existing
solution for this.


# How ?

Processing is done at Sample level. A reverse index that maps to samples is
generated during ingestion. Samples are serialized and compressed using zstd
and stored in a key value store.

Extensive use of roaring bitmaps for the index combined with apache arrow
allows faster and efficient sample lookup.

Fundamental data stored is a union of otel sample data 

```proto3
message Data {
  oneof data {
    opentelemetry.proto.metrics.v1.MetricsData metrics = 1;
    opentelemetry.proto.logs.v1.LogsData logs = 2;
    opentelemetry.proto.trace.v1.TracesData trace = 3;
  }
}
```


# Why Typescript as query Language ?

It is easy to learn but powerful, I consider this choice to be tactical. 


# Show me the code 

```js
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
you won't be disappointed (My email is on my github profile).

.


