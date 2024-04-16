# Permanent Storage For Open Telemetry Data

 `requiemdb` is a pure `go` database for open telemetry that uses modern
javascript or typescript as query language.


# Technology used

- [Open Telemetry](https://github.com/open-telemetry) standard for Metrics, Traces and Logs collection.
- [badger](https://github.com/dgraph-io/badger) the underlying key/value store
- [Roaring Bitmaps](https://github.com/RoaringBitmap/roaring) for attributes indexing
- [goja](https://github.com/dop251/goja) Embedded JS engine to power query api

# Why ?

Early iteration of [vince](https://www.vinceanalytics.com/) relied on open telemetry to track events. Due to my limited
budget I needed a very cheap open telemetry storage that is capable of serving the samples as 
they were observed.

Requirements were

 - very low resource usage (compute/memory/storage)
 - understand Open Telemetry standards
 - very fast scans

I would do further processing locally on my dev machine(which is more powerful) If I needed,
`vince` as a company has since folded, so I am sharing this hoping someone will find useful.

# How ?

Open Telemetry Data is compressed and stored in a key value store. An index is
generated during ingestion mapping attributes and various interesting properties 
of samples to the incoming sample ID.

Scans are against samples not individual data points within the samples. It is
possible to additionally pre process these samples before serving them.

A `js` runtime is used to drive the `go` backend.

Basically when  calling   `rq query cpu.ts`,  depending on `cpu.ts` contents it might involve scanning for 
interesting samples remotely and further processing them locally, either natively in `go` or you can
lift the native values into `js` and deal with them as you please.

`rq query` works locally. For efficiency,  samples are always compressed
when they are not being processed.

We use `ztsd compressed gRPC` for communication. In short, the whole setup is
very cost effective and insanely powerful. You have access to your `Metrics`, `Traces`,
 and `Logs` in a single `js` script and you can interpret the data as you see fit.


# Why Typescript as query Language ?

It is easy to learn but powerful, I consider this choice to be tactical. 


# Show me the code 

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

# Why a sad name ?

I'm sad and broke, I have run out of savings trying to bootstrap a now defunct web analytics company
[vince](https://github.com/vinceanalytics/vince).

My dream is gone, and so is my livelihood. I am desperate, anyone out there who
is looking for a humble, mid-level software engineer please give me a chance, I promise
you won't be disappointed [Here is my Resume](https://github.com/gernest/resume).



# Why feature x is missing?

My priority right now is to find employment. Contributions are welcome, but my effort will only  go towards things that will increase my odds of getting a job at the moment.

After I sort my work situation, I will go back to this.


