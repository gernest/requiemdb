
export interface TracesData {
    resource_spans: ResourceSpans[]
}

export interface ResourceSpans {
    resource: Resource
    scope_spans: ScopeSpans[]
    schema_url: string
}

export interface ScopeSpans {
    scope: InstrumentationScope
    spans: Span[]
    schema_url: string
}

export interface Span {
    trace_id: Uint8Array
    span_id: Uint8Array
    trace_state: string
    parent_span_id: Uint8Array
    name: string
    kind: SpanKind
    start_time_unix_nano: number
    end_time_unix_nano: number
    attributes: KeyValue[]
    events: Event[]
    links: Link[]
    status: Status
}

export enum SpanKind {
    SPAN_KIND_UNSPECIFIED = 0,
    SPAN_KIND_INTERNAL = 1,
    SPAN_KIND_SERVER = 2,
    SPAN_KIND_CLIENT = 3,
    SPAN_KIND_PRODUCER = 4,
    SPAN_KIND_CONSUMER = 5,
}

export interface Event {
    time_unix_nano: number
    name: string
    attributes: KeyValue[]
}

export interface Link {
    trace_id: Uint8Array
    span_id: Uint8Array
    trace_state: string
    attributes: KeyValue[]
}

export interface Status {
    message: string
    code: StatusCode
}


export enum StatusCode {
    STATUS_CODE_UNSET = 0,
    STATUS_CODE_OK = 1,
    STATUS_CODE_ERROR = 2,
};

export interface LogsData {
    resource_logs: ResourceLogs[]
}

export interface ResourceLogs {
    resource: Resource
    scope_logs: ScopeLogs[]
    schema_url: string
}

export interface ScopeLogs {
    scope: InstrumentationScope
    log_records: LogRecord
    schema_url: string
}

export interface LogRecord {
    time_unix_nano: number
    observed_time_unix_nano: number
    severity_number: SeverityNumber
    severity_text: string
    body: AnyValue
    attributes: KeyValue[]
    trace_id: Uint8Array
    span_id: Uint8Array
}

enum SeverityNumber {
    SEVERITY_NUMBER_UNSPECIFIED = 0,
    SEVERITY_NUMBER_TRACE = 1,
    SEVERITY_NUMBER_TRACE2 = 2,
    SEVERITY_NUMBER_TRACE3 = 3,
    SEVERITY_NUMBER_TRACE4 = 4,
    SEVERITY_NUMBER_DEBUG = 5,
    SEVERITY_NUMBER_DEBUG2 = 6,
    SEVERITY_NUMBER_DEBUG3 = 7,
    SEVERITY_NUMBER_DEBUG4 = 8,
    SEVERITY_NUMBER_INFO = 9,
    SEVERITY_NUMBER_INFO2 = 10,
    SEVERITY_NUMBER_INFO3 = 11,
    SEVERITY_NUMBER_INFO4 = 12,
    SEVERITY_NUMBER_WARN = 13,
    SEVERITY_NUMBER_WARN2 = 14,
    SEVERITY_NUMBER_WARN3 = 15,
    SEVERITY_NUMBER_WARN4 = 16,
    SEVERITY_NUMBER_ERROR = 17,
    SEVERITY_NUMBER_ERROR2 = 18,
    SEVERITY_NUMBER_ERROR3 = 19,
    SEVERITY_NUMBER_ERROR4 = 20,
    SEVERITY_NUMBER_FATAL = 21,
    SEVERITY_NUMBER_FATAL2 = 22,
    SEVERITY_NUMBER_FATAL3 = 23,
    SEVERITY_NUMBER_FATAL4 = 24,
}

export interface MetricsData {
    resource_metrics: ResourceMetrics[]
}

export interface ResourceMetrics {
    resource?: Resource
    scope_metrics: ScopeMetrics[]
    schema_url: string
}

export interface ScopeMetrics {
    scope?: InstrumentationScope
    metrics: Metric[]
    schema_url: string
}

export interface Metric {
    name: string
    description: string
    unit: string
    GetGauge(): Gauge
    GetSum(): Sum
    GetHistogram(): Histogram
    GetExponentialHistogram(): ExponentialHistogram
    GetSummary(): Summary
}

export interface Gauge {
    data_points: NumberDataPoint[]
}

export interface Sum {
    data_points: NumberDataPoint[]
    aggregation_temporality: AggregationTemporality
    is_monotonic: boolean
}

export interface Histogram {
    data_points: HistogramDataPoint[]
    aggregation_temporality: AggregationTemporality
}

export interface ExponentialHistogram {
    data_points: ExponentialHistogramDataPoint[]
    aggregation_temporality: AggregationTemporality
}

export interface Summary {
    data_points: SummaryDataPoint[]
}

export enum AggregationTemporality {
    AGGREGATION_TEMPORALITY_UNSPECIFIED = 0,
    AGGREGATION_TEMPORALITY_DELTA = 1,
    AGGREGATION_TEMPORALITY_CUMULATIVE = 2,
}

export interface NumberDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    GetAsDouble(): number
    GetAsInt(): number
    exemplars: Exemplar[]
}

export interface HistogramDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    count: number
    sum?: number
    bucket_counts: number[]
    explicit_bounds: number[]
    exemplars: Exemplar[]
    min: number[]
    max: number[]
}

export interface ExponentialHistogramDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    count: number
    sum: number[]
    scale: number
    zero_count: number
    positive: Buckets
    negative: Buckets
    exemplars: Exemplar[]
    min: number[]
    max: number[]
    zero_threshold: number[]
}

export interface Buckets {
    offset: number
    bucket_counts: number[]
}

export interface SummaryDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    count: number
    sum: number
    quantile_values: ValueAtQuantile[]
}

export interface ValueAtQuantile {
    quantile: number
    value: number
}

export interface Exemplar {
    filtered_attributes: KeyValue[]
    time_unix_nano: number
    span_id?: Uint8Array
    trace_id?: Uint8Array
    GetAsDouble(): number
    GetAsInt(): number
}

export interface Resource {
    attributes?: KeyValue[]
}

export interface KeyValue {
    key: string
    value: AnyValue
}

export interface AnyValue {
    GetStringValue(): string
    GetBoolValue(): boolean
    GetIntValue(): boolean
}

export interface InstrumentationScope {
    name: string
    version: string
    attributes?: KeyValue[]
}
